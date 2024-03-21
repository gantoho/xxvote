package model

import "fmt"

func GetVotes() []Vote {
	ret := make([]Vote, 0)
	err := Conn.Table("vote").Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
