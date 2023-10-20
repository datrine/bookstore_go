package utils

import "time"

type SanitizedRegisterInput struct {
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	DOB       *time.Time `json:"dob"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
}

type BookInput struct {
	Title         string `validate:"required"`
	ISBN          string `validate:"required"`
	YearOfPublish uint   `validate:"required"`
	AuthorId      uint   `validate:"required"`
}

type DeleteBookInput struct {
	DeleterId uint `validate:"required"`
	BookId    uint `validate:"required"`
}
