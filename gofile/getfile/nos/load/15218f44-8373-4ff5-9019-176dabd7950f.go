package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//引入 mysql的驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	//"github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/jinzhu/gorm/dialects/mssql"
)

//DB over db conns
var DB = Database{}

//DBConfig store config params
type DBConfig struct {
	Type      string
	User      string
	Passwd    string
	Host      string
	Port      string
	Name      string
	ParseTime bool
	Local     string
}

//Database store config and db
type Database struct {
	Self           *gorm.DB
	Docker         *gorm.DB
	DBConfig       *DBConfig
	DockerDBConfig *DBConfig
}

//Init will get config file and init dbs
func (db *Database) Init() {
	DB.DBConfig = GetDBConfig()
	DB.DockerDBConfig = GetDockerDBConfig()
	DB.Self = GetSelfDB()
	DB.Docker = GetDockerDB()
	if viper.GetBool("db.createDB") {
		createDB(GetSelfDB())
	}
}

//Close will close all conns
func (db *Database) Close() {
	db.Self.Close()
	db.Docker.Close()
}

//GetDockerDBConfig return DockerDBConfig
func GetDockerDBConfig() *DBConfig {
	return &DBConfig{
		Type:      viper.GetString("docker_db.type"),
		User:      viper.GetString("docker_db.user"),
		Passwd:    viper.GetString("docker_db.passwd"),
		Host:      viper.GetString("docker_db.host"),
		Port:      viper.GetString("docker_db.port"),
		Name:      viper.GetString("docker_db.name"),
		ParseTime: viper.GetBool("docker_db.ParseTime"),
		Local:     viper.GetString("docker_db.local")}

}

//GetDBConfig return DB config
func GetDBConfig() *DBConfig {
	return &DBConfig{
		Type:      viper.GetString("db.type"),
		User:      viper.GetString("db.user"),
		Passwd:    viper.GetString("db.passwd"),
		Host:      viper.GetString("db.host"),
		Port:      viper.GetString("db.port"),
		Name:      viper.GetString("db.name"),
		ParseTime: viper.GetBool("db.ParseTime"),
		Local:     viper.GetString("db.local")}

}

//GetSelfDB return SelfDB
func GetSelfDB() *gorm.DB {
	db := openDB(DB.DBConfig)
	return db
}

//GetDockerDB return DockerDB
func GetDockerDB() *gorm.DB {
	db := openDB(DB.DockerDBConfig)
	return db
}

//openDB will open and return db with gifted config
func openDB(dbConfig *DBConfig) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		dbConfig.User,
		dbConfig.Passwd,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.ParseTime,
		dbConfig.Local)
	db, err := gorm.Open(dbConfig.Type, config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", dbConfig.Name)
	}
	log.Info("Connect to db successfully!!!!!!!!!!!!")
	setupDB(db)
	return db
}

//setupDB set some db property with config file
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(viper.GetInt("gorm_max_conns")) //最大连接数
	db.DB().SetMaxIdleConns(0)                              //限制连接数，放入连接池中，下一次使用
}

func createDB(db *gorm.DB) {
	if !db.HasTable(&ItemInfo{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ItemInfo{}).Error; err != nil {
			panic(err)
		}
		log.Info("Table ItemInfo created")
	} else {
		log.Info("Table ItemInfo has exists")
	}
	if !db.HasTable(&TypeInfo{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&TypeInfo{}).Error; err != nil {
			panic(err)
		}
		log.Info("Table ItemInfo created")
	} else {
		log.Info("Table ItemInfo has exists")
	}
	if !db.HasTable(&CityInfo{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CityInfo{}).Error; err != nil {
			panic(err)
		}
		log.Info("Table CityInfo created")
	} else {
		log.Info("Table CityInfo has exists")
	}

}
