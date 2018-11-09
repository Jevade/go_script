package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	//"github.com/spf13/viper"
	//"../../model"
	"../../handler"
	"../../pkg/errno"
	"../../service"
)

//GetList is to create user
// @Summary List the users in the database
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param username query string false "name search" "username"
// @Param limit query int false "search limit" "limit"
// @Param offset query int false "serach offset"  "offset"
// @Success 200
// @Router /v1/user [get]
func GetList(c *gin.Context) {
	log.Info("Get user")
	var r ListRequest
	r.Username = c.Request.URL.Query().Get("username")
	r.Offset, _ = strconv.Atoi(c.Request.URL.Query().Get("Offset"))
	r.Limit, _ = strconv.Atoi(c.Request.URL.Query().Get("Limit"))

	userinfo, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	lr := ListResponse{
		TotalCount: count,
		UserList:   userinfo,
	}
	handler.SendResponse(c, nil, lr)
}
