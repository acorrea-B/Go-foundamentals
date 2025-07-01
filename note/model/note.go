package model

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	return &Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}

func (n *Note) String() string {
	return "Title: " + n.Title + "\nContent: " + n.Content + "\nCreated At: " + n.CreatedAt.Format(time.RFC3339)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}
