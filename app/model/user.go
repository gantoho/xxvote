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
