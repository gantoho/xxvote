package app

import (
	"xxvote/app/model"
	"xxvote/app/router"
	"xxvote/app/schedule"
	"xxvote/app/tools"
)

func Start() {
	model.NewMysql()
	defer func() {
		model.Close()
	}()
	schedule.Start()
	tools.NewLogger()
	router.New()
}
