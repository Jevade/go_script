package models

import "time"

const (
	CONTACT_CATE_USER      = 0x01
	CONTACT_CATE_COMMUNITY = 0x02
)

//Contact is to store contact
type Contact struct {
	Id      int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Dstobj  int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	Cate    int   `xorm:"int(10)" form:"Ccte" json:"cate"`
	//Memo
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime"  form:"createat" json:"createat"`
}
