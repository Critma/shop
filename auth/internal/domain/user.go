package domain

import "time"

type Password struct {
	Value string
	Hash  []byte
}

type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Phone    string    `json:"phone"`
	Password Password  `json:"-"`
	Updated  time.Time `json:"updated"`
	Created  time.Time `json:"created"`
}

func NewUser(email, name, surname, phone, password string) *User {
	user := &User{
		Email:   email,
		Name:    name,
		Surname: surname,
		Phone:   phone,
		Password: Password{
			Value: password,
		},
	}
	return user
}
