package entity

import "go.mod/domain/enum"

type Employer struct {
	ID          string      `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Office      enum.Office `json:"office" db:"office"`
	PhoneNumber string      `json:"phone_number" db:"phone_number"`
}
