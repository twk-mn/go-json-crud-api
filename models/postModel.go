package models

import "gorm.io/gorm"

// Basic structure for DB table
type Post struct {
	gorm.Model
	Title string
	Body  string
}
