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

func ResultInfo(context *gin.Context) {
	context.HTML(http.StatusOK, "result.html", nil)
}

type ResultData struct {
	Title string
	Count int64
	Opt   []*ResultVoteOpt
}

type ResultVoteOpt struct {
	Name  string
	Count int64
}

func ResultVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	data := ResultData{
		Title: ret.Vote.Title,
	}

	for _, v := range ret.Opt {
		data.Count = data.Count + v.Count
		tmp := ResultVoteOpt{
			Name:  v.Name,
			Count: v.Count,
		}
		data.Opt = append(data.Opt, &tmp)
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: data,
	})
}
