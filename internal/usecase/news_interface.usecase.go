package usecase

import (
	"news-topic-api/common"

	"news-topic-api/internal/delivery/data/dtos"
	response "news-topic-api/internal/delivery/data/responses"
)

type NewsUseCase interface {
	GetAllNews(pagination *common.Pagination, filter *dtos.FilterNewsRequest) (news []*response.NewsResponse, totalItems int, err error)
	CreateNews(newsDto dtos.CreateNewsRequest) (news *response.NewsResponse, err error)
	GetByUuid(uuid string) (news *response.NewsResponse, err error)
	UpdateByUuid(uuid string, newsDto dtos.UpdateNewsRequest) (*response.NewsResponse, error)
	DeleteByUuid(uuid string) error

	UpdateNewsStatus(uuid string, dto dtos.UpdateNewsStatus) (*response.NewsResponse, error)
}
