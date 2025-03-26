package models

type Status string

const (
	CREATED Status = "created"
	IN_PROGRESS Status = "in_progress"
	DONE Status = "done"
)

type Todo struct {
	Id int `json:"id"`
	Name string	`json:"name"`
	Status Status `json:"status"`
}