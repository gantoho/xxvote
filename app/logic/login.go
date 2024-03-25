package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
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
		context.JSON(http.StatusOK, tools.ECode{
			Message: err.Error(), // 这里有风险
		})
	}
	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.UserErr)
		return
	}

	// context.SetCookie("name", user.Name, 3600, "/", "localhost", false, true)
	// context.SetCookie("id", fmt.Sprint(ret.Id), 3600, "/", "localhost", false, true)

	_ = model.SetSession(context, user.Name, ret.Id)

	context.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
	})
	return
	//context.Redirect(http.StatusMovedPermanently, "index")
}

func Logout(context *gin.Context) {
	//context.SetCookie("name", "", 0, "/", "localhost", false, true)
	//context.SetCookie("id", "", 0, "/", "localhost", false, true)

	_ = model.FlushSession(context)
	context.Redirect(http.StatusFound, "/login")
}

type CUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Password2 string `json:"password_2"`
}

func CreateUser(context *gin.Context) {
	var user CUser
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(), // 这里有风险
		})
		return
	}

	if user.Name == "" || user.Password == "" || user.Password2 == "" {
		context.JSON(http.StatusOK, tools.ParamErr)
		return
	}

	if user.Password != user.Password2 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "两次密码不一致",
		})
		return
	}

	//这里有一个巨大的BUG，并发安全！！！
	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10004,
			Message: "用户已存在",
		})
		return
	}

	nameLen := len(user.Name)
	if nameLen < 5 || nameLen > 16 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "用户名长度不符合要求[5,16]",
		})
		return
	}

	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(user.Password) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "密码不能为纯数字",
		})
		return
	}

	newUser := model.User{
		Name:        user.Name,
		Password:    user.Password,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err = model.CreateUse(&newUser)
	if err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "用户注册失败",
		})
		return
	}

	context.JSON(http.StatusOK, tools.OK)
	return
}
