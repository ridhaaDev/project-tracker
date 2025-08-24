package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"project-tracker/api/db"
	"project-tracker/api/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

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

	token, err := utils.SignToken(entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SetCookie(w, token)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

type LoginUserType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	user, err := utils.ParseBody[LoginUserType](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get request cookie
	authCookie, err := r.Cookie("Auth")
	if err != nil {
		http.Error(w, "Auth cookie not found", http.StatusUnauthorized)
		return
	}

	// Initial check to verify the token
	userId, err := utils.VerifyToken(authCookie.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// If cookie is expired, create a new one
	if time.Now().UTC().After(authCookie.Expires.UTC()) {
		token, err := utils.SignToken(userId)
		if err != nil {
			http.Error(w, "Failed to sign token", http.StatusInternalServerError)
			return
		}
		utils.SetCookie(w, token)

		// Verify the new token
		userId, err = utils.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
