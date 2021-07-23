package models

import (
	"time"
)

type Author struct {
	ID          uint
	FirstName   string
	LastName    string
	Nationality string
	BirthDate   time.Time
}
