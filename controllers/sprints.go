package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project-tracker/api/db"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var ticket db.Ticket

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := queries.CreateTicket(ctx,
		db.CreateTicketParams{
			SprintID:    ticket.SprintID,
			Name:        ticket.Name,
			Description: ticket.Description,
			Status:      ticket.Status,
			StartDate:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func GetProjectTickets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	projectID := chi.URLParam(r, "id")

	fmt.Println("Project ID:", projectID)

	projectIDInt, err := strconv.ParseInt(projectID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pr := pgtype.Int4{Int32: int32(projectIDInt), Valid: true}

	conn, ctx := db.ConnectDB()
	defer conn.Close(ctx)

	queries := db.New(conn)

	tickets, err := queries.GetTicketsByProjectID(ctx, pr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tickets)
}
