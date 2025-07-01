package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s %s, Birthdate: %s, Created At: %s", u.firstName, u.lastName, u.birthdate, u.createdAt.Format(time.RFC3339))
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	var appUser *User
	appUser = newUser("John", "Doe", "1990-01-01")

	fmt.Println(appUser.String())
	appUser.clearUserName()
	fmt.Println(appUser.String())
}
