package routes

import (
	"net/http"

	"github.com/datrine/routes/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/login", basicLogin)

	bookRoutes := r.Group("/books", middleware.Auth)

	bookRoutes.GET("/", getBooks)
	bookRoutes.GET("/:id", getBookById)
	bookRoutes.POST("/", addBook)
	bookRoutes.DELETE("/:id", deleteBookByOwner)

	authRoutes := r.Group("/auth")
	authRoutes.POST("/login/basic", basicLogin)
	authRoutes.POST("/register", registerRoute)

	return r
}
