package routes

import (
	"net/http"
	_ "news-topic-api/docs"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// @title News Topic API
// @version 2.0
// @description This is a sample server for managing news topics.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:9000
// @BasePath /api/v1
// @schemes http
func InitRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(v1 chi.Router) {
		// swagger
		v1.Get("/swagger/*", httpSwagger.WrapHandler)

		// base route
		v1.Get("/", func(w http.ResponseWriter, r *http.Request) {
			timeNow := time.Now().Format("2006-01-02 15:04:05")

			w.Write([]byte(timeNow))
		})

		// topic
		v1.Mount("/", TopicRouter(db))
	})

	return r
}
