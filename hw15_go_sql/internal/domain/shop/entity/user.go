package entity

import "errors"

type User struct {
	ID       int32  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserUpdateParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type UserCreateParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Params struct {
	Limit  int64
	Offset int64
}

var ErrUserNotFound = errors.New("user not found")
