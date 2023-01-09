package bootstrap

import (
	article "2023-Shmily-backend/models/artical"
	"2023-Shmily-backend/models/category"
	"2023-Shmily-backend/models/user"
	"2023-Shmily-backend/pkg/config"
	"2023-Shmily-backend/pkg/model"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}
