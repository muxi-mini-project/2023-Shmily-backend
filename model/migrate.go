package model

func migration() {
	//自动迁移、创建数据库表
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Memo{}).AutoMigrate(&Letter{}).AutoMigrate(&Event{}).AutoMigrate(&AboutLover{}).AutoMigrate(&Friend{})

	DB.Model(&Memo{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")

	DB.Model(&Letter{}).AddForeignKey("FromUid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&Letter{}).AddForeignKey("ToUid", "User(id)", "CASCADE", "CASCADE")

	DB.Model(&Event{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")

	DB.Model(&AboutLover{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")
	DB.Model(&AboutLover{}).AddForeignKey("LoverUid", "User(id)", "CASCADE", "CASCADE")

	//DB.Model(&Friend{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")
	//DB.Model(&Friend{}).AddForeignKey("FriendUid", "User(id)", "CASCADE", "CASCADE")
}
