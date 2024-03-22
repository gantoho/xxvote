package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote/app/logic"
)

func New() {
	g := gin.Default()
	g.LoadHTMLGlob("app/view/*")

	index := g.Group("")
	index.Use(checkUser)
	index.GET("/index", logic.Index)
	index.GET("/vote", logic.GetVoteInfo)
	index.POST("/vote", logic.PostVote)

	g.GET("/", logic.Index)

	g.GET("/login", logic.LoginGet)
	g.POST("/login", logic.LoginPost)

	err := g.Run(":8090")
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}
}

func checkUser(context *gin.Context) {
	name, err := context.Cookie("name")
	if err != nil || name == "" {
		context.Redirect(http.StatusFound, "/login")
	}
	context.Next()
}
