package services

import (
	"time"

	"github.com/datrine/models"
	"github.com/datrine/repositories"
	"github.com/datrine/utils"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func AddBookService(data *utils.BookInput) (*Book, error) {
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	bookModel, err := repositories.AddBook(data)
	if err != nil {
		return nil, err
	}
	book := &Book{
		ID:            bookModel.ID,
		Title:         bookModel.Title,
		AuthorId:      bookModel.AuthorId,
		ISBN:          bookModel.ISBN,
		CreatedAt:     bookModel.CreatedAt,
		UpdatedAt:     bookModel.UpdatedAt,
		YearOfPublish: bookModel.YearOfPublish,
	}
	return book, nil
}

func GetBooksService(data *utils.GetBooksDTO) (*[]*Book, error) {
	var books []*Book = []*Book{}
	bookModels := &[]models.Book{}
	bookModels, err := repositories.GetBooks(data)
	if err != nil {
		return nil, err
	}
	for _, v := range *bookModels {
		book := &Book{
			ID:            v.ID,
			Title:         v.Title,
			AuthorId:      v.AuthorId,
			ISBN:          v.ISBN,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			YearOfPublish: v.YearOfPublish,
		}
		books = append(books, book)
	}
	return &books, nil
}

func GetBookById(bookId uint) (*Book, error) {
	bookModel, err := repositories.GetBookById(bookId)
	if err != nil {
		return nil, err
	}
	book := &Book{
		ID:            bookModel.ID,
		YearOfPublish: bookModel.YearOfPublish,
		AuthorId:      bookModel.AuthorId,
		ISBN:          bookModel.ISBN,
		Title:         bookModel.Title,
		CreatedAt:     bookModel.CreatedAt,
		UpdatedAt:     bookModel.UpdatedAt,
	}
	return book, nil
}

func DeleteBookByOwner(data *utils.DeleteBookByOwnerDTO) (*Book, error) {
	bookModel, err := repositories.DeleteBookByOwner(&utils.DeleteBookInput{
		DeleterId: data.AuthorId,
		BookId:    data.BookId,
	})
	if err != nil {
		return nil, err
	}
	book := &Book{
		ID:            bookModel.ID,
		YearOfPublish: bookModel.YearOfPublish,
		AuthorId:      bookModel.AuthorId,
		ISBN:          bookModel.ISBN,
		Title:         bookModel.Title,
		CreatedAt:     bookModel.CreatedAt,
		UpdatedAt:     bookModel.UpdatedAt,
	}
	return book, nil
}

type Book struct {
	ID            uint
	Title         string
	ISBN          string
	AuthorId      uint
	YearOfPublish int16
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
