package main 

import (
	"marketplace/cmd/app/config"
	"marketplace/pkg/logger"
	"marketplace/pkg/repository"

	"marketplace/internal/repository"
	"marketplace/internal/usecases"
	"marketplace/internal/service"
	"marketplace/internal/handlers"

	"net/http"
	"github.com/go-chi/chi/v5"
)

func main()  {
	router := chi.NewRouter()

	cfg := config.NewConfig()

	log := logger.NewSlogLogger(cfg.Env)

	db := database.NewDatabase(cfg.Repository)

	repository := repository.NewRepositoryDB(db, log)
	usecases := usecases.NewUsecases(repository, log)
	service := service.NewService(usecases, log)
	handlers := handlers.NewHandlers(service)

	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)
	router.Post("/post-card", handlers.PostCard)
	router.Get("/get-cards", handlers.GetCards)



	log.Info("config is successful")
	http.ListenAndServe(cfg.HTTPServer.Address, router)

}

//TODO docker container