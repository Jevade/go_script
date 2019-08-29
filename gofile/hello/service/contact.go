package service

import (
	"errors"
	"time"

	"../args"
	"../models"
)

type ContactService struct {
}

func (s *ContactService) LoadFriend(args *args.ContactArg) ([]map[string]interface{}, error) {
	userid := args.Userid
	// session := models.DbEngin.NewSession()
	contacts := make([]models.Contact, 0)
	objIds := make([]int64, 0)
	err := models.DbEngin.Where("ownerid=?", userid).And("cate=?", models.CONTACT_CATE_USER).Find(&contacts)
	if err != nil {
		return nil, err
	}
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}
	friends := make([]models.User, 0)
	result := make([](map[string]interface{}), 0)
	if 0 == len(objIds) {
		return result, err
	}
	models.DbEngin.In("id", objIds).Find(&friends)
	for _, friend := range friends {
		result = append(result, friend.Json())
	}
	return result, err
}

//Add 添加好友
func (s *ContactService) Add(args *args.ContactArg) (err error) {
	userid := args.Userid
	dstid := args.Dstid
	if userid == dstid {
		err = errors.New("不能添加自己")
		return
	}

	//查看是否存在需添加的人
	dstobj := models.User{}
	_, err = models.DbEngin.Where("id=?", dstid).Get(&dstobj)
	if dstobj.Id == 0 {
		err = errors.New("不存在")
		return
	}

	//查看是否已经添加
	contactobj := models.Contact{}
	_, err = models.DbEngin.Where("ownerid=?", userid).And("dstobj=?", dstid).And("cate=?", models.CONTACT_CATE_USER).Get(&contactobj)
	if contactobj.Id > 0 {
		err = errors.New("已经添加过好友")
		return
	}

	//开启事务
	session := models.DbEngin.NewSession()
	session.Begin()
	_, e2 := session.InsertOne(&models.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     models.CONTACT_CATE_USER,
		Createat: time.Now(),
	})
	_, e3 := session.InsertOne(&models.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     models.CONTACT_CATE_USER,
		Createat: time.Now(),
	})

	//没有错误
	if e2 == nil && e3 == nil {
		session.Commit()
	} else {
		session.Rollback()
		if e2 != nil {
			err = e2
		} else {
			err = e3
		}
	}
	return
}
