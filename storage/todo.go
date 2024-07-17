package storage

import (
	"errors"
	"todoAppWithTest/model"
)

var todos = []model.Todo{}
var idCounter = 1

func CreateTodo(todo model.Todo) model.Todo {
	todo.ID = idCounter
	idCounter++
	todos = append(todos, todo)
	return todo
}

func GetTodos() []model.Todo {
	return todos
}

func UpdateTodoById(id int, updatedTodo model.Todo) (model.Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Completed = updatedTodo.Completed
			return todos[i], nil
		}
	}
	return model.Todo{}, errors.New("Todo not found")
}

func DeleteTodoById(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}
