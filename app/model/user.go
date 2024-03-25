package model

import (
	"fmt"
)

func GetUser(name string) *User {
	var ret User
	err := Conn.Table("user").Where("name = ?", name).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}

func CreateUse(user *User) error {
	if err := Conn.Table("user").Create(user).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
