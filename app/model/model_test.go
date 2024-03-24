package model

import (
	"fmt"
	"testing"
	"time"
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

func TestAddVote(t *testing.T) {
	NewMysql()
	vote := Vote{
		Title:       "测试",
		Type:        0,
		Status:      0,
		Time:        0,
		UserId:      0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	opt := make([]VoteOpt, 0)
	opt = append(opt, VoteOpt{
		Name:        "测试选项1",
		Count:       0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	opt = append(opt, VoteOpt{
		Name:        "测试选项2",
		Count:       0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	err := AddVote(vote, opt)
	fmt.Printf("err:%s", err)
	Close()
}
