package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// MailAuth 定义验证码的存储结构
type MailAuth struct {
	Code     string
	ExpireAt time.Time
}

type StoreAuth map[string]*MailAuth

var store = make(StoreAuth, 0)

func getCode() *MailAuth {
	return &MailAuth{
		// 生成六位的验证码
		Code: fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000)),
		// 设置过期时间
		ExpireAt: time.Now().Add(time.Second * 2),
	}
}

func destroyCode() {
	// 每十分钟清理无效的验证
	ticker := time.NewTicker(time.Minute * 10)
	for {
		<-ticker.C
		for key := range store {
			if store[key].ExpireAt.Before(time.Now()) {
				delete(store, key)
			}
		}
	}
}

func main() {
	go destroyCode()
	engine := gin.Default()
	engine.GET("/email/code", SendEMail)
	engine.Run()
}

func SendEMail(c *gin.Context) {
	email := c.Qurey("email")
	if email == "" || 是否是邮箱 {
		return
	}
	auth := getCode()
	// 发送邮件
	Send()
	store[email] = auth
}
