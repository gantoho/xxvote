package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote/app/model"
)

func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)
	ret := make(map[string]any)
	ret = model.GetUser(&user)
	context.JSON(http.StatusOK, ret)
}
