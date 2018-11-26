package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	//"github.com/spf13/viper"
	//"../../model"
	"../../handler"
	"../../model"
	"../../pkg/errno"
)

//Info return info
func Info(c *gin.Context) {
	username := c.Param("username")
	handler.SendResponse(c, nil, CreateResponse{Username: username})
}

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [post]
func Create(c *gin.Context) {
	log.Info("Create user")

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Debugf("Username is: [%s]", r.Username)
	if err := r.checkParam(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
