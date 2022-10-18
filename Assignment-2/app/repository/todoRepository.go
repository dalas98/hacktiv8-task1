package repository

import (
	"dalas98/golang-lesson/Assignment-2/app"
	"dalas98/golang-lesson/Assignment-2/app/model"
	"database/sql"
	"errors"
)

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) app.ITodoRepository {
	return &TodoRepo{
		db: db,
	}
}

func (t *TodoRepo) CreateTask(todo *model.Todo) error {
	query := `
	INSERT INTO todos (
		task, status
	) VALUES (?, ?)
	`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Task, todo.Status)

	return err
}

func (t *TodoRepo) GetTask() (*[]model.Todo, error) {
	query := `
	SELECT id, task, status, created_at, updated_at FROM todos
	`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return &todos, nil
}
func (t *TodoRepo) UpdateTask(id string, todo *model.Todo) error {

	data, err := t.getById(id)
	if err != nil {
		return err
	}

	if len(data.Task) < 1 {
		return errors.New("data not found")
	}

	query := `
	UPDATE todos 
	SET task = ?, status = ?
	WHERE id = ?
	`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Task, todo.Status, id)

	return err
}

func (t *TodoRepo) getById(id string) (*model.Todo, error) {
	query := `
	SELECT id, task, status, created_at, updated_at FROM todos
	WHERE id = ?
	`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todo model.Todo
	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &todo, nil
}

func (t *TodoRepo) DeleteTask(id string) error {
	query := `
	DELETE FROM todos WHERE id = ?
	`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
