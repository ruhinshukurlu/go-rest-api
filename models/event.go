package models

import (
	"time"

	"github.com/go-rest-api/database"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	row := database.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	err := row.Scan(&e.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5
	`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	// row := database.DB.QueryRow(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES ($1, $2)"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) Cancel(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = $1 AND user_id = $2"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
