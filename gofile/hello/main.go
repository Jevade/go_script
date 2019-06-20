package main

import (
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
		// fmt.Println(tpl.Name())
		http.HandleFunc(tpl.Name(), func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, tpl.Name(), nil)
		})
	}
}

func main() {
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	RegisterViews()
	http.ListenAndServe(":8000", nil)
}
