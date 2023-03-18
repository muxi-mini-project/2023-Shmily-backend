package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Letter struct {
	gorm.Model
	UserFrom    User       `gorm:"ForeignKey:UidFrom"`
	UserTo      User       `gorm:"ForeignKey:UidTo"`
	UidFrom     uint       `gorm:"not null"`
	UidTo       uint       `gorm:"not null"`
	StampImgURL string     //邮票图片地址
	Title       string     //信封标题
	Salutation  string     //收件人称呼
	Content     string     `gorm:"type:longtext"` //信件内容
	Nickname    string     //发件人昵称
	Date        *time.Time //写信时间
}

func (l *Letter) Create(uid uint) error {
	err := DB.Create(&Letter{}).Error

	return err
}
