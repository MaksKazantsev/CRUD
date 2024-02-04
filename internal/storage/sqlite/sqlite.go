package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/go-crud/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func GetDB(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
CREATE TABLE IF NOT EXISTS url(
    id INTEGER PRIMARY KEY,
    alias TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_alias ON url(alias)
)
`)
	if err != nil {
		return nil, fmt.Errorf("#{op}: #{err}")
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("#{op}: #{err}")
	}
	return &Storage{db: db}, nil
}

func (s *Storage) saveURL(URLtoSave string, alias string) (int64, error) {
	const op = "storage.sqlite.New"
	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?,?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(URLtoSave, alias)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) getURL(alias string) (string, error) {
	const op = "storage.sqlite.New"

	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var resURL string
	err = stmt.QueryRow(alias).Scan(&resURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound
	}
	return "", fmt.Errorf("%s: %w", op, err)
}
