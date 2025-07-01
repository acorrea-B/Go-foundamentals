package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"explame.com/note/model"
	"explame.com/note/todo"
)

func main() {
	note, err := getNoteData()
	todoText := getUserImput("Todo text:")
	todo, tod_err := todo.New(todoText)

	if err != nil || tod_err != nil {
		fmt.Println("Error:", err, tod_err)
		return
	}
	fmt.Println("Note created successfully!")
	fmt.Println(note.String())
	fmt.Println("Todo created successfully!")
	fmt.Println(todo.Display())
	err = note.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved succeeded!")
	err = todo.Save()
	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	fmt.Println("Todo saved succeeded!")

}

func getNoteData() (*model.Note, error) {
	title := getUserImput("Note title:")
	content := getUserImput("Note content:")
	var note *model.Note
	note, err := model.New(title, content)
	return note, err
}

func getUserImput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
