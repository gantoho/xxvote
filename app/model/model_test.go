package model

import (
	"fmt"
	"testing"
)

func TestGetVotes(t *testing.T) {
	NewMysql()
	// 测试用例
	r := GetVotes()
	fmt.Printf("GetVotes: %v\n", r)
	Close()
}
