package repository

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name string `json:"name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Todo
}
