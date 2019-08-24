package service

import (
	"errors"
	"strconv"
	"time"

	"../args"
	"../models"
)

//CommunityService is to undertake  CommunityService
type CommunityService struct {
}

//Add is to  添加
func (s *CommunityService) Add(args *args.ContactArg) (err error) {
	userid := args.Userid
	dstid := args.Dstid

	//查看是否存在需添加的人
	dstobj := models.Community{}
	_, err = models.DbEngin.Where("id=?", dstid).Get(&dstobj)
	if dstobj.Id == 0 {
		err = errors.New("不存在该群")
		return
	}

	//查看是否已经添加
	contactobj := models.Contact{}
	_, err = models.DbEngin.Where("ownerid=?", userid).And("dstobj=?", dstid).And("cate=?", models.CONTACT_CATE_COMMUNITY).Get(&contactobj)
	if contactobj.Id > 0 {
		err = errors.New("已经添加过群")
		return
	}

	//开启事务
	session := models.DbEngin.NewSession()
	session.Begin()
	_, err = session.InsertOne(&models.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     models.CONTACT_CATE_COMMUNITY,
		Createat: time.Now(),
	})
	//没有错误
	if err != nil {
		session.Rollback()
	}
	session.Commit()
	return
}

//LoadCommunity is to  添加
func (s *CommunityService) LoadCommunity(args *args.ContactArg) ([]models.Community, error) {
	userid := args.Userid
	// session := models.DbEngin.NewSession()
	contacts := make([]models.Contact, 0)
	objIds := make([]int64, 0)
	err := models.DbEngin.Where("ownerid=?", userid).And("cate=?", models.CONTACT_CATE_COMMUNITY).Find(&contacts)
	if err != nil {
		return nil, err
	}
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}
	communities := make([]models.Community, 0)
	if 0 == len(objIds) {
		return communities, err
	}
	models.DbEngin.In("id", objIds).Find(&communities)
	return communities, err
}

//CreateCommunity is to  添加群
func (s *CommunityService) CreateCommunity(
	name, //群名称
	memo, //群介绍
	icon, //图像
	cate, ownerid string) (community models.Community, err error) {
	//if mobile exists
	tmp := models.Community{}
	_, err = models.DbEngin.Where("name=?", name).And("ownerid=?", ownerid).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//message exists
	if tmp.Id > 0 {
		return tmp, errors.New("用户有同名群存在，请修改名称")
		// return tmp,errors.New()
	}
	//create new community
	tmp.Name = name
	tmp.Ownerid, err = strconv.ParseInt(ownerid, 10, 64)
	tmp.Icon = icon
	tmp.Cate = models.COMMUNITY_CATE_COM
	tmp.Memo = memo
	tmp.Createat = time.Now()
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
