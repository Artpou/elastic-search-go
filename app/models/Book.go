package models

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	Title    string `gorm:"not null;size:500"`
	Author   string `gorm:"not null;size:100"`
	Abstract string `gorm:"not null;size:5000"`
	Year     string `gorm:"not null;size:4"`
}

func NewBook() {

}
