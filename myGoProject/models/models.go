package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Director    string  `json:"director" gorm:"not null"`
	ReleaseYear int     `json:"release_year" gorm:"not null"`
	Rating      float32 `json:"rating" gorm:"not null"`
}
