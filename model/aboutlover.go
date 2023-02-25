package model

import "github.com/jinzhu/gorm"

type AboutLover struct {
	gorm.Model
	User     User `gorm:"ForeignKey:Uid"`
	Lover    User `gorm:"ForeignKey:LoverUid"`
	Uid      uint `gorm:"not null"`
	LoverUid uint `gorm:"not null"`
	Title    string
	Content  string `gorm:"type:longtext"`
	ImageURL string
}
