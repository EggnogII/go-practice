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
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// Use $n as a safe way to inject values
	query := `INSERT INTO events(name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5)`
	statement, statement_err := db.DB.Prepare(query)
	if statement_err != nil {
		return statement_err
	}
	defer statement.Close()
	result, result_err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if result_err != nil {
		return result_err
	}
	id, id_err := result.LastInsertId()
	e.ID = id
	return id_err
}

func GetAllEvents() []Event {
	return events
}
