package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "Admin123_", "localhost:3306", "vote")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}
	Conn = conn
}

func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
}
