package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(user string, password string, host string, port string, database string) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", user, password, host, port, database)
	fmt.Println(dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	conn.AutoMigrate(&Author{})
	conn.AutoMigrate(&Tag{})
	conn.AutoMigrate(&Post{})
}
