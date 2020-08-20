package main

import (
	"github.com/jevade/hello/util"
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
	http.HandleFunc("/data", MSdataFunc)
    fmt.Println("start listen")
	http.ListenAndServe(":8787", nil)
}

func MSdataFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Println("get listen")
	connStr := "user=hailan dbname=hailan port=5432 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed")
		return
	}
	//查询数据
	rows, err := db.Query("SELECT id,sourcename,type,c1,c2 FROM sourcecollectinfos where type='NS'")
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

	util.RespOK(w, data)
}

