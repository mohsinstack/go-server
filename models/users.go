package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Firstname  string      `json:"firstname"`
	Lastname   string      `json:"lastname"`
	Email      string      `json:"email" gorm:"unique"`
	Educations []Education `json:"educations" gorm:"foreignKey:UserID"`
}
