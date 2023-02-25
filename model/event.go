package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Event struct {
	gorm.Model
	User    User   `gorm:"ForeignKey:Uid"`
	Uid     uint   `gorm:"not null"`
	Content string `gorm:"type:longtext"`
	Date    *time.Time
}
