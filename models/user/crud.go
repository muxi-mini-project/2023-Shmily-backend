package user

import (
	"2023-Shmily-backend/pkg/logger"
	"2023-Shmily-backend/pkg/model"
	"go/types"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Get 通过 ID 获取用户
func Get(idstr string) (User, error) {
	var user User
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Update 更新用户
func (user *User) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&user)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}
