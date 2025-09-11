package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin123"
	dbname   = "go-db"
)

func InitDB() {

	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		panic("Could not connect to database.")
	}

	err = DB.Ping()
	if err != nil {
		panic("Could not ping to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Couldn't create users table..")
	}

	createEventTables := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime TIMESTAMP NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventTables)

	if err != nil {
		panic("Couldn't create events table")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations(
			id SERIAL PRIMARY KEY,
			user_id INTEGER,
			event_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(event_id) REFERENCES events(id)
		)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Couldn't create registrations table")
	}
}
