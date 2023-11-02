package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string
	Position   string
	Department string
	Email      string
}
