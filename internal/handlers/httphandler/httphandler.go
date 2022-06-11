package httphandler

import (
	"CleanTemplate/internal/models"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	router      *chi.Mux
	logger      *logrus.Logger
	bookService bookService
}

type bookService interface {
	GetById(id string) (*models.Book, error)
	Add(book *models.Book)
}

func NewHandler(log *logrus.Logger, bookService bookService) *handler {
	return &handler{
		router:      chi.NewRouter(),
		logger:      log,
		bookService: bookService,
	}
}

func (h *handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")

	book, err := h.bookService.GetById(bookId)

	if err != nil {
		h.logger.Error("error when get book:", err)
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, "Book Name: %s, Book Author: %s", book.Name, book.Author)

}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func (h *handler) InitRoutes() *chi.Mux {
	h.router.Get("/health", h.HealthCheck)
	h.router.Get("/book/{bookId}", h.GetBookById)

	return h.router
}
