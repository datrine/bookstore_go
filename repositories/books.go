package repositories

import (
	"fmt"

	"github.com/datrine/conn"
	"github.com/datrine/models"
	"github.com/datrine/utils"
)

func AddBook(bookInput *utils.BookInput) (*models.Book, error) {
	db := conn.DB
	book := &models.Book{
		YearOfPublish: int16(bookInput.YearOfPublish),
		AuthorId:      bookInput.AuthorId,
		ISBN:          bookInput.ISBN,
		Title:         bookInput.Title,
	}
	result := db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func GetBooks(data *utils.GetBooksDTO) (*[]models.Book, error) {
	db := conn.DB
	bookModels := []models.Book{}
	result := db.Find(&bookModels)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, result.Error
	}
	return &bookModels, nil
}

func GetBookById(bookId uint) (*models.Book, error) {
	db := conn.DB
	book := &models.Book{
		ID: bookId,
	}
	tx := db.First(book)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBookByOwner(data *utils.DeleteBookInput) (*models.Book, error) {
	db := conn.DB
	book := models.Book{
		AdderId: data.DeleterId,
		ID:      data.BookId,
	}
	result := db.Unscoped().Delete(&book)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, result.Error
	}
	return &book, nil
}
