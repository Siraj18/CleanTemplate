package services

import "CleanTemplate/internal/models"

type booksRepository interface {
	GetById(id string) (*models.Book, error)
	Add(book *models.Book)
}

type booksService struct {
	booksRepo booksRepository
}

func NewBooksService(bookRepo booksRepository) *booksService {
	return &booksService{
		booksRepo: bookRepo,
	}
}

func (bookService *booksService) Add(book *models.Book) {
	bookService.booksRepo.Add(book)
}

func (bookService *booksService) GetById(id string) (*models.Book, error) {
	book, err := bookService.booksRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return book, nil
}
