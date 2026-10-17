package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTables()
	seedStudents()
}

func createTables() {
	schema := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		grade TEXT NOT NULL
	);`

	if _, err := DB.Exec(schema); err != nil {
		log.Fatal(err)
	}
}

func seedStudents() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM students").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		return
	}

	students := []struct {
		Name  string
		Email string
		Grade string
	}{
		{"Alice Johnson", "alice@example.com", "A"},
		{"Bob Smith", "bob@example.com", "B"},
		{"Carol White", "carol@example.com", "A"},
		{"David Brown", "david@example.com", "C"},
		{"Eve Davis", "eve@example.com", "B"},
		{"Frank Miller", "frank@example.com", "A"},
		{"Grace Wilson", "grace@example.com", "B"},
		{"Henry Moore", "henry@example.com", "C"},
		{"Ivy Taylor", "ivy@example.com", "A"},
		{"Jack Anderson", "jack@example.com", "B"},
	}

	for _, s := range students {
		_, err := DB.Exec(
			"INSERT INTO students (name, email, grade) VALUES (?, ?, ?)",
			s.Name, s.Email, s.Grade,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
