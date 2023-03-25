package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique"`
	Nickname       string
	Avatar         string
	PasswordDigest string     //密码加密后的密文
	Gender         string     //性别
	Birthday       *time.Time //生日
	Signature      string     //个性签名
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

func (user *User) UpdateInfo() error {
	//err := DB.Save(&user).Error
	// 根据 `struct` 更新属性，只会更新非零值的字段
	err := DB.Model(&user).Updates(User{}).Error
	return err
}

func UpdateAvatar(id uint, avatarPath string) error {
	err := DB.Model(&User{}).Where("id = ?", id).Update("avatar", avatarPath).Error
	return err
}

func DeleteUser(uid uint) error {
	err := DB.Where("id=?", uid).Delete(&User{}).Error
	return err
}
