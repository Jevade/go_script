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

//Create is to create user
// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {
	log.Info("Get user")
	username := c.Param("username")
	u, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, u)
}
