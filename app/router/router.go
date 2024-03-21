package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xxvote/app/logic"
)

func New() {
	g := gin.Default()
	g.LoadHTMLGlob("app/view/*")

	g.GET("/", logic.Index)

	g.GET("/login", logic.LoginGet)
	g.POST("/login", logic.LoginPost)

	err := g.Run(":8090")
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}
}
