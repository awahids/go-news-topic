package repositories

import (
	"news-topic-api/common"

	"news-topic-api/internal/entities"
)

type TopicRepository interface {
	GetTopics(pagination *common.Pagination) (topics []*entities.Topic, items int64, err error)
	GetByUuid(uuid string) (topic *entities.Topic, err error)
	CreateTopic(topic *entities.Topic) (*entities.Topic, error)
	UpdateByUuid(uuid string, topic *entities.Topic) (*entities.Topic, error)
	DeleteByUuid(uuid string) error
}
