package service

import (
	"todoAppWithTest/model"
	"todoAppWithTest/storage"
)

func CreateTodoService(todo model.Todo) model.Todo {
	return storage.CreateTodo(todo)
}

func GetTodosService() []model.Todo {
	return storage.GetTodos()
}

func UpdateTodoService(id int, todo model.Todo) (model.Todo, error) {
	return storage.UpdateTodoById(id, todo)
}

func DeleteTodoService(id int) error {
	return storage.DeleteTodoById(id)
}
