package main

import (
	"log"
	"net/http"
	"os"
	"project-tracker/api/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func getEndpointName(prefix string, endpointName string) string {
	return prefix + endpointName
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPrefix := "/api/" + os.Getenv("API_VERSION")
	r.Post(getEndpointName(apiPrefix, "/project/create"), controllers.CreateProject)

	http.ListenAndServe(":3000", r)
}
