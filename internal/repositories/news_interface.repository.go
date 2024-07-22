package repositories

import (
	"news-topic-api/common"

	"gorm.io/gorm"

	"news-topic-api/internal/delivery/data/dtos"
	"news-topic-api/internal/entities"
)

type NewsRepository interface {
	BeginTransaction() (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error

	GetNews(pagination *common.Pagination, filter *dtos.FilterNewsRequest) (news []*entities.News, items int64, err error)
	GetByUuid(uuid string) (*entities.News, error)
	CreateNews(news *entities.News) (*entities.News, error)
	UpdateByUuid(uuid string, news *entities.News) (*entities.News, error)
	DeleteByUuid(uuid string) error
	UpdateNewsStatus(uuid string, dto dtos.UpdateNewsStatus) (*entities.News, error)

	LoadTopics(news *entities.News) error
}
