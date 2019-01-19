package main

import (
	"github.com/astaxie/beego"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main2() {
	beego.Run()
}
func main() {
	// 打开数据库，sns是我的数据库名字，需要替换你自己的名字，（官网给的没有加tcp，跑不起来，具体有时 间看看源码分析下为何）
	// db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/godmin?charset=utf8mb4")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// topic是我本地数据库的表名，需要替换你自己的表名，这里面的英文注释都是引用github官网的~~
	//  嘿嘿 我只是想跑起来看看
	// rows, err := db.Query("SELECT * FROM goadmin_roles")
	// if err != nil {
	// panic(err.Error())
	// }
	// fmt.Println(rows)

	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "localhost",
				PORT:         "3306",
				USER:         "root",
				PWD:          "123456",
				NAME:         "goadmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "localhost", // 是cookie相关的，访问网站的域名
		PREFIX: "admin",
		// STORE 必须设置且保证有写权限，否则增加不了新的管理员用户
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
	}

	// Generators： 详见 https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// eng.AddConfig(cfg).Use(r)
	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	r.Run(":9033")
}
