package service

import (
	"database/sql"
	"time"

	"github.com/dhirajzope/todo/internal/model"
)

type TaskService struct {
	DB *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) Add(listID int, desc string) error {
	_, err := s.DB.Exec(`INSERT INTO tasks (list_id, description, created_at) VALUES (?, ?, ?)`, listID, desc, time.Now().UTC())
	return err
}

func (s *TaskService) Delete(id int) error {
	_, err := s.DB.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	return err
}

func (s *TaskService) Complete(id int) error {
	_, err := s.DB.Exec(`UPDATE tasks SET completed = 1, completed_at = ? WHERE id = ?`, time.Now().UTC(), id)
	return err
}

func (s *TaskService) ListBy(listID int) ([]model.Task, error) {
	rows, err := s.DB.Query(`SELECT id, list_id, description, completed, created_at, completed_at FROM tasks WHERE list_id = ?`, listID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		var nt sql.NullTime

		if err := rows.Scan(&t.ID, &t.ListID, &t.Description, &t.Completed, &t.CreatedAt, &nt); err != nil {
			return nil, err
		}

		if nt.Valid {
			t.CompletedAt = &nt.Time
		} else {
			t.CompletedAt = nil
		}

		tasks = append(tasks, t)
	}
	return tasks, nil
}
