package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../parse"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fields := RemoveDuplicatesAndEmpty(strings.Split(req.URL.Path, "/"))
	fmt.Println(len(fields))

	var resp string
	if len(fields) == 2 {
		switch fields[0] {
		case "hello":
			resp = fields[1]
		case "shouthello":
			resp = strings.ToUpper(fields[1])
		default:
			resp = fields[0] + "/" + fields[1]
		}

	} else {
		resp = string(req.URL.Path[1:])
	}
	fmt.Println(fields[0])
	fmt.Fprintf(w, "<h1>%s<h1><div>Hello,%s</div>", "Hello", resp)
}

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="num" />
			<input type="submit" value="submit"/>
		</form>
		<form action="#"  name="result">
		<input type="text" name="mean" value="%f" />
		<input type="text" name="max" value="%f"/>
		<input type="text" name="min" value="%f"/>
	</form>
	</body></html>
`

func FormServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "xml")
	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, "Hi:"+req.FormValue("in"))
	}
}
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

type Hello struct {
}

func (t Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>%s<h1><div>Hello,%s</div>", "Hello", req.URL.Path)
}

func TestChan() chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("1233")
		ch <- 1
		close(ch)
	}()
	return ch

}

func processPost(r *http.Request) (nums []int, err error) {
	numStr := strings.Split(r.FormValue("num"), ",")
	if err != nil {
		return
	}
	nums, err = myparse.Fields2number(numStr)
	if err != nil {
		return
	}
	return nums, nil
}
func calc(nums []int) (a, b, c float64) {
	if len(nums) < 1 {
		return
	}
	a, b = float64(nums[0]), float64(nums[0])
	for _, v := range nums {
		if a > float64(v) {
			a = float64(v)
		}
		if b < float64(v) {
			b = float64(v)
		}
		c += float64(v)
	}
	c /= float64(len(nums))
	return
}

func HomeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Begin process")
	switch r.Method {
	case "GET":
		fmt.Println("Begin process get")
		w.Header().Set("Content-Type", "xml")
		fmt.Fprintf(w, form, 0.0, 0.0, 0.0)
	case "POST":
		fmt.Println("Begin process post")
		w.Header().Set("Content-Type", "json")
		result, err := processPost(r)
		fmt.Println(result)
		if err != nil {
			fmt.Println("ERR", err)
			panic(err)
		}
		a, b, c := calc(result)
		fmt.Fprintf(w, form, a, b, c)
	}
}

func Server() {

	// http.Handler
	http.HandleFunc("/", HomeServer)
	// http.HandleFunc("/he/", HelloServer)
	// http.HandleFunc("/she/", FormServer)
	//  http.Handle("/", http.HandlerFunc(HelloServer))
	//  err := http.ListenAndServe("localhost:8080", Hello{})
	err := http.ListenAndServe("localhost:8080", nil)
	// err := http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListerAndServer:", err.Error())
	}
	ch := TestChan()
	for v := range ch {
		fmt.Println(v)
	}
}

func TalkToServer(url string) {
	res, err := http.Get(url)
	checkErr(err)
	fmt.Println(res.Cookies, res.Header)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(data[:200]))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Get : %v", err)
	}
}
