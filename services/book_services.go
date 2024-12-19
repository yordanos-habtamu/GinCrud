package services

import (
	"github.com/yordanos-habtamu/GinCrud/model"
	"github.com/yordanos-habtamu/GinCrud/repositories"
)

func GetBooks() []model.Book {
	return repositories.GetAllBooks()
}

func GetBookByID(id string) (model.Book, error) {
	return repositories.GetBookByID(id)
}

func CreateBook(book model.Book)(model.Book,error) {
	return repositories.AddBook(book)

}

func UpdateBook(id string, updatedBook model.Book) error {
	return repositories.UpdateBook(id, updatedBook)
}

func DeleteBook(id string) error {
	return repositories.DeleteBook(id)
}
