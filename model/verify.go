package model

import (
	"time"
)

type Verify struct {
	//gorm.Model
	//Email     string
	Code     string //系统发送的验证码
	ExpireAt time.Time
	//Character string //用户输入的验证码
	//Password  string
}
