package models

import "time"

type Student struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	BirthDate   time.Time `json:"birth_date"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}
