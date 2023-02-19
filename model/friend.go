package model

import "github.com/jinzhu/gorm"

type Friend struct {
	gorm.Model
	User         User `gorm:"ForeignKey:Uid"`
	Friend       User `gorm:"ForeignKey:FriendUid"`
	Uid          uint `gorm:"not null"`
	FriendUid    uint `gorm:"not null"`
	Relationship uint
}
