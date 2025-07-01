package model

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}
	return &Note{title: title, content: content, createdAt: time.Now()}, nil
}

func (n *Note) String() string {
	return "Title: " + n.title + "\nContent: " + n.content + "\nCreated At: " + n.createdAt.Format(time.RFC3339)
}
