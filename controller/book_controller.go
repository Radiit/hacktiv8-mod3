package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sesi4/model"
	"sesi4/repository"
	"strconv"
)

type BookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(bookRepository repository.BookRepository) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}

func (bc *BookController) GetAllBook(c *gin.Context) {
	mapBook, err := bc.bookRepository.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, mapBook)
}

func (bc *BookController) AddBook(c *gin.Context) {
	var newBook model.ItemBook

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedBook, err := bc.bookRepository.Save(newBook)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, savedBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	updatedBook := model.ItemBook{}

	err := c.ShouldBindJSON(&updatedBook)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uintID := uint(id)
	_, err = bc.bookRepository.GetID(uintID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	success, err := bc.bookRepository.Update(updatedBook, uintID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, success)
}

func (bc *BookController) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	uintId := uint(id)
	bookTarget, err := bc.bookRepository.GetID(uintId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, bookTarget)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	deletedBook := model.ItemBook{}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uintId := uint(id)
	_, err = bc.bookRepository.GetID(uintId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	bc.bookRepository.Delete(deletedBook, uintId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete the book",
	})
	return
}
