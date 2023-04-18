package main

import (
	"github.com/gin-gonic/gin"
	"sesi4/controller"
	"sesi4/database"
	"sesi4/repository"
)

func main() {
	database.StartDB()
	db := database.GetDB()

	bookRepository := repository.NewBookRepository(db)
	bookController := controller.NewBookController(*bookRepository)

	router := gin.Default()
	router.GET("book", bookController.GetAllBook)
	router.GET("book/:id", bookController.GetBookById)
	router.POST("book", bookController.AddBook)
	router.PUT("book/:id", bookController.UpdateBook)
	router.DELETE("book/:id", bookController.DeleteBook)
	//
	//router.GET("/shelfes", GetGroup)
	router.Run()
}
