package api

import (
	"github.com/gin-gonic/gin"
	"shmily/pkg/utils"
	"shmily/service"
	"strconv"
)

func SaveLocation(c *gin.Context) {
	var myLocation service.LocationService
	if err := c.ShouldBind(&myLocation); err == nil {
		claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
		myLocation.Uid = claim.Id
		res := myLocation.Save()
		c.JSON(200, res)
	}
}

func GetFriendLocation(c *gin.Context) {
	var myLocation service.LocationService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	myLocation.Uid = claim.Id
	id, _ := strconv.Atoi(c.Param("id"))
	res := myLocation.GetFriendLocation(uint(id))
	c.JSON(200, res)
}

func GetFriendsLocations(c *gin.Context) {
	var myLocation service.LocationService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	myLocation.Uid = claim.Id

	res := myLocation.GetFriendsLocations()
	c.JSON(200, res)
}
