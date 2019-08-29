package ctrl

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"../models"
	"../service"
	"../util"
)

var userService service.UserService

//UserLogin 处理用户登录逻辑
func UserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println(" Doing login jobs")
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("password")
	if !(mobile != "" && password != "") {
		util.RespFail(w, errors.New("空数据").Error())
		return
	}
	log.Println(1222222, mobile, password)
	user, err := userService.Login(mobile, password)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		id := fmt.Sprintf("%d", user.Id)
		w.Header().Set("id", id)
		data := make(map[string]interface{})
		data["id"] = user.Id
		data["token"] = user.Token
		data["mobile"] = user.Mobile
		data["avatar"] = user.Avatar
		data["memo"] = user.Memo
		data["sex"] = user.Sex
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
