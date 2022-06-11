package repositories

import (
	"CleanTemplate/internal/models"
	"fmt"
)

// BooksRepository - using map for example
type booksRepository struct {
	books map[string]*models.Book
}

func NewBooksRepository() *booksRepository {
	return &booksRepository{
		books: map[string]*models.Book{
			"1234": &models.Book{
				Id:     "1234",
				Name:   "Clean Template",
				Author: "Siraj Guseynov",
			},
		},
	}
}

func (bookRepo *booksRepository) Add(book *models.Book) {
	bookRepo.books[book.Id] = book
}

func (bookRepo *booksRepository) GetById(id string) (*models.Book, error) {
	book, ok := bookRepo.books[id]

	if !ok {
		return nil, fmt.Errorf("book not found")
	}

	return book, nil
}
