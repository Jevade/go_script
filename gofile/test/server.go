package main

import (
	"fmt"
	"net/http"
	"text/template"
	"reflect"
	"./point"
	"./employee"
	"./car"
	"./vendo"
	"./integer"
)

type Address struct{
	code string "mail code of address"
	street string "street of address"
	city string "city of address"
	country string "country of address"
}

type VCard struct{
	adress *Address
	name string
	telNum string
	corpation string
	age uint16
}
func print(T *VCard){
	fmt.Println(*T.adress)
	fmt.Println(T.age)
	fmt.Println(T.corpation)
}
func main(){
	myVCard := new(VCard)
	myVCard.adress = new(Address)
	myVCard.adress.city="bj"
	myVCard.adress.country="CN"
	myVCard.adress.street="YCBYJ"
	myVCard.adress.code="102201"
	myVCard.age=25
	myVCard.corpation = "COMAC"
	myVCard.name = "Jevade"
	myVCard.telNum = "13112345678"
	print(myVCard)
	p1 := new(point.Point2D)
	p1.Set(2,3)
	p2 := &point.Point2D{3,5}
	fmt.Println(p1.ABS2D())
	p1.Scale(2)
	fmt.Println(p1.ABS2D())
	fmt.Println(p1.Dis(p2))
	p3 :=  point.Point3D{1,point.Point2D{2,3}}
    fmt.Println((&p3).ABS2D())
	http.HandleFunc("/",root)
	http.HandleFunc("/long",long)
	// http.ListenAndServe("localhost:8080",nil)
	reflectType(*(myVCard.adress),1)
	employee := employee.NewEmployee(1000,"li",1980,1)
	employee.GiveRaise(1.2,1)
	fmt.Println(employee.Age(),employee.Salary())
	car := car.NewMecideris(4,"MD")
	car.Start()
	fmt.Println(car.NumberOfWheel())

	vendo := vendo.Vendo{}
	vendo.Magic()
	vendo.MoreMagic()
	integer.Test()
	integer.Test2()
}
func reflectType(tt Address,idx int){
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(idx)
	fmt.Printf("%v\n",ixField.Tag)
}
var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))
func root(w http.ResponseWriter,r *http.Request){
	fmt.Println(r)
	rootHtmlTmpl.Execute(w,nil)
}
func long(w http.ResponseWriter,r *http.Request){
	shortUrl := r.FormValue("shortUrl")

	longUrl := "long__" + shortUrl
	rootHtmlTmpl.Execute(w,fmt.Sprintf("Longer version of %s is : %s",shortUrl,longUrl))
}
