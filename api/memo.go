package api

import (
	"2023-Shmily-bakend/pkg/utils"
	"2023-Shmily-bakend/serializer"
	"2023-Shmily-bakend/service"
	"github.com/gin-gonic/gin"
)

func CreateMemo(c *gin.Context) {
	var memo service.CreateMemoService
	if err := c.ShouldBind(&memo); err == nil {
		//解析token 写到哪一个用户下面呢
		claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
		//把小纸条数据保存到数据库中
		res := memo.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		})
	}

	//1）验证token权限
	//2）解析网络传过来的小纸条数据
	//3）保存到数据库
}
