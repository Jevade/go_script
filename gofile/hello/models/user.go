package models

import (
	"time"
)

const (
	SEX_FEMALE = "F"
	SEX_MALE   = "M"
	SEX_UNKNOW = "U"
)

//User
type User struct {
	Id       int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"passwd"`
	Avatar   string `xorm:"varchar(120)" form:"avatar" json:"avatar"`
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt     string `xorm:"varchar(10)" form:"salt" json:"salt"`
	Online   int    `xorm:"int(10)" form:"online" json:"online"`
	Token    string `xorm:"varchar(40)" form:"token" json:"token"`
	Memo     string `xorm:"varchar(140)" form:"memo" json:"memo"`
	//Memo
	Createat time.Time `xorm:"datetime"  form:"createat" json:"createat"`
}