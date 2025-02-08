package repository

import (
	"fmt"
	"strings"

	"github.com/senyabanana/todo-app/internal/database"
	"github.com/senyabanana/todo-app/internal/entity"
	
	"github.com/jmoiron/sqlx"
)

type TodoListsPostgres struct {
	db *sqlx.DB
}

func NewTodoListsPostgres(db *sqlx.DB) *TodoListsPostgres {
	return &TodoListsPostgres{db: db}
}

func (r *TodoListsPostgres) Create(userId int, list entity.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", database.TodoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", database.UsersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}

func (r *TodoListsPostgres) GetAll(userId int) ([]entity.TodoList, error) {
	var lists []entity.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s AS ul on tl.id = ul.list_id WHERE ul.user_id=$1",
		database.TodoListsTable, database.UsersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListsPostgres) GetById(userId, listId int) (entity.TodoList, error) {
	var list entity.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s AS tl
		INNER JOIN %s AS ul on tl.id=ul.list_id WHERE ul.user_id=$1 AND ul.list_id=$2`,
		database.TodoListsTable, database.UsersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListsPostgres) Update(userId, listId int, input entity.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s AS tl SET %s FROM %s AS ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		database.TodoListsTable, setQuery, database.UsersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TodoListsPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s AS tl USING %s AS ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		database.TodoListsTable, database.UsersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
