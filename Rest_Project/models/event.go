package models

import (
	"time"

	"example.com/rest-project/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	// Use $n as a safe way to inject values
	query := `INSERT INTO events(name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`
	statement, statement_err := db.DB.Prepare(query)
	if statement_err != nil {
		return statement_err
	}
	defer statement.Close()

	var id_num int
	result_err := statement.QueryRow(e.Name, e.Description, e.Location, e.DateTime, 1).Scan(&id_num)
	if result_err != nil {
		return result_err
	}

	e.ID = int64(id_num)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, rows_err := db.DB.Query(query)
	if rows_err != nil {
		return nil, rows_err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		scan_err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if scan_err != nil {
			return nil, scan_err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)
	var event Event
	scan_err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if scan_err != nil {
		return nil, scan_err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = $2, description = $3, location = $4, datetime = $5
	WHERE id = $1
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES ($1, $2)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err

}
