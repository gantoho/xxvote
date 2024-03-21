package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote/app/model"
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
		context.JSON(http.StatusOK, map[string]string{
			"msg": "参数错误",
		})
	}
	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, map[string]string{
			"msg": "用户名或密码错误",
		})
		return
	}
	context.SetCookie("name", user.Name, 3600, "/", "localhost", false, true)
	//context.JSON(http.StatusOK, ret)
	context.Redirect(http.StatusMovedPermanently, "index")
}
