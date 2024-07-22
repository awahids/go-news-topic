package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"news-topic-api/internal/delivery/handlers"
	"news-topic-api/internal/repositories"
	"news-topic-api/internal/usecase"
)

func NewsRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	validate := validator.New()

	newsRepo := repositories.NewNewsRepositoryGorm(db)
	topicRepo := repositories.NewTopicRepositoryGorm(db)
	newsUc := usecase.NewNewsUseCase(newsRepo, topicRepo, validate)
	handler := handlers.NewNewsHandler(newsUc)

	r.Get("/", handler.GetNews)
	r.Post("/", handler.CreateNews)
	r.Put("/status/{uuid}", handler.UpdateNewsStatus)

	r.Route("/{uuid}", func(r chi.Router) {
		r.Get("/", handler.GetNewsByUuid)
		r.Put("/", handler.UpdateNews)
		r.Delete("/", handler.DeleteNews)
	})

	return r
}
