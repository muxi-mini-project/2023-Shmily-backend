package service

import (
	"github.com/gin-gonic/gin"
	"shmily/model"
)

func GetId(c *gin.Context) uint {
	id, _ := c.Get("user_id")
	ID, _ := id.(uint)
	return ID
}

func WhetherRe(re model.Relationship) bool {
	var d model.Relationship
	err1 := model.DB.Model(&model.Relationship{}).Where("ID2 = ? AND ID1 = ? AND num = ?", re.ID2, re.ID1, re.Num).Take(&d).Error
	re1 := re
	re1.ID1 = re.ID2
	re1.ID2 = re.ID1
	var err2 error = nil
	if re1.Num == 2 {
		err2 = model.DB.Model(&model.Relationship{}).Where("ID2 = ? AND ID1 = ? AND num = 2", re.ID2, re.ID1).Take(&d).Error
	}
	if err1 != nil || err2 != nil {
		return false
	}
	return true
}
