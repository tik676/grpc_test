package postgres

import (
	"database/sql"
	"grpc_test/internal/domain"
	"log"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

func (p *PostgresRepo) CreateTask(title, description string) domain.Task {
	var id int64
	query := `INSERT INTO todo(title,description,completed)VALUES($1,$2,$3)RETURNING id;`
	err := p.DB.QueryRow(query, title, description, false).Scan(&id)
	if err != nil {
		log.Println("error to create task", err)
		return domain.Task{}
	}

	return domain.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
	}
}

func (p *PostgresRepo) ListTasks() domain.TasksList {
	query := `SELECT * FROM todo;`
	rows, err := p.DB.Query(query)
	if err != nil {
		log.Println("error to send list task", err)
		return domain.TasksList{}
	}

	defer rows.Close()

	var response []domain.Task
	for rows.Next() {
		var list domain.Task

		err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Completed)
		if err != nil {
			log.Println("error to send list task")
			return domain.TasksList{}
		}

		response = append(response, list)
	}

	return domain.TasksList{Tasks: response}
}

func (p *PostgresRepo) EditTask(taskRequest domain.Task) domain.Task {
	query := `UPDATE todo SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING id, title, description, completed`
	var task domain.Task
	err := p.DB.QueryRow(query, taskRequest.Title, taskRequest.Description, taskRequest.Completed, taskRequest.ID).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		log.Println("error updating task:", err)
		return domain.Task{}
	}

	return domain.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
}

func (p *PostgresRepo) DeleteTask(id int) domain.Task {
	var task domain.Task
	query := `SELECT id,title,description,completed FROM todo`

	err := p.DB.QueryRow(query).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		log.Println("task not found for delete:", err)
		return domain.Task{}
	}

	queryDelete := `DELETE FROM todo WHERE id = $1`
	_, err = p.DB.Exec(queryDelete, task.ID)
	if err != nil {
		log.Println("error to delete task:", err)
		return domain.Task{}
	}

	return task
}
