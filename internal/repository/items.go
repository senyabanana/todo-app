package repository

import (
	"fmt"
	"strings"

	"github.com/senyabanana/todo-app/internal/database"
	"github.com/senyabanana/todo-app/internal/entity"
	
	"github.com/jmoiron/sqlx"
)

type TodoItemsRepository struct {
	db *sqlx.DB
}

func NewTodoItemsRepository(db *sqlx.DB) *TodoItemsRepository {
	return &TodoItemsRepository{db: db}
}

func (r *TodoItemsRepository) Create(listId int, input entity.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", database.TodoItemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", database.ListsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemsRepository) GetAll(userId, listId int) ([]entity.TodoItem, error) {
	var items []entity.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s AS ti INNER JOIN %s AS li ON ti.id = li.item_id
									INNER JOIN %s AS ul ON li.list_id = ul.list_id WHERE li.list_id=$1 AND ul.user_id=$2`,
		database.TodoItemsTable, database.ListsItemsTable, database.UsersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemsRepository) GetById(userId, itemId int) (entity.TodoItem, error) {
	var item entity.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s AS ti INNER JOIN %s AS li ON ti.id = li.item_id
									INNER JOIN %s AS ul ON li.list_id = ul.list_id WHERE ti.id=$1 AND ul.user_id=$2`,
		database.TodoItemsTable, database.ListsItemsTable, database.UsersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemsRepository) Update(userId, itemId int, input entity.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s AS ti SET %s FROM %s AS li, %s AS ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		database.TodoItemsTable, setQuery, database.ListsItemsTable, database.UsersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TodoItemsRepository) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s AS ti USING %s AS li, %s AS ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$1 AND ti.id=$2`,
		database.TodoItemsTable, database.ListsItemsTable, database.UsersListsTable)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}
