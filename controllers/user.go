package controllers

import (
	"encoding/json"
	"net/http"
	"project-tracker/api/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user db.User

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := queries.CreateUser(ctx,
		db.CreateUserParams{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
