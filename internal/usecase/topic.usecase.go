package usecase

import (
	"errors"
	"news-topic-api/common"

	"github.com/go-playground/validator/v10"

	"news-topic-api/internal/delivery/data/dtos"
	response "news-topic-api/internal/delivery/data/responses"
	"news-topic-api/internal/entities"
	"news-topic-api/internal/repositories"
)

type topicUseCase struct {
	topicRepo repositories.TopicRepository
	validate  *validator.Validate
}

func NewTopicUseCase(topicRepo repositories.TopicRepository, validate *validator.Validate) TopicUseCase {
	return &topicUseCase{
		topicRepo,
		validate,
	}
}

func (uc *topicUseCase) GetAllTopics(pagination *common.Pagination) (topics []*response.TopicResponse, totalItems int, err error) {
	topicModel, totalItems64, err := uc.topicRepo.GetTopics(pagination)
	if err != nil {
		return nil, 0, err
	}

	totalItems = int(totalItems64)

	for _, topic := range topicModel {
		topics = append(topics, &response.TopicResponse{
			Id:    uint(topic.Id),
			UUID:  topic.UUID,
			Title: topic.Title,
			Value: topic.Value,
		})
	}

	return topics, totalItems, nil
}

func (uc *topicUseCase) GetByUuid(uuid string) (topic *response.TopicResponse, err error) {
	topicModel, err := uc.topicRepo.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	topic = &response.TopicResponse{
		Id:    uint(topicModel.Id),
		UUID:  topicModel.UUID,
		Title: topicModel.Title,
		Value: topicModel.Value,
	}

	return topic, nil
}

func (uc *topicUseCase) CreateTopic(topicDto dtos.CreateTopicRequest) (*response.TopicResponse, error) {
	if err := uc.validate.Struct(&topicDto); err != nil {
		return nil, err
	}

	if topicDto.Title == "" {
		return nil, errors.New("topic name cannot be empty")
	}

	if topicDto.Value == "" {
		return nil, errors.New("topic value cannot be empty")
	}

	createTopic, err := uc.topicRepo.CreateTopic(
		&entities.Topic{
			Title: topicDto.Title,
			Value: topicDto.Value,
		},
	)

	if err != nil {
		return nil, err
	}

	topicRes := &response.TopicResponse{
		Id:    createTopic.Id,
		UUID:  createTopic.UUID,
		Title: createTopic.Title,
		Value: createTopic.Value,
	}

	return topicRes, nil
}

func (uc *topicUseCase) UpdateByUuid(uuid string, topicDto dtos.UpdateTopicRequest) (*response.TopicResponse, error) {
	if err := uc.validate.Struct(&topicDto); err != nil {
		return nil, err
	}

	topicRes, topicErr := uc.topicRepo.UpdateByUuid(
		uuid,
		&entities.Topic{
			Title: topicDto.Title,
		},
	)

	if topicErr != nil {
		return nil, topicErr
	}

	topicResponse := &response.TopicResponse{
		Id:    topicRes.Id,
		UUID:  topicRes.UUID,
		Title: topicRes.Title,
		Value: topicRes.Value,
	}

	return topicResponse, nil
}

func (uc *topicUseCase) DeleteByUuid(uuid string) error {
	return uc.topicRepo.DeleteByUuid(uuid)
}
