package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DataSource struct {
	db *sql.DB
}

func (d *DataSource) InitDB(dbName string) error {
	var err error
	d.db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}

	q := `
		CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			status TEXT
		)
	`
	_, err = d.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataSource) FetchTodos() ([]Todo, error) {
	q := `SELECT id, title, status FROM todos`
	rows, err := d.db.Query(q)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (d *DataSource) AddTodo(todo Todo) error {
	q := `
	INSERT INTO todos(title, status)
	VALUES(?, ?)
	`
	_, err := d.db.Exec(q, todo.Title, todo.Status)
	return err
}

func (d *DataSource) UpdateTodo(todo Todo) error {
	q := `
	UPDATE todos 
	SET title=?, status=?
	WHERE id=?
	`
	_, err := d.db.Exec(q, todo.Title, todo.Status, todo.ID)
	return err
}
