package app

import (
	"xxvote/app/model"
	"xxvote/app/router"
)

func Start() {
	model.NewMysql()
	defer func() {
		model.Close()
	}()
	router.New()
}
