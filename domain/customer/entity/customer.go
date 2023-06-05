package entity

type Customer struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Email       string  `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number,omitempty" db:"phone_number"`
}
