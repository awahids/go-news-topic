package repositories

import (
	"errors"
	"news-topic-api/common"

	"gorm.io/gorm"

	"news-topic-api/internal/delivery/data/dtos"
	"news-topic-api/internal/entities"
)

type newsRepositoryGorm struct {
	db *gorm.DB
}

func NewNewsRepositoryGorm(db *gorm.DB) NewsRepository {
	return &newsRepositoryGorm{db}
}

func (r *newsRepositoryGorm) BeginTransaction() (*gorm.DB, error) {
	return r.db.Begin(), nil
}

func (r *newsRepositoryGorm) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *newsRepositoryGorm) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *newsRepositoryGorm) GetNews(pagination *common.Pagination, filter *dtos.FilterNewsRequest) (news []*entities.News, items int64, err error) {
	query := r.db.Model(&entities.News{})

	if filter.Title != nil {
		query = query.Where("title ILIKE ?", "%"+*filter.Title+"%")
	}
	if filter.Topic != nil {
		query = query.Joins("JOIN news_topics nt ON nt.news_id = news.id").
			Joins("JOIN topics t ON t.id = nt.topic_id").
			Where("t.value = ?", filter.Topic)
	}
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	err = query.Count(&items).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at desc").
		Preload("Topics").
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Find(&news).Error

	if err != nil {
		return nil, 0, err
	}

	return news, items, nil
}

func (r *newsRepositoryGorm) GetByUuid(uuid string) (news *entities.News, err error) {
	result := r.db.Where("uuid = ?", uuid).First(&news)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("news not found")
	}

	return news, nil
}

func (r *newsRepositoryGorm) CreateNews(news *entities.News) (*entities.News, error) {
	result := r.db.Create(news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}

func (r *newsRepositoryGorm) UpdateByUuid(uuid string, news *entities.News) (*entities.News, error) {
	existingNews, err := r.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(existingNews).Updates(news).Error; err != nil {
		return nil, err
	}

	return existingNews, nil
}

func (r *newsRepositoryGorm) DeleteByUuid(uuid string) error {
	if err := r.db.Where("uuid = ?", uuid).Delete(&entities.News{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *newsRepositoryGorm) UpdateNewsStatus(uuid string, dto dtos.UpdateNewsStatus) (*entities.News, error) {
	existingNews, err := r.GetByUuid(uuid)
	if err != nil {
		return nil, err
	}

	existingNews.Status = entities.StatusType(dto.Status)

	if err := r.db.Save(existingNews).Error; err != nil {
		return nil, err
	}

	return existingNews, nil
}

func (r *newsRepositoryGorm) LoadTopics(news *entities.News) error {
	return r.db.Model(news).Association("Topics").Find(&news.Topics)
}
