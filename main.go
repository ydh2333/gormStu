package main

import (
	_ "github.com/go-sql-driver/mysql"
	gormstu "github.com/ydh2333/gormStu/homeWork/gormStu"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormstu.Run(db)
}

// sqlx
// func main() {
// 	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := sqlx.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err) // 实际项目中需优雅处理错误
// 	}
// 	defer db.Close()

// 	sqlxstu.Run(db)
// }
