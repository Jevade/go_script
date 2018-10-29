package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	//"github.com/spf13/viper"
	//"../../model"
	"../../handler"
	"../../pkg/errno"
)

//Info return info
func Info(c *gin.Context) {
	username := c.Param("username")
	handler.SendResponse(c, nil, CreateResponse{Username: username})
}

//Create is to create user
func Create(c *gin.Context) {
	log.Info("Create user")

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username is: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc is: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type is: %s", contentType)

	log.Debugf("Username is: [%s],Password is: [%s]", r.Username, r.Password)
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can't found in db:xx.xx.xx.xx")).Add("This is added message"), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("Password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)

}
