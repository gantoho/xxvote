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

func GetVotes(context *gin.Context) {
	ret := model.GetVotes()
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}

func GetVoteInfo(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	voteWithOpt := model.GetVote(id)

	//log.Printf("[Printf]ret:%+v\n", voteWithOpt)
	//log.Panicf("[Panicf]ret:%+v\n", voteWithOpt)
	//log.Fatalf("[Fatalf]ret:%+v\n", voteWithOpt)

	//logrus.Errorf("[Error]ret:%+v\n", voteWithOpt)
	tools.Logger.Error("[Error]ret:%+v\n", voteWithOpt)

	context.JSON(http.StatusOK, tools.ECode{
		Data: voteWithOpt,
	})
	//context.HTML(http.StatusOK, "vote.html", gin.H{"voteWithOpt": voteWithOpt})
}

func PostVote(context *gin.Context) {
	userIDStr, _ := context.Cookie("id")
	voteIdStr, _ := context.GetPostForm("vote_id")
	optStr, _ := context.GetPostFormArray("opt[]")

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	voteId, _ := strconv.ParseInt(voteIdStr, 10, 64)

	old := model.GetVoteHistory(userID, voteId)
	if len(old) >= 1 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "您已投过票了",
		})
		return
	}

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
