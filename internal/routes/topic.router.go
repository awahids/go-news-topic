package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"news-topic-api/internal/delivery/handlers"
	"news-topic-api/internal/repositories"
	"news-topic-api/internal/usecase"
)

func TopicRouter(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	validate := validator.New()

	topicRepo := repositories.NewTopicRepositoryGorm(db)
	topicUc := usecase.NewTopicUseCase(topicRepo, validate)
	handler := handlers.NewTopicHandler(topicUc)

	r.Post("/", handler.CreateTopic)
	r.Get("/", handler.GetTopics)

	r.Route("/{uuid}", func(r chi.Router) {
		r.Get("/", handler.GetTopic)
		r.Put("/", handler.UpdateTopic)
		r.Delete("/", handler.DeleteTopic)
	})

	return r
}
