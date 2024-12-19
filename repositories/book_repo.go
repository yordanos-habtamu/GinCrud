package repositories

import (
	"github.com/yordanos-habtamu/GinCrud/model"
	"errors"
)

var books = []model.Book{
	{ID: "1", Title: "1984", Author: "George Orwell", Year: "1949"},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: "1960"},
}

func GetAllBooks() []model.Book {
	return books
}

func GetBookByID(id string) (model.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, errors.New("Book not found")
}

func AddBook(book model.Book)  (model.Book,error){
	books = append(books, book)
	return book,nil;
}

func UpdateBook(id string, updatedBook model.Book) error {
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			return nil
		}
	}
	return errors.New("Book not found")
}

func DeleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("Book not found")
}
