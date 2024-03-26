package logic

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	if ret.Id < 1 || ret.Password != encryptV1(user.Password) {
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

	//测试
	//encrypt(user.Password)
	//encryptV1(user.Password)
	//encryptV2(user.Password)
	//return

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

	//这里有一个巨大的BUG，并发安全！！！
	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10004,
			Message: "用户已存在",
		})
		return
	}

	newUser := model.User{
		Name:        user.Name,
		Password:    encryptV1(user.Password),
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

// 最基础的版本
func encrypt(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)
	return hashString
}

func encryptV1(pwd string) string {
	newPwd := pwd + "香香编程喵喵喵" //不能随便起，且不能暴露
	hash := md5.New()
	hash.Write([]byte(newPwd))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)
	return hashString
}

func encryptV2(pwd string) string {
	//基于Blowfish 实现加密。简单快速，但有安全风险
	//golang.org/x/crypto/ 中有大量的加密算法
	newPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败：", err)
		return ""
	}
	newPwdStr := string(newPwd)
	fmt.Printf("加密后的密码：%s\n", newPwdStr)
	return newPwdStr
}
