package model

import (
)

type User struct {
	ID    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
