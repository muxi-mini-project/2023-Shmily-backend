package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"shmily/serializer"
	"shmily/service"
)

func Search(c *gin.Context) {
	var user service.UserSearch
	var res serializer.Response
	if err := c.ShouldBind(&user); err == nil {
		log.Printf("Search api:%v\n", user)

		if user.Email != "" {
			res = user.SearchByEmail(user.Email)
		} else if user.Nickname != "" {
			res = user.SearchByNickname(user.Nickname)
		} else if user.ID != "" {
			res = user.SearchByID(user.ID)
		}
		c.JSON(200, res)
	}
}
