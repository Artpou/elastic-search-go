package models

import (
	"time"
)

type Book struct {
	ID          uint
	Author      Author
	Title       string
	Preface     string
	Category    string
	ReleaseDate time.Time
}
