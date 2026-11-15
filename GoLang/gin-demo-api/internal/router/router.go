package router

import (
	"my-app/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(bookHandler *handler.BookHandler) *gin.Engine { 

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		books := api.Group("/books")
		books.GET("" , bookHandler.ListBooks)
		books.POST("" , bookHandler.CreateBook)
		books.GET("/:id", bookHandler.GetBookById)
		books.PATCH("/:id" , bookHandler.Update)
		books.DELETE("/:id" , bookHandler.Delete)
	}

	return router
}