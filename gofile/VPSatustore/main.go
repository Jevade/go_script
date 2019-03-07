package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
     connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
      db, err := gorm.Open("mysql", connArgs)
      if err != nil {
            log.Fatal(err)
      }
       db.SingularTable(true)
        return db

}

func init() {
	db = DbConn("root","123456","10.90.226.243","vps_information",3306)
	if !db.HasTable(&VPS{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&VPS{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table VPS created")
	} else {
		log.Println("Table VPS has exists")
	}

	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table User created")
	} else {
		log.Println("Table User has exists")
	}

	if !db.HasTable(&VPSOwn{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&VPSOwn{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table VPSOwn created")
	} else {
		log.Println("Table VPSOwn has exists")
	}

	if !db.HasTable(&Network{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Network{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table Network created")
	} else {
		log.Println("Table Network has exists")
	}

	if !db.HasTable(&Image{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Image{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table Image created")
	} else {
		log.Println("Table Image has exists")
	}
	// db.DropTableIfExists(&UsageInfo{})
	if !db.HasTable(&UsageInfo{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UsageInfo{}).Error; err != nil {
			panic(err)
		}
		log.Println("Table UsageInfo created")
	} else {
		log.Println("Table UsageInfo has exists")
	}
}

func main() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	ip :=[4]string{ "127.0.0.1","10.90.226.97","10.90.226.189","10.90.226.243"}
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		GetVPSInfo(uint64(i%len(ip)), ip[i%len(ip)])
	}
}

//GetVPSInfo return vps info
func GetVPSInfo(vpsID uint64, ip string) {
	resp, err := http.Get("http://" + ip + ":6669/sd/info")
	if err != nil {
		fmt.Println(err)
        return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var in map[string]interface{}
	json.Unmarshal(data, &in)
	for k, v := range in {
		switch v.(type) {
		case float32:
			in[k] = float64(v.(float32))
		default:

		}
		fmt.Println(k, ":", v)
	}
	usageInfo := &UsageInfo{
		VPSID: vpsID,
		L1:    (in["l1"]).(float64),
		L5:    (in["l5"]).(float64),
		L15:   (in["l15"]).(float64),
		CPU:   (in["busyCPU"]).(float64),
		Disk:  (in["usedDisk"]).(float64),
		MEM:   (in["usedMEM"]).(float64),
	}
	db.Save(usageInfo)

}
