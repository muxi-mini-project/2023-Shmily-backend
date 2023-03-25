package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"log"
	"os"
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

	SetupLogger()

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

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./shmily.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

type Config struct {
	Name string
}

func Initi(cfg string, prefix string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(prefix); err != nil {
		return err
	}

	return nil
}

func (c *Config) initConfig(prefix string) error {
	if c.Name != "" {
		// 如果制定配置文件，则解析配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 	如果没有制指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("/home/pro1/2023-Shmily-backend1/conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 默认配置文件类型为yaml
	viper.AutomaticEnv()        // 读取默认的环境变量
	viper.SetEnvPrefix(prefix)  // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
