package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database() {
	dns := "root:wyy666@tcp(124.223.28.250:3306)/shmily?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic("Mysql连接错误")
	} else {
		fmt.Println("Mysql连接成功")
	}
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration() //自动创建数据库表
}
