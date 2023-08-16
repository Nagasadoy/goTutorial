package main

import "fmt"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	NumChild int    `json:"-"`
	Age      int    `json:"age"`
}

func NewUser(name, email string, age int) (*User, error) {

	if age < 0 {
		return nil, fmt.Errorf("возраст не может быть отрицательным")
	}

	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}, nil
}
