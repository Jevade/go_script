package ctrl

import (
	"fmt"
	"net/http"

	"../args"
	"../service"
	"../util"
)

var contactService service.ContactService

//AddFriend 处理用户登录逻辑
func AddFriend(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	fmt.Println(arg.Userid)
	err := contactService.Add(&arg)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	util.RespOK(w, nil)
}

//LoadFriend 处理用户好友列表
func LoadFriend(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	friends, err := contactService.LoadFriend(&arg)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	fmt.Println(friends, len(friends))
	util.RespOKList(w, friends, len(friends))
}
