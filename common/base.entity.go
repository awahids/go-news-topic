package common

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id   uint   `gorm:"type:int;not null auto_increment primary_key" `
	UUID string `gorm:"type:string;default:gen_random_uuid()"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.NewString()
	return
}
