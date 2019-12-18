package infra

import (
	"fmt"
	"todo/domain/model"
	"todo/domain/repository"
)

// TodoRepository ...
type TodoRepository struct {
	SQLHandler
}

// NewTodoRepository ...
func NewTodoRepository(sqlHandler SQLHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}

// FindAll ...
func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
	rows, err := todoRepo.SQLHandler.Conn.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		todo := model.Todo{}

		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

		todos = append(todos, &todo)
	}
	return
}

// Find ...
func (todoRepo *TodoRepository) Find(word string) (todos []*model.Todo, err error) {
	rows, err := todoRepo.SQLHandler.Conn.Query("SELECT * FROM todos WHERE task LIKE ?", "%"+word+"%")
	defer rows.Close()
	if err != nil {
		fmt.Print(err)
		return
	}
	for rows.Next() {
		todo := model.Todo{}

		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

		todos = append(todos, &todo)
	}
	return
}

// Create ...
func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SQLHandler.Conn.Exec("INSERT INTO todos (task,limitDate,status) VALUES (?, ?, ?) ", todo.Task, todo.LimitDate, todo.Status)
	return todo, err
}

// Update ...
func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.SQLHandler.Conn.Exec("UPDATE todos SET task = ?,limitDate = ? ,status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.ID)
	return todo, err
}
