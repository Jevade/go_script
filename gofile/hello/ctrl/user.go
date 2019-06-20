package ctrl

import (
	"errors"
	"net/http"

	"../models"
	"../service"
	"../util"
)

var userService service.UserService

//UserLogin 处理用户登录逻辑
func UserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("passwd")
	user, err := userService.Login(mobile, password)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		data := make(map[string]interface{})
		data["id"] = user.Id
		data["token"] = user.Token
		util.RespOK(w, data)
	}
}

//UserRegister 处理用户注册逻辑
func UserRegister(w http.ResponseWriter,
	r *http.Request) {
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	plainpwdagain := r.PostForm.Get("passwdagain")
	if plainpwd != plainpwdagain {
		util.RespFail(w, errors.New("密码不一致").Error())
		return
	}
	// nickname := r.PostForm.Get("nickname")
	nickname := "nickname"
	sex := models.SEX_UNKNOW
	user, err := userService.Register(mobile, plainpwd, nickname, "", sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOK(w, user)
	}
	return
}
