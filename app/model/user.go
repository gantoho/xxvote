package model

import (
	"fmt"
	"time"
)

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name        string    `gorm:"column:name;default:NULL"`
	Password    string    `gorm:"column:password;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func GetUser(name string) *User {
	var ret User
	err := Conn.Table("user").Where("name = ?", name).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}
