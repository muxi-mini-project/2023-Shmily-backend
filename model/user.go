package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Email          string     `gorm:"unique" json:"email" form:"email"`
	Nickname       string     `json:"nickname" form:"nickname"`
	Avatar         string     `json:"avatar" form:"avatar"`
	PasswordDigest string     `json:"passwordDigest" form:"passwordDigest"` //密码加密后的密文
	Gender         string     `json:"gender" form:"gender"`                 //性别
	Birthday       *time.Time `json:"birthday" form:"birthday"`             //生日
	Signature      string     `json:"signature" for,m:"signature"`          //个性签名
}

// SetPassword 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.PasswordDigest = string(bytes)
	return err
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) UpdateInfo(u User, id uint) error {
	// 查询单个用户
	if err := DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}

	// 更新用户信息
	if err := DB.Model(&user).Updates(u).Error; err != nil {
		return err
	}

	return nil
}

func UpdateAvatar(id uint, avatarPath string) error {
	err := DB.Model(&User{}).Where("id = ?", id).Update("avatar", avatarPath).Error
	return err
}

func DeleteUser(uid uint) error {
	err := DB.Where("id=?", uid).Delete(&User{}).Error
	return err
}
