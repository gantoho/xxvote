package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote/app/logic"
	"xxvote/app/model"
	"xxvote/app/tools"
)

func New() {
	g := gin.Default()
	g.LoadHTMLGlob("app/view/*")

	{
		index := g.Group("")
		index.Use(checkUser)
		// vote
		index.GET("/index", logic.Index)

		index.GET("/votes", logic.GetVotes)
		index.GET("/vote", logic.GetVoteInfo)
		index.POST("/vote", logic.PostVote)

		index.POST("/vote/add", logic.AddVote)
		index.POST("/vote/update", logic.UpdateVote)
		index.POST("/vote/del", logic.DeleteVote)

		index.GET("/result", logic.ResultInfo)
		index.GET("/result/info", logic.ResultVote)
	}

	g.GET("/", logic.Index)

	{
		g.GET("/logout", logic.Logout)
		g.GET("/login", logic.LoginGet)
		g.POST("/login", logic.LoginPost)
	}

	err := g.Run(":8090")
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}
}

func checkUser(context *gin.Context) {
	var name string
	var id int64
	values := model.GetSession(context)
	if v, ok := values["name"]; ok {
		name = v.(string)
	}
	if v, ok := values["id"]; ok {
		id = v.(int64)
	}
	if name == "" || id <= 0 {
		context.JSON(http.StatusOK, tools.NotLogin)
		context.Redirect(http.StatusFound, "/login")
		context.Abort()
	}
	context.Next()
}
