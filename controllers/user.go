package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"project-tracker/api/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func SignToken(userId int32) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"ttl":    time.Now().Add(time.Hour * 24 * 100).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "Unauthorized", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	exp, ok := claims["ttl"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid expiration claim")
	}

	if int64(exp) < time.Now().Unix() {
		return 0, fmt.Errorf("token expired")
	}

	userIdInt, ok := claims["userId"].(float64)
	if !ok {
		return 0, fmt.Errorf("userId claim is not a float64")
	}

	return int32(userIdInt), nil
}

func setCookie(w http.ResponseWriter, tokenString string) {
	cookie := http.Cookie{
		Name:    "Auth",
		Value:   tokenString,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour * 100),
	}
	http.SetCookie(w, &cookie)

}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Config = struct {
	Context context.Context
	Conn    *pgx.Conn
}

func CheckEmailExists(email string, config Config) bool {
	queries := db.New(config.Conn)
	_, err := queries.GetUserByEmail(config.Context, email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Email does not exist
			return false
		}
		// Handle other errors
		return false
	}
	// Email exists
	return true
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user db.User

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if CheckEmailExists(user.Email, Config{Context: ctx, Conn: conn}) {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := HashPassword(user.Password.String)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := queries.CreateUserAndReturnId(ctx,
		db.CreateUserAndReturnIdParams{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  pgtype.Text{String: hashedPassword, Valid: true},
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := SignToken(entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	setCookie(w, token)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

type LoginUserType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user LoginUserType

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get request cookie
	authCookie, err := r.Cookie("Auth")
	if err != nil {
		http.Error(w, "Auth cookie not found", http.StatusUnauthorized)
		return
	}

	userId, err := VerifyToken(authCookie.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// If cookie is expired, create a new one
	if time.Now().UTC().After(authCookie.Expires.UTC()) {
		token, err := SignToken(userId)
		if err != nil {
			http.Error(w, "BLAH", http.StatusInternalServerError)
			return
		}
		setCookie(w, token)
	}

	userId, err = VerifyToken(authCookie.Value)
	if err != nil {
		http.Error(w, "Invalid token again", http.StatusUnauthorized)
		return
	}

	hashedPassword, err := queries.GetHashedPassword(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	passwordIsValid := VerifyPassword(user.Password, hashedPassword.String)
	if !passwordIsValid {
		http.Error(w, "Invalid request, please try again", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
