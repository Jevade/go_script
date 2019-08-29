package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./ctrl"
	_ "./models"
	_ "github.com/go-sql-driver/mysql"
)

//全局扫描模板
func RegisterViews() {
	tpls, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(tpls.Templates())
	for _, tpl := range tpls.Templates() {
		v := tpl.Name()
		log.Println(v)
		http.HandleFunc(v, func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(tpl.Name())
			tpl.ExecuteTemplate(w, v, nil)
		})
	}
}

// 函数式编程，闭包外变量被 局部函数多次引用时，各闭包的值会同步，
//可以实现数据实时交换，如果需要不同的值，需要赋值先取出来，在使用取出的值引入闭包中

func main() {
	log.Println(" I am doing jobs")
	RegisterViews()
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/contact/addfriend", ctrl.AddFriend)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.HandleFunc("/chat", ctrl.Chat)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
