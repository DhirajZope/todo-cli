package db

import (
	"database/sql"
	"fmt"
	"time"
)

func Init(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", path+"?_foreign_keys=on")
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE IF NOT EXISTS lists (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      created_at DATETIME NOT NULL
    );
    CREATE TABLE IF NOT EXISTS tasks (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      description TEXT NOT NULL,
      completed BOOLEAN NOT NULL DEFAULT 0,
      created_at DATETIME NOT NULL,
      completed_at DATETIME,
      list_id INTEGER, 
      FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
    );
  `

	if _, err := conn.Exec(schema); err != nil {
		return nil, fmt.Errorf("failed to init the schema: %w", err)
	}

	return conn, nil
}

func now() time.Time {
	return time.Now().UTC()
}
