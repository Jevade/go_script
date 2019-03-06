package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./config"
	"./handler"
	"./model"
	v "./pkg/version"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "", "58spy config files path")
	version = pflag.BoolP("version", "v", false, "show version info")
	help    = pflag.StringP("help", "h", "", "58spy help info")

	gitTag       = ""
	gitCommit    = "$:%H$"
	gitTreeState = "not a git tree"
	buildDate    = "1970-01-01T00:00:00Z"
)

func main() {
	//读取配置文件，初始化链接选项
	pflag.Parse()
	if *help == "v" || *help == "version" {
		fmt.Println("-h --help 58spy help info")
		fmt.Println("-c --config 58spy config yaml/json")
		fmt.Println("-v --version 58spy versions")
	}
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
	model.DB.Init()
	// itemspy.GetItemInfo("https://bj.58.com/shouji/pn06/", make(chan interface{}, 4))
	// handler.ProcessType()
	handler.Process()
}
