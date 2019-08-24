package models

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

//DbEngin is to link db
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
	fmt.Println(123)
	DsName := "root:123456@tcp(192.168.199.132:3306)/chat?charset=utf8mb4"
	fmt.Println(DsName)
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect to databases", DsName)
	//是否显示sql
	// GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION;
	// update user set host='%' where user='root';
	DbEngin.ShowSQL(true)
	//数据库最大连接数
	DbEngin.SetMaxOpenConns(2)
	err = DbEngin.Sync2(
		new(User),
		new(Community),
		new(Contact))
	errDefine(err)
	log.Println("init database success")
}

func errDefine(err error) {
	if err != nil {
		fmt.Println("here")
		log.Println(err)
	}
}
