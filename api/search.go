package api

import (
	"github.com/gin-gonic/gin"
	"shmily/serializer"
	"shmily/service"
)

func Search(c *gin.Context) {
	var user service.UserSearch
	var res serializer.Response
	if err := c.ShouldBind(&user); err == nil {

		if user.Email != "" {
			res = user.SearchByEmail()
		} else if user.Nickname != "" {
			res = user.SearchByNickname()
		} else if user.ID != "" {
			res = user.SearchByID()
		}
		c.JSON(200, res)
	}
}
