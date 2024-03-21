package model

import (
	"fmt"
)

type User struct {
	UserName string `json:"name" form:"username"`
	Password string `json:"pass" form:"password"`
}

func GetUser(user *User) map[string]any {
	ret := make(map[string]any)
	err := Conn.Table("user").Where("name = ?", user.UserName).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
