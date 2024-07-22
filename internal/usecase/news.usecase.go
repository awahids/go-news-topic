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

type newsUseCase struct {
	newsRepo  repositories.NewsRepository
	topicRepo repositories.TopicRepository
	validate  *validator.Validate
}

func NewNewsUseCase(newsRepo repositories.NewsRepository, topicRepo repositories.TopicRepository, validate *validator.Validate) NewsUseCase {
	return &newsUseCase{
		newsRepo:  newsRepo,
		topicRepo: topicRepo,
		validate:  validate,
	}
}

func (uc *newsUseCase) GetAllNews(pagination *common.Pagination, filter *dtos.FilterNewsRequest) (news []*response.NewsResponse, totalItems int, err error) {
	newsEntities, totalItems64, err := uc.newsRepo.GetNews(pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	newsResponses := []*response.NewsResponse{}
	for _, newsEntity := range newsEntities {
		if err := uc.newsRepo.LoadTopics(newsEntity); err != nil {
			return nil, 0, err
		}

		topics := newsEntity.Topics
		topicResponses := make([]response.TopicResponse, len(topics))
		for i, topic := range topics {
			topicResponses[i] = response.TopicResponse{
				Id:    topic.Id,
				UUID:  topic.UUID,
				Title: topic.Title,
				Value: topic.Value,
			}
		}

		newsResponses = append(newsResponses, &response.NewsResponse{
			Id:      newsEntity.Id,
			UUID:    newsEntity.UUID,
			Title:   newsEntity.Title,
			Content: newsEntity.Content,
			Status:  string(newsEntity.Status),
			Topics:  topicResponses,
		})
	}

	return newsResponses, int(totalItems64), nil
}

func (uc *newsUseCase) GetByUuid(uuid string) (*response.NewsResponse, error) {
	newsEntity, err := uc.newsRepo.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	if err := uc.newsRepo.LoadTopics(newsEntity); err != nil {
		return nil, err
	}

	topicResponses := make([]response.TopicResponse, len(newsEntity.Topics))
	for i, topic := range newsEntity.Topics {
		topicResponses[i] = response.TopicResponse{
			Id:    topic.Id,
			UUID:  topic.UUID,
			Title: topic.Title,
			Value: topic.Value,
		}
	}

	newsResponse := &response.NewsResponse{
		Id:      newsEntity.Id,
		UUID:    newsEntity.UUID,
		Title:   newsEntity.Title,
		Content: newsEntity.Content,
		Status:  string(newsEntity.Status),
		Topics:  topicResponses,
	}

	return newsResponse, nil
}

func (uc *newsUseCase) CreateNews(newsDto dtos.CreateNewsRequest) (*response.NewsResponse, error) {
	if err := uc.validate.Struct(&newsDto); err != nil {
		return nil, err
	}

	tx, err := uc.newsRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			uc.newsRepo.RollbackTransaction(tx)
		} else if err != nil {
			uc.newsRepo.RollbackTransaction(tx)
		} else {
			err = uc.newsRepo.CommitTransaction(tx)
		}
	}()

	var status entities.StatusType
	switch newsDto.Status {
	case "published":
		status = entities.NewsStatusPublished
	case "draft":
		status = entities.NewsStatusDraft
	default:
		return nil, errors.New("invalid status")
	}

	var topicEntities []entities.Topic
	var topicResponses []response.TopicResponse
	for _, topicDto := range newsDto.Topics {
		topicEntity, err := uc.topicRepo.GetByUuid(topicDto.Uuid)
		if err != nil {
			return nil, err
		}
		if topicEntity == nil {
			return nil, errors.New("topic entity not found")
		}

		topicEntities = append(topicEntities, *topicEntity)
		topicResponses = append(topicResponses, response.TopicResponse{
			Id:    topicEntity.Id,
			UUID:  topicEntity.UUID,
			Title: topicEntity.Title,
			Value: topicEntity.Value,
		})
	}

	newsEntity := &entities.News{
		Title:   newsDto.Title,
		Content: newsDto.Content,
		Status:  status,
		Topics:  topicEntities,
	}

	newsEntity, err = uc.newsRepo.CreateNews(newsEntity)
	if err != nil {
		return nil, err
	}

	newsResponse := &response.NewsResponse{
		Id:      newsEntity.ID,
		UUID:    newsEntity.UUID,
		Title:   newsEntity.Title,
		Content: newsEntity.Content,
		Status:  string(newsEntity.Status),
		Topics:  topicResponses,
	}

	return newsResponse, nil
}

func (uc *newsUseCase) UpdateByUuid(uuid string, newsDto dtos.UpdateNewsRequest) (*response.NewsResponse, error) {
	if err := uc.validate.Struct(&newsDto); err != nil {
		return nil, err
	}

	existingNews, err := uc.newsRepo.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	if existingNews.Status != entities.NewsStatusDraft {
		return nil, errors.New("news is not in draft status")
	}

	if newsDto.Title != "" {
		existingNews.Title = newsDto.Title
	}
	if newsDto.Content != "" {
		existingNews.Content = newsDto.Content
	}

	if newsDto.Status != "" {
		var status entities.StatusType
		switch newsDto.Status {
		case "published":
			status = entities.NewsStatusPublished
		default:
			return nil, errors.New("invalid status")
		}
		existingNews.Status = status
	}

	if len(newsDto.Topics) > 0 {
		topicEntities := []entities.Topic{}
		for _, topicDto := range newsDto.Topics {
			topicEntity, err := uc.topicRepo.GetByUuid(topicDto.Uuid)
			if err != nil {
				return nil, err
			}
			if topicEntity == nil {
				return nil, errors.New("topic entity not found")
			}
			topicEntities = append(topicEntities, *topicEntity)
		}
		existingNews.Topics = topicEntities
	}

	updatedNews, err := uc.newsRepo.UpdateByUuid(uuid, existingNews)
	if err != nil {
		return nil, err
	}

	topicResponses := make([]response.TopicResponse, len(updatedNews.Topics))
	for i, topic := range updatedNews.Topics {
		topicResponses[i] = response.TopicResponse{
			Id:    topic.Id,
			UUID:  topic.UUID,
			Title: topic.Title,
			Value: topic.Value,
		}
	}

	newsResponse := &response.NewsResponse{
		Id:      updatedNews.Id,
		UUID:    updatedNews.UUID,
		Title:   updatedNews.Title,
		Content: updatedNews.Content,
		Status:  string(updatedNews.Status),
		Topics:  topicResponses,
	}

	return newsResponse, nil
}

func (uc *newsUseCase) DeleteByUuid(uuid string) error {
	newsExisting, err := uc.newsRepo.GetByUuid(uuid)
	if err != nil {
		return err
	}

	if newsExisting.Status == entities.NewsStatusDeleted && newsExisting.DeletedAt.Valid {
		return errors.New("news already deleted")
	}

	updateStatusDto := dtos.UpdateNewsStatus{
		Status: string(entities.NewsStatusDeleted),
	}
	if _, err := uc.UpdateNewsStatus(uuid, updateStatusDto); err != nil {
		return err
	}

	return uc.newsRepo.DeleteByUuid(uuid)
}

func (uc *newsUseCase) UpdateNewsStatus(uuid string, dto dtos.UpdateNewsStatus) (*response.NewsResponse, error) {
	status := entities.StatusType(dto.Status)
	if status != entities.NewsStatusPublished && status != entities.NewsStatusDeleted {
		return nil, errors.New("invalid status")
	}

	updatedNews, err := uc.newsRepo.UpdateNewsStatus(uuid, dto)
	if err != nil {
		return nil, err
	}

	newsResponse := &response.NewsResponse{
		Id:      updatedNews.Id,
		UUID:    updatedNews.UUID,
		Title:   updatedNews.Title,
		Content: updatedNews.Content,
		Status:  string(updatedNews.Status),
	}

	return newsResponse, nil
}
