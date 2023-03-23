package serializer

import (
	"github.com/gin-gonic/gin"
	"shmily/model"
)

func BuildMemo(item model.Memo) gin.H {
	return gin.H{
		"ID":      item.ID,
		"Color":   item.Color,
		"Content": item.Content,
		"Date":    item.Date,
	}
}
