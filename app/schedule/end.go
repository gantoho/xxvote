package schedule

import (
	"time"
	"xxvote/app/model"
)

func Start() {
	go voteEnd()
	return
}

func voteEnd() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			//fmt.Printf("定时器 voteEnd 启动")
			_ = model.EndVote()
		}
	}
}
