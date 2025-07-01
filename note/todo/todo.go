package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo *Todo) Display() string {
	return todo.Text
}

func (todo *Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid input: text cannot be empty")
	}

	return &Todo{Text: text}, nil
}
