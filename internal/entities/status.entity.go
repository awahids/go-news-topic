package entities

import (
	"news-topic-api/common"

	"gorm.io/gorm"
)

type Statuses struct {
	common.Base
	Title string `gorm:"unique;type:varchar(255)" json:"title"`
	Value string `gorm:"unique;type:varchar(255)" json:"value"`
	gorm.Model
}
