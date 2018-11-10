package token

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	//ErrMissingHeader  means "Authorization" is missing
	ErrMissingHeader = errors.New("The length of the 'Authorization' header is zero")
)

//Context is context of the ison web token
type Context struct {
	ID       uint64
	Username string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

//Parse will
func Parse(tokenString, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, _ := strconv.Atoi(claims["id"].(string))
		ctx.ID = uint64(id)
		ctx.Username = claims["username"].(string)
		return ctx, nil

	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	secret := viper.GetString("jwt_secret")

	if len(secret) == 0 {
		return &Context{}, ErrMissingHeader
	}
	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	if len(secret) == 0 {
		secret = viper.GetString("jwt_secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.Username,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}
