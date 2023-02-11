package model

func migration() {
	//自动迁移、创建数据库表
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Memo{})
	DB.Model(&Memo{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")

}
