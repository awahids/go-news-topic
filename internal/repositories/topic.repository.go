package repositories

import (
	"errors"
	"news-topic-api/common"

	"gorm.io/gorm"

	"news-topic-api/internal/entities"
)

type topicRepositoryGorm struct {
	db *gorm.DB
}

func NewTopicRepositoryGorm(db *gorm.DB) TopicRepository {
	return &topicRepositoryGorm{db}
}

func (r *topicRepositoryGorm) GetTopics(pagination *common.Pagination) (topics []*entities.Topic, items int64, err error) {
	err = r.db.Model(&topics).
		Count(&items).
		Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Order("created_at desc").
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Find(&topics).
		Error

	if err != nil {
		return nil, 0, err
	}

	return topics, items, nil
}

func (r *topicRepositoryGorm) GetByUuid(uuid string) (topic *entities.Topic, err error) {
	result := r.db.Find(&topic, "uuid = ?", uuid)

	if result.Error != nil {
		return topic, result.Error
	} else if result.RowsAffected == 0 {
		return topic, errors.New("topic not found")
	}

	return topic, nil
}

func (r *topicRepositoryGorm) CreateTopic(topic *entities.Topic) (*entities.Topic, error) {
	result := r.db.Create(topic)
	if result.Error != nil {
		return topic, result.Error
	}
	return topic, nil
}

func (r *topicRepositoryGorm) UpdateByUuid(uuid string, topic *entities.Topic) (*entities.Topic, error) {
	findTopic, _ := r.GetByUuid(uuid)

	result := r.db.Model(&entities.Topic{}).
		Where("id = ?", findTopic.Id).
		Updates(topic)

	if result.Error != nil {
		return nil, result.Error
	}

	updatedTopic := &entities.Topic{}
	r.db.Where("id = ?", findTopic.Id).First(updatedTopic)
	return updatedTopic, nil
}

func (r *topicRepositoryGorm) DeleteByUuid(uuid string) error {
	result := r.db.Delete(&entities.Topic{}, "uuid = ?", uuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
