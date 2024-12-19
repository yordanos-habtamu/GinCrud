package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yordanos-habtamu/GinCrud/model"
	"github.com/yordanos-habtamu/GinCrud/services"
)

func GetBooks(c *gin.Context) {
	books := services.GetBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func CreateBook (c *gin.Context){
	var book model.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	_,err := services.CreateBook(book)
    if err != nil {
		c.JSON(00, gin.H{"error": err.Error()})
        return
	}
    c.JSON(200, gin.H{"message": "book created", "book": book})
}
func UpdateBook (c *gin.Context){
	var book model.Book
	id := c.Param("id")
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	err := services.UpdateBook(id,book)
    if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
        return
	}
    c.JSON(200, gin.H{"message": "book updated", "book": book})
}
func DeleteBook(c *gin.Context) {
	// Get ID from URL parameters
	id := c.Param("id")

	// Check if the book exists
	_, err := services.GetBookByID(id)
	if err != nil {
		// If the book is not found, return 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Attempt to delete the book
	err = services.DeleteBook(id)
	if err != nil {
		// Return internal server error if deletion fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}
