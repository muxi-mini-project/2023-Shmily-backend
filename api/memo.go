package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"shmily/pkg/utils"
	"shmily/serializer"
	"shmily/service"
)

// @Summary      Create a memo
// @Description  get a memo
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param		 memo body service.CreateMemoService true "小纸条"
// @Success      200  {string}  string"{"成功"}
// @Router       /api/v1/memo [post]

func CreateMemo(c *gin.Context) {
	var memo service.MemoService
	if err := c.ShouldBind(&memo); err == nil {
		//解析token 写到哪一个用户下面呢
		claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
		log.Printf("CreateMemo api:email=%v memo=%v\n", claim.Email, memo)
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

// @Summary      Show a memo
// @Description  get a memo
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param		 memo body service.CreateMemoService true "小纸条"
// @Success      200  {string}  string"{"成功"}"
// @Router       /api/v1/memo/:id [get]

func ShowMemo(c *gin.Context) {
	var showMemo service.MemoService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	log.Printf("ShowMemo api:email=%v memoID%v\n", claim.Email, c.Param("id"))
	res := showMemo.Show(claim.Id, c.Param("id"))
	c.JSON(200, res)
}

func RandMemo(c *gin.Context) {
	var randMemo service.MemoService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	log.Printf("RandMemo api:email=%v\n", claim.Email)
	res := randMemo.Rand(claim.Id)
	c.JSON(200, res)
}

// @Summary      Show memos
// @Description  get memos
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param		 memo body service.CreateMemoService true "小纸条"
// @Success      200  {string}  string"{"查询成功"}"
// @Router       /api/v1/memo/list [get]

func ListMemoId(c *gin.Context) {
	var showMemo service.MemoService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	log.Printf("ListMemo api:email=%v\n", claim.Email)
	res := showMemo.List(claim.Id)
	c.JSON(200, res)
}
