package db

import "database/sql"

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	Location    string `json:"location,omitempty"`
}

func GetEvents(db *sql.DB) ([]Event, error) {
	var rows *sql.Rows
	var err error

	rows, err  = db.Query(`
		SELECT id, title, description, event_date, location 
		FROM events 
		WHERE show = TRUE
		ORDER BY event_date ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.EventDate, &event.Location); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventsByID(db *sql.DB, id int) (*Event, error) {
	var event Event
	err := db.QueryRow(`
		SELECT id, title, description, event_date, location
		FROM events
		WHERE id = $1 AND show = TRUE
	`, id).Scan(&event.ID, &event.Title, &event.Description, &event.EventDate, &event.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func CreateEvent(db *sql.DB, title, description, eventDate, location string) (*Event, error) {
	var event Event
	err := db.QueryRow(`
		INSERT INTO events (title, description, event_date, location)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, description, event_date, location
	`, title, description, eventDate, location).Scan(&event.ID, &event.Title, &event.Description, &event.EventDate, &event.Location)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
