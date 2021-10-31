package model

import "github.com/google/uuid"

// User is a canonical user model
type User struct {
	Base
	ID       int       `json:"id" pg:"id,notnull,pk"`
	UUID     uuid.UUID `json:"uuid" pg:"uuid,unique,notnull"`
	Login    string    `json:"login" validate:"required" pg:"login,unique,notnull"`
	Password string    `json:"-" validate:"required" pg:"password"`
}

//UsersList is a list of users
type UsersList struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

type UserRegisterForm struct {
	Login    string `json:"login" validate:"required" pg:"login,unique,notnull"`
	Password string `json:"password" validate:"required" pg:"password,notnull"`
}
