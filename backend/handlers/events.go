package handlers

import (
	"database/sql"
	"encoding/json"
	"event-guest-manager/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EventsHandler struct {
	db *sql.DB
}

type CreateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	Location    string `json:"location,omitempty"`
}

func NewEventsHandler(database *sql.DB) *EventsHandler {
	return &EventsHandler{db: database}
}

func (h *EventsHandler) GetEvents(w http.ResponseWriter,r *http.Request) {
	events, err := db.GetEvents(h.db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch events")
		return
	}
	respondWithJSON(w, http.StatusOK, events)
}

func (h *EventsHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid event ID")
		return
	}
	event, err := db.GetEventsByID(h.db, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch event")
		return
	}
	respondWithJSON(w, http.StatusOK, event)
}

func (h *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	event, err := db.CreateEvent(h.db, req.Title, req.Description, req.EventDate, req.Location)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create event")
		return
	}
	respondWithJSON(w, http.StatusCreated, event)
}
