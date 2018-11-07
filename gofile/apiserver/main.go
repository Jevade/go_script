package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"./config"
	"./model"
	v "./pkg/version"
	"./router"
	"./router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg          = pflag.StringP("config", "c", "", "apiserver config files path")
	version      = pflag.BoolP("version", "v", false, "show version info")
	xyz          string
	gitTag       string
	gitCommit    string = "$:%H$"
	gitTreeState string = "not a git tree"
	buildDate    string = "1970-01-01T00:00:00Z"
)

func main() {
	fmt.Println(xyz)
	v.PrintStr()
	//读取配置文件，初始化链接选项
	pflag.Parse()
	if *version {
		v := v.Get()
		v.GitTag = gitTag
		v.GitCommit = gitCommit
		v.GitTreeState = gitTreeState
		v.BuildDate = buildDate

		marshalled, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	//初始化数据库连接
	model.DB.Init()
	defer model.DB.Close()

	//设置运行模式
	gin.SetMode("debug")
	g := gin.New()
	//定义需要载入的全局中间件
	globalmiddlerwares := []gin.HandlerFunc{
		middleware.RequestID(),
		middleware.Logging(),
	}

	router.Load(
		g,
		//middleware
		globalmiddlerwares...,
	)
	go func() {
		if err := PingServer(); err != nil {
			log.Fatal("The router no response,or it might took too long to start up.", err)
		}
		log.Info("The router has been started successfully")
	}()
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Info("Start https server")
			log.Infof("Start to listening the incoming  http address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(":"+viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	log.Infof("Start to listening the incoming  http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(":"+viper.GetString("addr"), g).Error())
}

//PingServer is to test
func PingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		value := viper.GetFloat64("value")
		fmt.Println("Value is :", value)

		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("http it's ok")
			return nil
		}
		log.Info("Waitting for the router,retry in 1 second ")
		time.Sleep(time.Second)
	}
	//log.Fatal()	会跟着推迟，退出符号
	return errors.New("Cannot connect to the router")
}
