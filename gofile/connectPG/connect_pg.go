package main

import (
	"github.com/jevade/hello"
	"net/http"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type SourceCollection struct {
	id         int
	sourcename string
	sctype     string
	c1         float32
	c2         int
	c3         int
	c4         int
	c5         int
	c6         int
} //查询数据
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main(){
	http.HandleFunc("/user/login", MSdataFunc)
	http.ListenAndServe(":8787", nil)
}
//UserLogin 处理用户登录逻辑
func UserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println(" Doing login jobs")
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("password")
	if !(mobile != "" && password != "") {
		util.RespFail(w, errors.New("空数据").Error())
		return
	}
	log.Println(1222222, mobile, password)
	user, err := userService.Login(mobile, password)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		id := fmt.Sprintf("%d", user.Id)
		w.Header().Set("id", id)
		data := make(map[string]interface{})
		data["id"] = user.Id
		data["token"] = user.Token
		data["mobile"] = user.Mobile
		data["avatar"] = user.Avatar
		data["memo"] = user.Memo
		data["sex"] = user.Sex
		util.RespOK(w, data)
	}
}

func MSdataFunc(w http.ResponseWriter, r *http.Request) {
	connStr := "user=hailan dbname=hailan port=5432 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed")
		return
	}
	//查询数据
	rows, err := db.Query("SELECT id,sourcename FROM sourcecollectinfos where type='NS'")
	checkErr(err)

	data := make(map[string][]interface{})
	for rows.Next() {
		sourceCollection := SourceCollection{}
		err = rows.Scan(&sourceCollection.id, &sourceCollection.sourcename, &sourceCollection.sctype, &sourceCollection.c1, &sourceCollection.c2)
		checkErr(err)
		fmt.Println(sourceCollection.sourcename)
		data["id"]=append(data["id"],sourceCollection.id)
		data["sourcename"]=append(data["sourcename"],sourceCollection.sourcename)
		data["type"]=append(data["type"],sourceCollection.sctype)
		data["c1"]=append(data["c1"],sourceCollection.c1)
		data["c2"]=append(data["c2"],sourceCollection.c2)
	}

	hello.util.RespOK(w, data)
}

