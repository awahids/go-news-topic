package usecase

import (
	"news-topic-api/common"

	"news-topic-api/internal/delivery/data/dtos"
	response "news-topic-api/internal/delivery/data/responses"
)

type TopicUseCase interface {
	GetAllTopics(pagination *common.Pagination) (topics []*response.TopicResponse, totalItems int, err error)
	GetByUuid(uuid string) (topic *response.TopicResponse, err error)
	CreateTopic(topicDto dtos.CreateTopicRequest) (topicRes *response.TopicResponse, err error)
	UpdateByUuid(uuid string, topicDto dtos.UpdateTopicRequest) (*response.TopicResponse, error)
	DeleteByUuid(uuid string) error
}
