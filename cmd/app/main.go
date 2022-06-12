package main

import (
	"CleanTemplate/internal/config"
	"CleanTemplate/internal/handlers/httphandler"
	"CleanTemplate/internal/repositories"
	"CleanTemplate/internal/services"
	"CleanTemplate/pkg/httpserver"
	"CleanTemplate/pkg/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()

	logger := logging.NewLogger()

	bookRepo := repositories.NewBooksRepository()

	bookService := services.NewBooksService(bookRepo)

	bookHandler := httphandler.NewHandler(logger, bookService)

	server := httpserver.NewServer(bookHandler.InitRoutes(), logger, httpserver.SetAddress(cfg.Address))

	if err := server.Run(); err != nil {
		logrus.Fatal(err)
	}
}
