package handler

import (
	"net/http"

	"../pkg/errno"
	"github.com/gin-gonic/gin"
)

//Response define response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//SendResponse will modify gin.Context
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
