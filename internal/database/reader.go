package database

import (
	"encoding/json"
	"os"
)

type ReaderOperations interface {
	Read(v any) error
	Write(v any) (bool, error)
}

type Reader struct {
	Name string
}

func (reader Reader) Read(v any) error {
	rawList, err := os.ReadFile(reader.Name)

	if (err != nil) {
		return err
	}

	return json.Unmarshal(rawList, v)
}

func (reader Reader) Write(v any) (bool, error) {
	bytes, bError := json.Marshal(v)

	if bError != nil {
		return false, bError
	}

	writeError := os.WriteFile("todos.json", bytes, 0666)

	if writeError != nil {
		return false, writeError
	}	

	return true, nil
}