// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	User   *User  `json:"user"`
	Active bool   `json:"active"`
}

type NewAccount struct {
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts"`
}
