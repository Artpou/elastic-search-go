package models

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	Title  string `gorm:"not null;size:500"`
	Author string `gorm:"not null;size:500"`
	Year   string `gorm:"not null;size:20"`
}
