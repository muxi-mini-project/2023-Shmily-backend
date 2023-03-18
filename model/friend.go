package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Relationship struct {
	gorm.Model
	ID1  uint
	ID2  uint   `form:"id_object"`
	Type string `form:"type"`
	Num  uint   `form:"num"` //=0为双向但是待验证，=1为id1对id2单向
}

type Friend struct {
	gorm.Model
	User         User `gorm:"ForeignKey:Uid"`
	Friend       User `gorm:"ForeignKey:FriendUid"`
	Uid          uint `gorm:"not null"`
	FriendUid    uint `gorm:"not null"`
	Relationship uint
}

type Res struct {
	One  []string
	Both []string
}

// FriendsAdd 添加好友
func FriendsAdd(relationship Relationship) error {
	err := DB.Create(&relationship).Error
	return err
}

// AddedFailure 删除好友
func AddedFailure(ID interface{}, ID2 interface{}) error {
	err := DB.Where("ID2 = ? AND ID1 = ? AND num = 0", ID, ID2).Delete(&Relationship{}).Error
	return err
}

// AddedSuccess 更新
func AddedSuccess(ID interface{}, ID2 interface{}) error {
	err := DB.Where("ID2 = ? AND ID1 = ? AND num = 1 || ID2 = ? AND ID1 = ? AND num = 1 ", ID, ID2, ID2, ID).Delete(&Relationship{}).Error
	err = DB.Model(&Relationship{}).Where("ID2 = ? AND ID1 = ? AND num = 0 ", ID, ID2).Update("num", 2).Error
	return err
}

// QueryFriends 好友查询
func QueryFriends(ID interface{}, Type string) (Res, error) {
	type IDS struct {
		ID1 string
		ID2 string
		Num string
	}
	var ids []IDS
	err := DB.Model(&Relationship{}).Select("relationships.id1,relationships.id2,relationships.num").Joins("join users on (relationships.id1 = users.id AND relationships.num != 0) || (relationships.id2 = users.id AND relationships.num = 2) ").Where("users.id = ? AND type =? ", ID, Type).Find(&ids).Error
	var one []string
	var both []string
	I := fmt.Sprintf("%v", ID)
	fmt.Println(ids)
	for _, id := range ids {
		if id.ID1 != I {
			if id.Num == "1" {
				one = append(one, id.ID1)
			} else {
				both = append(both, id.ID1)
			}
		} else {
			if id.Num == "1" {
				one = append(one, id.ID2)
			} else {
				both = append(both, id.ID2)
			}
		}
	}

	res := Res{
		One:  one,
		Both: both,
	}
	return res, err
}

func FriendsAddedRequest(ID interface{}) ([]string, error) {
	var users []string
	err := DB.Table("relationships").Select("users.id").Joins("join users on relationships.id1 = users.id").Where("id2 = ? AND num = 0 ", ID).Find(&users).Error
	return users, err
}

func IdToUser(id string, b bool) (User, error) {
	var user User
	err := DB.Model(&User{}).Where("id = ?", id).Find(&user).Error
	if b == false {
		user.PasswordDigest = ""
	}
	return user, err
}
