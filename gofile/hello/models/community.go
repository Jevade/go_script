package models

import "time"

const (
	COMMUNITY_CATE_COM = 0x01
)

//Community is to store group
type Community struct {
	Id   int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Name string `xorm:"varchar(20)" form:"name" json:"name"`
	//群主Id
	Ownerid int64  `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Icon    string `xorm:"varchar(120)" form:"icon" json:"icon"`
	Cate    int    `xorm:"int(10)" form:"cate" json:"cate"`
	//Memo
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime"  form:"createat" json:"createat"`
}
