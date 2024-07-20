package entities

import (
	"news-topic-api/common"

	"gorm.io/gorm"
)

type StatusType string

const (
	NewsStatusPublished StatusType = "published"
	NewsStatusDraft     StatusType = "draft"
	NewsStatusDeleted   StatusType = "deleted"
)

type News struct {
	common.Base
	Title    string   `gorm:"type:varchar(255)" json:"title"`
	Content  string   `gorm:"type:text" json:"content"`
	StatusId *uint64  `gorm:"type:int" json:"status_id"`
	Status   Statuses `gorm:"foreignKey:StatusId;references:ID" json:"status"`
	Topics   []Topic  `gorm:"many2many:news_topics" json:"topics"`
	gorm.Model
}
