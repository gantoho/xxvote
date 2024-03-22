package model

import (
	"fmt"
	"testing"
)

func TestGetVotes(t *testing.T) {
	NewMysql()
	// 测试用例
	r := GetVotes()
	fmt.Printf("GetVotes: %v+\n", r)
	Close()
}

func TestGetVote(t *testing.T) {
	NewMysql()
	r := GetVote(1)
	fmt.Printf("GetVote: %+v\n", r)
	Close()
}

func TestPostVote(t *testing.T) {
	NewMysql()
	r := PostVote(1, 1, []int64{1, 2})
	fmt.Printf("PostVote: %+v\n", r)
	Close()
}
