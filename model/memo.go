package model

import (
	"github.com/jinzhu/gorm"
)

// Memo 备忘录内容
type Memo struct {
	gorm.Model
	User    User `gorm:"ForeignKey:Uid"` //外键
	Uid     uint `gorm:"not null"`       //userid 属于某人(user)的
	Color   string
	Content string `gorm:"type:longtext"` //长字符串
}
