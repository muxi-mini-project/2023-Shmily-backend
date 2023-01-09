package main

import (
	"2023-Shmily-backend/models/artical"
	"2023-Shmily-backend/models/user"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"time"
)

var router = mux.NewRouter()
var db *sql.DB

func initDB() {

	var err error
	// 设置数据库连接信息
	//注意： 请按你的实际情况修改数据库连接信息。
	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "shmily",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())
	checkError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := db.Exec(createArticlesSQL)
	checkError(err)
}

func saveArticleToDB(title string, body string) (int64, error) {

	// 变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	// 1. 获取一个 prepare 声明语句
	stmt, err = db.Prepare("INSERT INTO articles (title, body) VALUES(?,?)")
	// 例行的错误检测
	if err != nil {
		return 0, err
	}

	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 连接
	defer stmt.Close()

	// 3. 执行请求，传参进入绑定的内容
	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	// 4. 插入成功的话，会返回自增 ID
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, err
}

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	//.
	//.
	//.
	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
