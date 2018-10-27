package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"./config"
	"./model"
	"./router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config files path")
)

func main() {
	//读取配置文件，初始化链接选项
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	//初始化数据库连接
	model.DB.Init()
	defer model.DB.Close()

	//设置运行模式
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	middlerwares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlerwares...,
	)
	go func() {
		if err := PingServer(); err != nil {
			log.Fatal("The router no response,or it might took too long to start up.", err)
		}
		log.Info("The router has been started successfully")
	}()

	log.Infof("Start to listening the incoming  http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(":"+viper.GetString("addr"), g).Error())
	// log.Printf(http.ListenAndServe(":8080", g))
}

//PingServer is to test
func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		value := viper.GetFloat64("value")
		fmt.Println("Value is :", value)
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("it's ok")
			return nil
		}
		log.Info("Waitting for the router,retry in 1 second ")
		time.Sleep(time.Second)
	}
	//log.Fatal()	会跟着推迟，退出符号1
	return errors.New("Cannot connect to the router")
}
