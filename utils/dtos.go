package utils

import "time"

type RegisterDto struct {
	Email       string     `json:"email" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	DOB         *time.Time `json:"dob"`
	FirstName   string     `json:"firstname"`
	LastName    string     `json:"lastname"`
	PhoneNumber string     `json:"phone_number"`
}

type BasicLoginDto struct {
	Email    string
	Password string
}

type AddBookDTO struct {
	Title         string `json:"title" binding:"required"`
	ISBN          string `json:"isbn" binding:"required"`
	YearOfPublish uint   `json:"year_of_publish" binding:"required"`
	AuthorId      uint
}

type GetBooksDTO struct {
	Filters BookFiltersDTO
}

type BookFiltersDTO struct {
	Title      *string
	Like_Title *string
}

type DeleteBookByOwnerDTO struct {
	AuthorId uint `json:"author_id" binding:"required"`
	BookId   uint `json:"book_id" binding:"required"`
}
