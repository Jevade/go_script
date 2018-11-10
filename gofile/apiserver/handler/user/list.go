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
// @Param user body user.ListRequest true "List users"
// @Success 200
// @Router /user [get]
func GetList(c *gin.Context) {
	log.Info("Get user")
	var r ListRequest

	// if err := c.Bind(&r); err != nil {
	// 	handler.SendResponse(c, errno.ErrBind, nil)
	// 	return
	// }
	r.Username = c.Param("username")
	r.Offset, _ = strconv.Atoi(c.Param("offset"))
	r.Limit, _ = strconv.Atoi(c.Param("limit"))

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
