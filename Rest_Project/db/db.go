package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

type Manifest struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"dbname"`
}

var DB *sql.DB

func InitDB() {

	jsonFile, json_err := os.Open("manifest.json")
	if json_err != nil {
		fmt.Println(json_err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	var manifest Manifest
	json.Unmarshal(jsonBytes, &manifest)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", manifest.Host, manifest.Port, manifest.User, manifest.Password, manifest.DatabaseName)
	var err error
	DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	 id SERIAL PRIMARY KEY,
	 email TEXT NOT NULL UNIQUE,
	 password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    datetime TIMESTAMP NOT NULL,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id SERIAL PRIMARY KEY,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create registrations table")
	}
}
