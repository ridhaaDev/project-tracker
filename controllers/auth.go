package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30 // 30 days in seconds
	IsProd = false
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("CLIENT_ID")
	googleClientSecret := os.Getenv("CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(google.New(googleClientId, googleClientSecret, os.Getenv("REDIRECT_URL")))
}

func getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, "Failed to complete user auth", http.StatusInternalServerError)
		return
	}

	// Here you can handle the user object as needed
	log.Printf("User: %+v\n", user)

	// Redirect or respond as needed
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		http.Error(w, "Provider not specified", http.StatusBadRequest)
		return
	}

	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getUserWithoutAuth(w http.ResponseWriter, r *http.Request) (goth.User, error) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		return gothUser, nil
	} else {
		gothic.BeginAuthHandler(w, r)
		return goth.User{}, err
	}
}
