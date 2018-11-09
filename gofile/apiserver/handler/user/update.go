package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	//"github.com/spf13/viper"
	//"../../model"
	"../../handler"
	"../../model"
	"../../pkg/errno"
)

//Update is to create user
// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/user/{id} [put]
func Update(c *gin.Context) {
	log.Info("Create user")
	userID, _ := strconv.Atoi(c.Param("id"))
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.ID = uint64(userID)

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: u.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
