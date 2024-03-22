package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xxvote/app/model"
)

func Index(context *gin.Context) {
	ret := model.GetVotes()
	context.HTML(http.StatusOK, "index.html", gin.H{"vote": ret})
}

func GetVoteInfo(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	voteWithOpt := model.GetVote(id)
	context.HTML(http.StatusOK, "vote.html", gin.H{"voteWithOpt": voteWithOpt})
}

func PostVote(context *gin.Context) {
	//voteOptId := context.PostForm("voteOptId")
}
