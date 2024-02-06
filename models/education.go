package models

import (
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	InstitutionName string  `json:"institutionName"`
	PassoutYear     int     `json:"passoutYear"`
	CGPI            float64 `json:"cgpi"`
	UserID          uint    `json:"userId"`
}
