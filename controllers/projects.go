package controllers

import (
	"encoding/json"
	"net/http"
	"project-tracker/api/db" // adjust import path if needed
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
			StartDate:   project.StartDate,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
