package user

import (
    "fmt"
    "net/http"
    "github.com/lexkong/log"
	"github.com/gin-gonic/gin"
    //"github.com/spf13/viper"
    //"../../model"
    "../../pkg/errno"
)


//Create is check health
func Create(c *gin.Context) {
	log.Info("Create user")
    var r struct{
        Username string `json:"username"`
        Password string `json:"password"`
    }

    var err error
    if err = c.Bind(&r);err!=nil{
        c.JSON(http.StatusOK,gin.H{"error":errno.ErrBind})
        return
    }

    log.Debugf("Username is: [%s],Password is: [%s]",r.Username,r.Password)
    if r.Username==""{
        err = errno.New(errno.ErrUserNotFound,fmt.Errorf("username can't found in db:xx.xx.xx.xx")).Add("This is added message")
        log.Errorf(err,"Get an error")
    }

    if errno.IsErrUserNotFound(err){
        log.Debug("err type is ErrUserNotFound")
    }

    if r.Password ==""{
        err=fmt.Errorf("Password is empty")
    }

    code, message := errno.DecodeErr(err)
    c.JSON(http.StatusOK,gin.H{"code":code,"message":message})
}
