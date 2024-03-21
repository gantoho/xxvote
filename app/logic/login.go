package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote/app/model"
	"xxvote/app/tools"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"pass" form:"password"`
}

func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(context *gin.Context) {
	var user User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusOK, tools.Ecode{
			Message: err.Error(), // 这里有风险
		})
	}
	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.Ecode{
			Message: "账号或密码错误",
		})
		return
	}
	context.SetCookie("name", user.Name, 3600, "/", "localhost", false, true)
	context.JSON(http.StatusOK, tools.Ecode{
		Message: "登录成功",
	})
	//context.Redirect(http.StatusMovedPermanently, "index")
}
