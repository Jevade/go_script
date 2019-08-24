package ctrl

import (
	"fmt"
	"net/http"

	"../args"
	"../service"
	"../util"
)

var communityService service.CommunityService

//JoinCommunity 处理用户好友列表
func JoinCommunity(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	err := communityService.Add(&arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOK(w, nil)
	}
	return
}

//LoadCommunity 处理用户好友列表
func LoadCommunity(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	community, err := communityService.LoadCommunity(&arg)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	fmt.Println(community, len(community))
	util.RespOKList(w, community, len(community))
}

//CreateCommunity 处理群组注册逻辑
func CreateCommunity(w http.ResponseWriter,
	r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name")
	memo := r.PostForm.Get("memo")
	icon := r.PostForm.Get("name")
	ownerid := r.PostForm.Get("ownerid")
	cate := r.PostForm.Get("cate")
	// nickname := r.PostForm.Get("nickname")
	community, err := communityService.CreateCommunity(name, memo, icon, cate, ownerid)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOK(w, community)
	}
	return
}
