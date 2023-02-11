package middleware

import (
	"2023-Shmily-bakend/pkg/utils"
	"2023-Shmily-bakend/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404 //没有token,无权限
		} else {
			//解析token
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //token错误，无权限
			} else if time.Now().Unix() > claim.ExpireseAt {
				code = 401 //token过期
			}
			//1)token对不对
			//2)token是否过期
		}
		//将错误码返回给前端
		if code != 200 {
			c.JSON(400, serializer.Response{
				Status: code,
				Msg:    "token错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
