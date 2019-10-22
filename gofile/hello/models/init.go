package models

import (
	"errors"
	"fmt"
	"log"

	//mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	//presgres
	_ "github.com/lib/pq"
)

//DbEngin is to link db
var DbEngin *xorm.Engine

func initPg() (*xorm.Engine, error) {
	drivename := "postgres"
	DsName := "user=liu dbname=chat sslmode=disable"
	log.Println("Connect to postgres databases", DsName)
	pgDbEngin, err := xorm.NewEngine(drivename, DsName)
	return pgDbEngin, err

}

func initMs() (*xorm.Engine, error) {
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
	//DsName := "root:123456@tcp(192.168.199.167:3306)/chat?charset=utf8mb4"
	DsName := "root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4"
	log.Println("Connect to mysql databases", DsName)
	msDbEngin, err := xorm.NewEngine(drivename, DsName)
	return msDbEngin, err
}

func init() {
	//DbEngin, err = initMs()
	err := errors.New("")
	DbEngin, err = initPg()
	if err != nil {
		log.Println("failed")
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		log.Println(i)
		if DbEngin.Ping() != nil {
			fmt.Println("ping连接失败")
		}
		fmt.Println("ping连接成功")
	}
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
	if err != nil {
		log.Println("init database failed")
	}
	log.Println("init database success")
}

func errDefine(err error) {
	if err != nil {
		fmt.Println("here")
		log.Println(err)
	}
}
