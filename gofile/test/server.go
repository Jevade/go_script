package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"

	"./car"
	"./employee"
	"./file"
	"./integer"
	"./myexec"
	"./parse"
	"./point"
	"./roution"
	"./vendo"
	"./web"
)

//Go中需要被外部访问到的值均需要首字母大写，
//小写的话是私有值，无法被外部访问
// ,ok 1 作为函数返回值，衡量函数是否出错
// if value,err := func(v),err!=nil{
//   process(v)
// }

//2 检验映射里面里面是否有键值
// if v, ok = map[key]; ok {
//     process(v)
// }

//3 类型断言 检查接口类型是否包含了某类型
// if value,ok := v.(T);ok{
//     process(value)
// }

// 4 判断通道是否关闭
// if value,ok := <-ch;ok{
//     process(value)
//  }

// 修改字符串字符
// str := "Hello"
// c := []byte(str)
// c[0]='c'
// s2 := string(c)

//获取字符串子串
// subStr := str[m,n]

// 获取字符串的字符数
// len(str)

// utf8.RuneCountInString(str)

// len([]byte(str))

// 连接字符串 str1,str2
// 使用缓存的方式最快速
// var buffer bytes.Buffer
// buffer.WriteString(str1)
// buffer.WriteString(str2)
//也可以使用 +=
// str1 += str2

//解析命令行参数 os.Args flag.Parse

//map1 := make(map[keytype]valuetype)
//for key,value := range map1{
// }

//val1 , isPresent = map1[key1]
//delete(map1,key1)
//

// type struct1 struct{
// field1 type1
// field2 type2
// field3 type3
// }
// ms := new(struct1)
// ms := &struct1{1,2,3}
// ms := NewStruct(1,2,3)
//构建函数，作为唯一的获得类的方式，类名小写，对包外不可见，从而隐藏包的细节
// func NewStruct(a,b,c int)(* struct1){
// return &struct1(a,b,c)
// }

// 接口
// if v,ok := v.(Stringer),ok{
// 	fmt.Println("v 实现了接口Stringer")
// }

// func classifier(items ...interface{}){
// 	for i,x := range items{
// 		switch x.(type){
// 		case bool:
// 			fmt.Println("Bool")
// 		case int:
// 			fmt.Println("int")
// 		case float64:
// 			fmt.Println("float")
// 		default:
// 			fmt.Println(x.(type))

// 	}
// }
//函数 防止panic程序崩溃,使用内建函数recover终止panic
// func protect(g func()){
// 	defer func(){
// 		fmt.Println("done")
// 		if x:= recover();x!=nil{
// 			log.Panicln("run time panic")
// 		}
// 	}()
// 	log.Println("Start")
// 	g()
// }

//取消耗时很长的同步调用
// func quary(conns[]Conn,query string)Result{
// 	ch := make(chan Result ,1)//缓冲为保证至少第一个发送过来的数据可以存放，
// 	for _,conn := conns{
// 		go func(conn Conn){
// 			select{
// 			case ch <- conn.Do(query):
// 			default:
// 			}
// 			// result,err:= conn.Do(query)
// 			// if err!=nil{
// 			// 	fmt.Println("Connect error")
// 			// 	return
// 			// }
// 			// ch <- result
// 		}(conn)
// 	}
// 	return <-ch
// }

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

func startServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("It is err to listen tcp link")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error to Accept ")
			return
		}
		go solveConn(conn)

	}

}

func solveConn(conn net.Conn) {
	for {
		data := make([]byte, 512)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Read data field")
			return
		}
		fmt.Printf("Received data:%v\n", string(data[:n]))

	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Failed to Diel")
		return
	}
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("input you data")
		data, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(data, "\r\n")
		conn.Write([]byte(trimmedInput))
	}
}

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main2() {
	// 发起请求查询推特Goodland用户的状态
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	// 初始化XML返回值的结构
	user := User{xml.Name{"", "user"}, Status{""}}
	// 将XML解析为我们的结构
	data, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(data, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
func main3() {
	var values = [5]int{10, 11, 12, 13, 14}
	for ix := range values {
		fmt.Println(values[ix])
	}
	fmt.Println()
	time.Sleep(5e9)
	for ix := range values {
		go func() {
			fmt.Println(values[ix])
		}()

	}
	fmt.Println()
	time.Sleep(5e9)
	for ix := range values {
		go func(i int) {
			fmt.Println(i)
		}(ix)

	}
	fmt.Println()
	time.Sleep(5e9)
	for ix := range values {
		value := values[ix]
		go func() {
			fmt.Println(value)
		}()

	}
}
func main() {
	main3()
	web.Server()
	main2()
	var url string
	for {
		fmt.Scanln(&url)
		web.TalkToServer(url)
	}
	c := make(chan int)
	go func() {
		fmt.Println("ggogoogog")
		c <- 1
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	go startServer()
	client()
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
	TestParse()
	fmt.Println("END1")
	go myexec.MyStartProcessls()
	go myexec.MyStartProcessps()
	go myexec.MyStartProcessCMD()
	fmt.Println("END")
	time.Sleep(2 * 1e9)
	fmt.Println(roution.MutilCump([]float64{1, 2, 3, 4, 5}, roution.Log))
	fmt.Println(roution.MutilCump([]float64{1, 2, 3, 4, 5}, roution.Mutil))
	fmt.Println(roution.MutilCump([]float64{1, 2, 3, 4, 5}, roution.Sqrt))
	// roution.PC()
	// roution.RunRoution()
	roution.TestLoop()
}

//TestParse is TestParse
func TestParse() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}
	for idx := range examples {
		fmt.Printf("Parsing %q:\n", examples[idx])
		nums, err := myparse.MyParse(examples[idx])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}

func Check() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("Unknown user: no value for $USER")
	} else {
		fmt.Println(user, "is runing")
	}
}

func protect(g func()) {
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic:%v", err)
		}
	}()
	log.Println("start")
	g()
}

func TestPaninc() {
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
