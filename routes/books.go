package routes

import (
	"fmt"
	"strconv"

	"github.com/datrine/services"
	"github.com/datrine/utils"
	"github.com/gin-gonic/gin"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} utils.OkResponse
// @Failure 400 {object} utils.ErrResponse
// @Router /books [post]
func addBook(c *gin.Context) {
	data := &utils.AddBookDTO{}
	c.Bind(data)
	authUser, err := utils.GetSessionUser(c)
	if err != nil {
		c.AbortWithStatusJSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	data.AuthorId = authUser.ID
	dataInput := &utils.BookInput{
		AuthorId:      data.AuthorId,
		Title:         data.Title,
		ISBN:          data.ISBN,
		YearOfPublish: data.YearOfPublish,
	}

	book, err := services.AddBookService(dataInput)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, &utils.OkResponse{
		Data: book,
	})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object}  utils.OkResponse
// @Failure 400 {object}  utils.ErrResponse
// @Router /books [get]
func getBooks(c *gin.Context) {
	data := &utils.GetBooksDTO{}
	books, err := services.GetBooksService(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, utils.OkResponse{
		Data: books,
	})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object}  utils.OkResponse
// @Failure 400 {object}  utils.ErrResponse
// @Router /books [get]
func getBookById(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		println(err.Error())
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	book, err := services.GetBookById(uint(bookId))
	if err != nil {
		println(err.Error())
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, &utils.OkResponse{Data: &book})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object}  utils.OkResponse
// @Failure 400 {object}  utils.ErrResponse
// @Router /books [get]
func deleteBookByOwner(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	authUser, err := utils.GetSessionUser(c)
	if err != nil {
		c.AbortWithStatusJSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	data := &utils.DeleteBookByOwnerDTO{
		AuthorId: authUser.ID,
		BookId:   uint(bookId),
	}
	if err != nil {
		println(err.Error())
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	book, err := services.DeleteBookByOwner(data)
	if err != nil {
		println(err.Error())
		c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, &utils.OkResponse{Data: &book})
}
