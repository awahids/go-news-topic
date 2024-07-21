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

	r.Post("/topic", handler.CreateTopic)
	r.Get("/topics", handler.GetTopics)
	r.Get("/topic/{uuid}", handler.GetTopic)
	r.Put("/topic/{uuid}", handler.UpdateTopic)
	r.Delete("/topic/{uuid}", handler.DeleteTopic)

	return r
}
