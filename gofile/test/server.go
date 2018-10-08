package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"text/template"

	"./car"
	"./employee"
	"./file"
	"./integer"
	"./point"
	"./vendo"
)

//Go中需要被外部访问到的值均需要首字母大写，
//小写的话是私有值，无法被外部访问

// Address store address
type Address struct {
	code    string "mail code of address"
	street  string "street of address"
	city    string "city of address"
	country string "country of address"
	Num     int    "num"
}

//VCard info
type VCard struct {
	adress    *Address
	name      string
	telNum    string
	corpation string
	age       uint16
}

func print(T *VCard) {
	fmt.Println(*T.adress)
	fmt.Println(T.age)
	fmt.Println(T.corpation)
}

// SaveJson will save []byte to json file
func SaveJson(js []byte, jsonName string) {
	file, _ := os.OpenFile(jsonName, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(js)
	if err != nil {
		log.Println("error in encoding json!", err)
	}
}

// ReadJson get json string
func ReadJson(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	dec := json.NewDecoder(file)
	var str []byte
	err := dec.Decode(&str)
	var in interface{}
	json.Unmarshal(str, &in)
	if err != nil {
		log.Println("decode error", err)
		return
	}
	fmt.Println(in)
}
func main() {
	myVCard := VCard{}
	myVCard.adress = new(Address)
	myVCard.adress.city = "bj"
	myVCard.adress.country = "CN"
	myVCard.adress.street = "YCBYJ"
	myVCard.adress.code = "102201"
	myVCard.age = 25
	myVCard.corpation = "COMAC"
	myVCard.name = "Jevade"
	myVCard.telNum = "13112345678"
	print(&myVCard)
	p1 := new(point.Point2D)
	p1.Set(2, 3)
	p2 := &point.Point2D{3, 5}
	fmt.Println(p1.ABS2D())
	p1.Scale(2)
	fmt.Println(p1.ABS2D())
	//数学 网络
	fmt.Println(p1.Dis(p2))
	p3 := point.Point3D{1, point.Point2D{2, 3}}
	fmt.Println((&p3).ABS2D())
	http.HandleFunc("/", root)
	http.HandleFunc("/long", long)
	// http.ListenAndServe("localhost:8080", nil)
	reflectType(*(myVCard.adress), 1)
	employee := employee.NewEmployee(1000, "li", 1980, 1)
	employee.GiveRaise(1.2, 1)
	fmt.Println(employee.Age(), employee.Salary())
	acar := car.NewMecideris(4, "MD")
	acar.Start()
	fmt.Println(acar.NumberOfWheel())
	car.TestCar()
	vendo := vendo.Vendo{}
	vendo.Magic()
	vendo.MoreMagic()
	integer.Test()
	integer.Test2()
	readFile("insert.go")

	//读入csv文件
	// ps := file.ReadProduct("product.txt")
	// file.PrintProducts(ps)
	// file.UnCompress("1.zip")
	// file.WriteFile("1.txt")
	// file.CopyFile("2.txt", "1.txt")
	//TestMyFlag()
	// file.MyCat()

	fmt.Println(myVCard.adress)
	adress := Address{street: "1", city: "2", code: "3", country: "4", Num: 123}
	js, err := json.Marshal(adress)
	fmt.Println("err:", err, adress)
	os.Stdout.Write(js)
	fmt.Printf("Json format:%s\n", js)

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	// SaveJson(b, "Vcard.json")
	// var mygropu interface{}
	ReadJson("Vcard.json")
	var mygropu ColorGroup
	json.Unmarshal(b, &mygropu)
	fmt.Println(mygropu)
    Check()
	protect(TestPaninc)
	fmt.Println("END")
}

func Check(){
	var user = os.Getenv("USER")
	if user == ""{
		panic("Unknown user: no value for $USER")
	}else{
		fmt.Println(user,"is runing")
	}
}

func protect(g func()){
   defer func(){
	   log.Println("done")
	   if err := recover();err!=nil{
		   log.Printf("run time panic:%v", err)
       }
   }()
	log.Println("start")
	g()
}

func TestPaninc(){
	panic(12222)
}

//TestMyFlag is to func as  cat
func TestMyFlag() {
	var NewLine = flag.Bool("n", false, "print newline")
	var Newline = "\n"
	// flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for ix := 0; ix < flag.NArg(); ix++ {
		if ix > 0 {
			s += ""
			if *NewLine {
				s += Newline
			}
		}
		s += flag.Arg(ix)
	}
	os.Stdout.WriteString(s)
}
func readFile(name string) {
	page := &file.Page{}
	page.Read(name)
	page.PrintTitle()
	page.PrintContent(100)
}

func reflectType(tt Address, idx int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(idx)
	fmt.Printf("%v\n", ixField.Tag)
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

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	rootHtmlTmpl.Execute(w, nil)
}
func long(w http.ResponseWriter, r *http.Request) {
	shortURL := r.FormValue("shortUrl")

	longURL := "long__" + shortURL
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Longer version of %s is : %s", shortURL, longURL))
}
