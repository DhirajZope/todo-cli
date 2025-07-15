package service

import (
	"database/sql"
	"time"

	"github.com/dhirajzope/todo/internal/model"
)

type ListService struct {
	DB *sql.DB
}

func NewListService(db *sql.DB) *ListService {
	return &ListService{DB: db}
}

func (s *ListService) Create(name string) error {
	_, err := s.DB.Exec(
		`INSERT INTO lists (name, created_at) VALUES (?, ?)`,
		name, time.Now().UTC(),
	)
	return err
}

func (s *ListService) Delete(id int) error {
	_, err := s.DB.Exec(`DELETE FROM lists WHERE id = ?`, id)
	return err
}

func (s *ListService) All() ([]model.List, error) {
	rows, err := s.DB.Query(`SELECT id, name, created_at FROM lists`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lists []model.List

	for rows.Next() {
		var l model.List

		if err := rows.Scan(&l.ID, &l.Name, &l.CreatedAt); err != nil {
			return nil, err
		}

		lists = append(lists, l)
	}

	return lists, nil
}
