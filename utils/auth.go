package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func SetCookie(w http.ResponseWriter, tokenString string) {
	cookie := http.Cookie{
		Name:    "Auth",
		Value:   tokenString,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour * 100),
	}
	http.SetCookie(w, &cookie)

}
