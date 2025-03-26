package handlers

import (
	"task_traker/internal/database"
	"task_traker/internal/models"
)

type ListOperations interface {
	Add() (int8, error)
	Delete() (bool, error)
}

type List struct {}

func (reader List) Add(todoName string) (int, error) {
	var id int = 1

	db := database.Reader {
		Name: "todos.json",
	}

	todoList := make([]models.Todo,0)

	rError := db.Read(&todoList)

	if (rError == nil) {
		id = len(todoList) + 1
	}

	todo := models.Todo{
		Id: id,
		Name: todoName,
		Status: models.CREATED,
	}

	todoList = append(todoList, todo)

	_, err := db.Write(todoList)

	if (err != nil) {
		return -1, err
	}

	return id, nil
}

func (reader List) Delete(id int) (bool, error) {
	return false, nil
}