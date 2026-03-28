package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./blog.db"
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbPath))
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// restrict file permissions to owner-only (600)
	if err := os.Chmod(dbPath, 0600); err != nil {
		log.Println("Warning: could not set DB file permissions:", err)
	}

	schema, err := os.ReadFile("seed.sql")
	if err != nil {
		log.Fatal("Failed to read seed.sql:", err)
	}

	// execute schema
	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal("Failed to execute schema:", err)
	}

	return db
}
