package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"xxvote/app/model"
	"xxvote/app/tools"
)

func AddVote(context *gin.Context) {
	idStr := context.Query("title")
	optStr, _ := context.GetPostFormArray("opt_name[]")
	//构建结构体
	vote := model.Vote{
		Title:       idStr,
		Type:        0,
		Status:      0,
		CreatedTime: time.Now(),
	}

	opt := make([]model.VoteOpt, 0)
	for _, v := range optStr {
		opt = append(opt, model.VoteOpt{
			Name:        v,
			CreatedTime: time.Now(),
		})
	}

	if err := model.AddVote(vote, opt); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, tools.OK)
	return
}

func UpdateVote(context *gin.Context) {
	idStr := context.PostForm("id")
	title := context.PostForm("title")
	idInt, _ := strconv.ParseInt(idStr, 10, 64)
	voteWithOpt := model.GetVote(idInt)
	fmt.Printf("%+v", voteWithOpt)
	voteWithOpt.Vote.Title = title

	err := model.UpdateVote(voteWithOpt.Vote, voteWithOpt.Opt)
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.OK)
}

func DeleteVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	if err := model.DeleteVote(id); err != true {
		context.JSON(http.StatusOK, tools.ECode{
			Code: 10006,
		})
		return
	}

	context.JSON(http.StatusOK, tools.OK)
	return
}
