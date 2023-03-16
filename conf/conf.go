package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"shmily/model"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

var (
	EmailHost     string
	EmailAddr     string
	EmailPassword string
)

func LoadEmail(file *ini.File) {
	EmailHost = file.Section("email").Key("EmailHost").String()
	EmailAddr = file.Section("email").Key("EmailAddr").String()
	EmailPassword = file.Section("email").Key("EmailPassword").String()
}

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败")
	}

	LoadEmail(file)
	LoadMysql(file)

	dns := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")

	model.Database(dns)
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
