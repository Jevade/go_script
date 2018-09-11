package main

//https://github.com/Jevade/the-way-to-go_ZH_CN/blob/master/eBook/07.3.m
import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

var re = regexp.MustCompile("[0-9]+")
var fibs [50]int

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Print(file, ":", line)
	}
	a := insertSort(1, 2)
	fmt.Println(a)
	where()
	result, ok := mysqrt(1121)
	if ok {
		fmt.Println(result)
	}
	where()
	fmt.Println(IsNumPosi(17.0))
	season, err := Season(3)
	if err {
		fmt.Println("not in range")
		return
	}
	where()
	fmt.Println(season)
	itera(7)
	iteraStr("This is a very import question!")
	gotointera(9)
	printIttera(10)
	printGoIttera(10)
	fmt.Println(MultiPly3Num(1, 2, 3))
	fmt.Println(add(1, 2))
	fmt.Println(add_3(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(sub_3(1, 2))
	fmt.Println(multi(1, 2))
	fmt.Println(multi_3(1, 2))
	fmt.Println(mysqrts_3(123))
	min, max := MinMax(1, 0)
	fmt.Printf("Min is :%d,Max is %d\n", min, max)
	n := 0
	reply := &n
	Multiply(3, 3, reply)
	fmt.Println(n)

	fmt.Println(MinList(3, 1, 6, 4, -6, 9))

	F1("a", "b", "c")
	typecheck("123", 123, true)
	A()
	doDB()
	bbb := new(string) //new 返回构造的变量的地址，指向分配好的一块空间
	fmt.Println(len(*bbb))
	fmt.Println(getFib_n(5))
	myprint(10)
	printStepList(9)
	fmt.Println(replaceNoASSIC("HELLOJLJ$#@HJK刘家伟"))
	fun := func(x, y int) int { return x * y }
	fmt.Println(fun(1, 2))
	fmt.Println(f())
	fmt.Println("闭包")
	num := 10

	now := time.Now()
	for i := 0; i < num; i++ {
		fmt.Println(fibonacci(i))
	}
	end := time.Now()
	delta := end.Sub(now)
	fmt.Printf("recurisive fib took this amount of time: %s\n", delta)
	now = time.Now()
	for i := 0; i < num; i++ {
		fmt.Println(fibonacci_cache(i))
	}
	end = time.Now()
	delta = end.Sub(now)
	fmt.Printf("cache fib took this amount of time: %s\n", delta)

	a_c := fibonacciC()
	now = time.Now()
	for i := 0; i < num; i++ {
		fmt.Println(a_c())
	}
	end = time.Now()
	delta = end.Sub(now)
	fmt.Printf("close package fid took this amount of time: %s\n", delta)

	where()
	array()

	as := new([3]int)
	ff(*as)
	ffp(as)
	var lists [5]int = [...]int{1, 2, 3, 4, 5}
	list1 := [5]int{3: 1233, 4: 234}
	list2 := []int{2, 3, 4}[:2]
	ss := [3]int{1, 2, 3}
	s := ss[:]
	list4 := [6]int{2, 3, 4, 5, 6, 7}
	fmt.Println(lists, list1, len(list2), cap(list2), list4, list4[:2])
	fmt.Println(sumArray(list4[:]))
	fmt.Println(s, list4)
	pfList(list4[:], "string")
	list5 := []string{"old", "old", "old"}[:]
	pfListIx(list5)
	arrayOp()
	list6 := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
	fmt.Println(Sum(list6[:3]))

	fmt.Println("the min is ", minSlice(list4[:]))
	fmt.Println("the max is ", maxSlice(list4[:]))
	fmt.Println(SumAndAverage(2, 2.4))
	diffNewMake()
	fmt.Println(numfib(35))
	z := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(len(z), cap(z))
	s2 := z[2:]
	s2[1] = 't'
	qz := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(z)
	fmt.Println(len(z), cap(z))
	z = apps(z, qz)
	fmt.Println(z)
	fmt.Println(len(z), cap(z))
	ka, ab := z[:5], z[5:]
	fmt.Println(ka, ab)
	appslice()
	testSliceCopyAppend()
	fmt.Println(appendbyte(qz, 'A', 'b', 'c'))
	fmt.Println(len(qz), cap(qz), len(resize(qz, 2)), cap(qz))
	fmt.Println(filterOdd([]int{1, 2, 3, 4, 5, 6}, IsOdd))
	fmt.Println(InsertStringSlice(qz, 'l', -9))
	fmt.Println(RemoveStringSlice(qz, 0, 5))
	s2b := func(s string) []byte {
		return []byte(s)
	}
	nnn, _ := fmt.Println(s2b("it is "))
	fmt.Println(nnn)
	fmt.Println(chstring("hello", 2, 'g'))
	fmt.Println(StrComp(qz, appendbyte(qz, 'a')))
	numSlice := []int{2, 5, 1, 8, 5, 3, 2}
	fmt.Println(qz)
	sort.Ints(numSlice)
	fmt.Println(sort.IntsAreSorted(numSlice))
	fmt.Println(sort.SearchInts(numSlice, 5), numSlice)
	fmt.Println(sl713("it is a very useful question"))
	fmt.Println(copyRepu("itt iss a very useful question"))
<<<<<<< HEAD
	fmt.Println(pop([]int{2,5,4,1,9,5,3,7,2}))
	fun1 := func(a int)int{return a * 10}
	fmt.Println(mapfunc(fun1,[]int{1,2,3,4,5}))
	fmt.Println(123)
=======
	fmt.Println(pop([]int{2, 5, 4, 1, 9, 5, 3, 7, 2}))
	fun1 := func(a int) int { return a * 10 }
	fmt.Println(mapfunc(fun1, []int{1, 2, 3, 4, 5}))
	MapTest()

}
func MapTest() {
	newMap := make(map[string]uint32, 100)
	newMap2 := newMap
	newMap["it"] = 2
	fmt.Println(len(newMap), newMap)
	newMap2["is"] = 3
	fmt.Println(newMap)
	for i := 0; i < 200; i++ {
		newMap[string(i)] = uint32(i)
	}

	fmt.Println(MapSearch("333", newMap))

}
func MapSearch(key string, themap map[string]uint32) uint32 {
	value, isStored := themap[key]
	if !isStored {
		panic("not exist")
	}
	return value
>>>>>>> ad95354943097fdb34996bbe89dbe6434a2603bf
}

// func map1 map[key]
func mapfunc(fun func(int) int, slice []int) []int {
	for ix := range slice {
		slice[ix] = fun(slice[ix])
	}
	return slice
}
func pop(sli []int) []int {
	var flag int
	for ix := 0; ix < len(sli); ix++ {
		for iy := 0; iy < len(sli)-ix-1; iy++ {
			if sli[iy] > sli[iy+1] {
				flag = sli[iy+1]
				sli[iy+1] = sli[iy]
				sli[iy] = flag
			}
		}
	}
	return sli
}
func copyRepu(str string) []byte {
	result := make([]byte, 0)
	for ix := 1; ix < len(str); ix++ {
		if str[ix] == str[ix-1] {
			result = append(result, str[ix])
		}
	}
	return result
}
func sl713(str string) string { //reverse a string
	if len(str) < 2 {
		return str
	}
	return sl713(str[len(str)/2:]) + sl713(str[:len(str)/2])
}
func findOneNum(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("open error")
		return nil
	}
	b = re.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
	//if a slice is not referenced,it will be recollected,
	//Just return a slice of a big array will take up large
	// memory, so we can make another slice and cope values to
	//this new slice and return it to save memory.
}
func splitSlice(sli string, pos int) (string, string, string) {
	return sli[:pos], "|", sli[pos+1:]
}
func findNum(filename string) []byte {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("open error")
		return nil
	}
	b := re.FindAll(f, len(f))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}
func StrComp(a, b []byte) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return true
			fallthrough
		case a[i] < b[i]:
			return false
		}
	}
	return bool(len(a) > len(b))
}
func chstring(str string, pos int, c byte) string {
	s := []byte(str)
	s[pos] = c
	return string(s)
}
func RemoveStringSlice(slice []byte, start, end int) []byte {
	if start < 0 {
		start = 0
	}
	if end > len(slice)-1 {
		end = len(slice) - 1
	}

	newslice := slice[end+1 : len(slice)]
	slice = slice[:start+len(newslice)]
	copy(slice[start:], newslice)
	return slice
}
func InsertStringSlice(slice []byte, c byte, pos int) []byte {
	fmt.Println("origin slice is", slice)
	if pos > len(slice)-1 {
		pos = len(slice)
	}
	if pos < 0 {
		pos = 0
	}
	newslice := make([]byte, len(slice)+1)
	if pos != 0 {
		copy(newslice[:pos], slice[:pos])
	}
	newslice[pos] = c

	if pos < len(slice) {
		copy(newslice[pos+1:], slice[pos:])
	}
	return newslice
}
func filterOdd(slice []int, fun func(int) bool) []int {
	newslice := make([]int, 0, 10)
	for ix := range slice {
		if fun(slice[ix]) {
			continue
		} else {
			newslice = newslice[0 : len(newslice)+1]
			newslice[len(newslice)-1] = slice[ix]
			// newslice = append(newslice,slice[ix])
		}
	}
	return newslice
}

func IsOdd(num int) bool {
	if num%2 == 1 {
		return true
	}
	return false
}
func resize(slice []byte, factor int) []byte {
	m := len(slice) * factor
	if m > cap(slice) {
		newslice := make([]byte, 2*m+2)
		copy(newslice, slice)
		slice = newslice
	}
	slice = slice[0:m]
	return slice
}
func testSliceCopyAppend() {
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	fmt.Println(1234)
	sl3 = append(sl3, []int{4, 5, 6}...)
	fmt.Println(sl3)
	fmt.Println(1234)

}
func appendbyte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := len(data) + m
	if n > cap(slice) {
		newslice := make([]byte, (n+1)*2)
		copy(newslice, slice)
		slice = newslice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func appslice() {
	slice := make([]int, 0, 10)
	for ix := 0; ix < cap(slice); ix++ {
		slice = slice[0 : ix+1]
		slice[ix] = ix
		fmt.Println("the length of slice is ", len(slice))
	}
	fmt.Println(slice)
}
func apps(slice, data []byte) []byte {
	if cap(slice) > len(slice)+len(data) {
		for ix := 0; ix < len(data); ix++ {
			slice[len(slice)] = data[ix]
		}
	} else {
		var buffer bytes.Buffer
		for ix := 0; ix < len(slice); ix++ {
			buffer.WriteByte(slice[ix])
		}
		for ix := 0; ix < len(data); ix++ {
			buffer.WriteByte(data[ix])
		}
		slice = buffer.Bytes()
	}
	return slice
}
func numfib(num int) []int {
	re := make([]int, num)
	for ix := 0; ix < num; ix++ {
		re[ix] = fibonacci_cache(ix)
	}
	return re
}
func diffNewMake() {
	var makeslice []int = make([]int, 5, 10)
	for ix := range makeslice {
		makeslice[ix] = ix * 5
	}
	for ix := 0; ix < len(makeslice); ix++ {
		fmt.Println("the", ix, "num is", makeslice[ix])
	}
	fmt.Println("len: ", len(makeslice), "cap:", cap(makeslice))
	// var newslice []int = new([]int,5)
	fmt.Println(makeslice)
}
func minSlice(slice []int) (min int) {
	min = slice[0]
	for i := 1; i < len(slice); i++ {
		if min > slice[i] {
			min = slice[i]
		}
	}
	return
}
func maxSlice(slice []int) (max int) {
	max = slice[0]
	for i := 1; i < len(slice); i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}
	return
}
func SumAndAverage(a int, b float32) (sum, ave float32) {
	sum = float32(a) + b
	ave = (sum / 2)
	return
}
func Sum(arr []float32) (re float32) {
	for ix := range arr {
		re += arr[ix]
	}
	return
}

func arrayOp() {
	items := [...]int{10, 20, 30, 40, 50}
	for ix := range items {
		items[ix] *= 2
	}
	fmt.Println(items)
}

func pfListIx(list []string) {
	for ix := range list {
		fmt.Println("the ", ix, "th num", list[ix])
		list[ix] = "New line"
	}
	fmt.Println(list)
}

func pfList(list ...interface{}) {
	for ix, value := range list {
		fmt.Println("the ", ix, "th num is", value)
	}
}

func sumArray(arr []int) (s int) {
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return
}

func ff(a [3]int)   { fmt.Println(a) }
func ffp(a *[3]int) { fmt.Println(*(a)) }

func array() {
	var list [20]int
	for j := 0; j < len(list); j++ {
		fmt.Println(j, ":", list[j])
	}
}

func f() (ret int) {
	defer func() {
		ret++
		fmt.Println("ret", ret)
	}()
	return
}

//close len cap new make copy append panic recover print println complex real imag

func print() {
	fmt.Println(123)
}

func insertSort(array float64, lenth float64) (a float64) {
	return array
}

func mysqrt(num float64) (result float64, ok bool) {
	if num < 0 {
		return
	}
	result, ok = math.Sqrt(num), true
	return
}

func IsNumPosi(num float64) (flag bool) {
	switch {
	case num > 0:
		flag = true
	default:
		flag = false
	}
	return
}
func Season(num int) (season string, err bool) {

	switch num {
	case 1:
		season = "1"
	case 2:
		season = "2"
	case 3:
		season = "3"
	case 4:
		season = "4"
	case 5:
		season = "5"
	case 6:
		season = "6"
	case 7:
		season = "7" // struct{
	case 8:
		season = "8" //         age int;
	case 9:
		season = "9" //         name string;
	case 10:
		season = "10" //        lenth float64;
	case 11:
		season = "11" //        id string;
	case 12:
		season = "12" // }Person
	default:
		err = true
	}
	return
}

func itera(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("this is %d th iteraations\n", i)
	}
}

func iteraStr(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("the %dth chacractor is %c\n", i, str[i])
	}
}

func gotointera(num int) {
	i := 0
LABLE1:
	fmt.Printf("this is %d th iteraations\n", i)
	i += 1
	if i == num {
		return
	}
	goto LABLE1

}

func printIttera(num int) {
	for i := 1; i < num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}

}
func printGoIttera(num int) {
	str := strings.Repeat("G", num)
	for i := 1; i < num; i++ {
		fmt.Println(str)
	}
	// for {
	//         fmt.Print(1)
	// }
}

func MultiPly3Num(a, b, c int) (multi int) { //multi 被初始化为0
	multi = a * b * c
	return
}

func add(a, b int) (re int) {
	re = a + b
	return
}

func add_3(a, b int) (re int) {
	return a + b
}

func sub(a, b int) (re int) {
	re = a - b
	return
}

func sub_3(a, b int) (re int) {
	return a - b
}

func multi(a, b int) (re int) {
	re = a * b
	return
}

func multi_3(a, b int) (re int) {
	return a * b
}

func mysqrts(a float64) (re float64, err error) {
	if a > 0 {
		re = math.Sqrt(a)
		return
	} else {
		err = errors.New("Not a positive num")
		return
	}
}

func mysqrts_3(a float64) (re float64, err error) {
	if a > 0 {
		re = math.Sqrt(a)
	} else {
		err = errors.New("Not a positive num")
	}
	return re, err
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func MinList(list ...int) (min int) {
	if (len(list)) == 0 {
		return
	}
	min = list[0]
	for i, s := range list {
		if s < min {
			min = s
			fmt.Println(i, "th is litter")
		}

	}
	return
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}
func F2(s ...string) {
	if len(s) == 0 {
		return
	}
	for _, v := range s {
		fmt.Println(v)
	}
}
func F3(s []string) {
	if len(s) == 0 {
		return
	}
	for _, v := range s {
		fmt.Println(v)
	}
}

func typecheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case float64:
			fmt.Println("float")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println(v)
		}
	}
}
func trace(s string) string {
	fmt.Println("entering", s)
	return s
}

func un(s string) {
	fmt.Println("leaving", s)
}

func A() {
	defer un(trace("a"))
	fmt.Println("in a")
	b()
}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
}

func connectToDB() {
	fmt.Println("connect")
}

func disconnectDB() {
	fmt.Println("disconnect")
}

func doDB() {
	connectToDB()
	defer disconnectDB()
	fmt.Println("do something")
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacci_cache(n int) (res int) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci_cache(n-1) + fibonacci_cache(n-2)
	}
	fibs[n] = res
	return
}

func getFib_n(n int) (int, int) {
	return n, fibonacci(n)

}
func myprint(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	myprint(n - 1)
}

func step(n int) (re int) {
	if n == 1 {
		re = 1
		return
	}
	if n == 0 {
		re = 1
		return
	}
	re = n * step(n-1)
	return
}

func printStepList(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i, step(i), ",")
	}
}
func IsAssic(c rune) bool {
	if int(c) > 255 {
		return false
	}
	return true
}

func replaceNoASSIC(s string) string {
	strs := ""
	for i, v := range s {
		if IsAssic(v) == false {
			fmt.Println(s[i])
			strs += "?"
		} else {
			strs += string(s[i])
		}
	}
	return strs
}

func fibonacciC() func() int {
	var result_0 int
	result := 1
	a := func() int {
		temp := result
		result += result_0
		result_0 = temp
		return result
	}
	return a
}
