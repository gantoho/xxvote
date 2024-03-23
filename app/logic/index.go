package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xxvote/app/model"
	"xxvote/app/tools"
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
	userIDStr, _ := context.Cookie("id")
	voteIdStr, _ := context.GetPostForm("vote_id")
	optStr, _ := context.GetPostFormArray("opt[]")

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	voteId, _ := strconv.ParseInt(voteIdStr, 10, 64)
	opt := make([]int64, 0)
	for _, v := range optStr {
		optId, _ := strconv.ParseInt(v, 10, 64)
		opt = append(opt, optId)
	}
	err := model.PostVote(userID, voteId, opt)
	if err != true {
		//context.JSON(http.StatusOK, gin.H{"code": 1, "msg": "投票失败"})
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1,
			Message: "投票失败",
		})
		return
	}
	context.JSON(http.StatusOK, tools.ECode{
		Message: "投票成功",
	})
}
