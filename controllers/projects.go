package controllers

import (
	"encoding/json"
	"net/http"
	"project-tracker/api/db" // adjust import path if needed
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var project db.Project

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := queries.CreateProject(ctx,
		db.CreateProjectParams{
			Name:        project.Name,
			Description: project.Description,
			StartDate:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	projects, err := queries.GetProjects(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}
	id := int32(idInt)

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	project, err := queries.GetProjectByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}
