package handlers

import (
	"errors"
	"slices"
	"task_traker/internal/database"
	"task_traker/internal/models"
	"time"
)

type ListOperations interface {
	Add() (int8, error)
	Delete(id int16) (bool, error)
	Update(id int16, name string) (bool, error)

	UpdateStatus(id int16, status models.Status) (bool, error)

	GetList(status *models.Status) ([]models.Todo, error)
}

type List struct{}

func (reader List) Add(todoName string) (int, error) {
	var id int = 1

	db := database.Reader{
		Name: "todos.json",
	}

	todoList := make([]models.Todo, 0)

	rError := db.Read(&todoList)

	if rError == nil {
		id = todoList[len(todoList)-1].Id + 1
	}

	currentTime := time.Now()

	todo := models.Todo{
		Id:        id,
		Name:      todoName,
		Status:    models.CREATED,
		CreatedAt: currentTime.Format("2006-01-2 15:4:5"),
		UpdatedAt: "",
	}

	todoList = append(todoList, todo)

	_, err := db.Write(todoList)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (reader List) Delete(id int) (bool, error) {
	db := database.Reader{
		Name: "todos.json",
	}

	todoList := make([]models.Todo, 0)

	rError := db.Read(&todoList)

	if rError != nil {
		return false, rError
	}

	idx := slices.IndexFunc(todoList, func(todo models.Todo) bool {
		return todo.Id == id
	})

	if idx < 0 {
		return false, errors.New("the todo do not exist")
	}

	firstPart := todoList[0:idx]
	var secondPart []models.Todo

	if idx+1 < len(todoList) {
		secondPart = todoList[idx+1:]
	}

	newTodoList := append(firstPart, secondPart...)

	return db.Write(newTodoList)
}

func (reader List) Update(id int16, name string) (bool, error) {
	db := database.Reader{
		Name: "todos.json",
	}

	todoList := make([]models.Todo, 0)

	rError := db.Read(&todoList)

	if rError != nil {
		return false, rError
	}

	idx := slices.IndexFunc(todoList, func(todo models.Todo) bool {
		return todo.Id == int(id)
	})

	if idx < 0 {
		return false, errors.New("the todo do not exist")
	}

	currentTime := time.Now()

	todoList[idx].Name = name
	todoList[idx].UpdatedAt = currentTime.Format("2006-01-2 15:4:5")

	_, err := db.Write(todoList)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (reader List) UpdateStatus(id int16, status models.Status) (bool, error) {
	db := database.Reader{
		Name: "todos.json",
	}

	todoList := make([]models.Todo, 0)

	rError := db.Read(&todoList)

	if rError != nil {
		return false, rError
	}

	idx := slices.IndexFunc(todoList, func(todo models.Todo) bool {
		return todo.Id == int(id)
	})

	if idx < 0 {
		return false, errors.New("the todo do not exist")
	}

	currentTime := time.Now()

	todoList[idx].Status = status
	todoList[idx].UpdatedAt = currentTime.Format("2006-01-2 15:4:5")

	_, err := db.Write(todoList)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (reader List) GetList(status models.Status) ([]models.Todo, error) {
	db := database.Reader{
		Name: "todos.json",
	}

	todoList := make([]models.Todo, 0)

	rError := db.Read(&todoList)

	if rError != nil || status == "" {
		return todoList, rError
	}

	cleaned := make([]models.Todo, 0)

	for _, todo := range todoList {
		if todo.Status == status {
			cleaned = append(cleaned, todo)
		}
	}

	return cleaned, nil
}
