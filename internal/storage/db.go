package storage

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func getDBPath() (string, error) {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("invalid database: unable to find home directory")
	}

	dir := filepath.Join(home, ".local", "state", "tsk")

	// Auto-create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}

	return filepath.Join(dir, "tsk.db"), nil
}

func connect() (*sql.DB, error) {
	path, err := getDBPath()
	if err != nil {
		return nil, errors.New("Invalid database path")
	}

	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}

	query := `
	    CREATE TABLE IF NOT EXISTS tasks (
	        id TEXT PRIMARY KEY,
	        template_id INTEGER,
	        data TEXT NOT NULL
	    );

	    CREATE TABLE IF NOT EXISTS templates (
	        id TEXT PRIMARY KEY,
	        data TEXT NOT NULL
	    );

	    CREATE TABLE IF NOT EXISTS assignments (
	        id INTEGER PRIMARY KEY,
	        task_id TEXT NOT NULL
	    );
	`

	_, err = conn.Exec(query)
	if err != nil {
		return nil, errors.New("Migration failed")
	}

	return conn, nil
}
