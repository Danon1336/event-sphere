package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/argon2"
)

var DB *sql.DB

func InitDB() {
	var err error
	if _, err := os.Stat("users.db"); os.IsNotExist(err) {
		file, err := os.Create("users.db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	DB, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			surname TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL,
			salt TEXT NOT NULL
		);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return string(hash)
}
