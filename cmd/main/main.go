package main

import (
	"log"
	"net/http"
	"os"
	"project-tracker/api/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func getEndpointName(prefix string, endpointName string) string {
	return prefix + endpointName
}

func main() {

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"}, // Specify allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware.Handler)

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiPrefix := "/api/" + os.Getenv("API_VERSION")

	r.Post(getEndpointName(apiPrefix, "/user/signup"), controllers.SignupUser)
	r.Post(getEndpointName(apiPrefix, "/user/login"), controllers.LoginUser)

	r.Post(getEndpointName(apiPrefix, "/projects/create"), controllers.CreateProject)
	r.Get(getEndpointName(apiPrefix, "/projects"), controllers.GetProjects)
	r.Get(getEndpointName(apiPrefix, "/projects/{id}"), controllers.GetProjectByID)
	r.Get(getEndpointName(apiPrefix, "/projects/{id}/tickets"), controllers.GetProjectTickets)
	r.Post(getEndpointName(apiPrefix, "/projects/{id}/tickets/create"), controllers.CreateTicket)

	http.ListenAndServe(":3000", r)
}
