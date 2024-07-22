package entities

import (
	"news-topic-api/common"

	"gorm.io/gorm"
)

type Topic struct {
	common.Base
	Title string `gorm:"unique;type:varchar(255)" json:"title"`
	Value string `gorm:"unique;type:varchar(255)" json:"value"`
	News  []News `gorm:"many2many:news_topics;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"news"`
	gorm.Model
}
