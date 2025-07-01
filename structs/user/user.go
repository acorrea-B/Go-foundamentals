package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // Embedding User struct
}

// Method for User struct, prints the user details
func (u *User) String() string {
	return fmt.Sprintf("Name: %s %s, Birthdate: %s, Created At: %s", u.firstName, u.lastName, u.birthdate, u.createdAt.Format(time.RFC3339))
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// New is a constructor function that initializes a User instance.
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields must be provided")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password, firstName, lastName, birthdate string) *Admin {

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthdate: birthdate,
			createdAt: time.Now(),
		},
	}
}
