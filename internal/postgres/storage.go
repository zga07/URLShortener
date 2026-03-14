package postgres

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    alias TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Не удалось создать/найти таблицу:", err)
	}
}
