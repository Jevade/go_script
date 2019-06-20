package models

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	// "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	//             dbConfig.User,
	//             dbConfig.Passwd,
	//             dbConfig.Host,
	//             dbConfig.Port,
	//             dbConfig.Name,
	//             dbConfig.ParseTime,
	//             dbConfig.Local)
	DsName := "root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if err != nil {
		log.Fatal(err)
	}
	//是否显示sql
	DbEngin.ShowSQL(true)
	//数据库最大连接数
	DbEngin.SetMaxOpenConns(2)
	err = DbEngin.Sync2(new(User))
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)

	}
	log.Println("init data success")

}
