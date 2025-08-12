package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// type User struct {
// 	FirstName string `json:"name,omitempty"`
// 	LastName  string `json:"description,omitempty"`
// 	Username  string `json:"startDate,omitempty"`
// 	Password  string `json:"endDate,omitempty"`
// }

type Project struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
}

func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		w.Write([]byte("Something went wrong"))
	}

	fmt.Print(project)

	w.Write([]byte("welcome"))
}

func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// var project Project
	// err = db.Select(&project, "SELECT * FROM place ORDER BY telcode ASC")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	r.Get("/", getProject)
	r.Post("/", createProject)
	http.ListenAndServe(":3000", r)
}
