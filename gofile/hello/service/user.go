package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"../models"
	"../util"
)

//UserService 用户服务
type UserService struct {
}

//Register 提供用户注册服务
func (s *UserService) Register(
	mobile, //手机号
	plainpwd, //明文密码
	nickname, //昵称
	avatar, sex string) (user models.User, err error) {
	//if mobile exists
	tmp := models.User{}
	_, err = models.DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//message exists
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经被占用")
		// return tmp,errors.New()
	}
	//create new user
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(100000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//insert
	_, err = models.DbEngin.InsertOne(&tmp)
	//前段恶意插入特殊zifu
	//数据库连接操作失败
	//
	if err != nil {
		return tmp, err
	}
	return tmp, nil
}

//提供用户登录服务
func (s *UserService) Login(
	mobile,
	plainpwd string) (user models.User, err error) {
	tmp := models.User{}
	_, err = models.DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}

	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	str := fmt.Sprintf("%v", time.Now())
	token := util.MD5Encode(str)
	tmp.Token = token
	models.DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}
