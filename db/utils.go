package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(host string, port string, user string, password string, database string) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", user, password, host, port, database)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	return conn
}

func InitDB(user string, password string, host string, port string, database string) {
	conn := NewDB(host, port, user, password, database)
	conn.AutoMigrate(&Author{})
	conn.AutoMigrate(&Tag{})
	conn.AutoMigrate(&Post{})
}
