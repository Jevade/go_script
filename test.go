package main

import (
	"fmt"
	"math"
	"strings"
	"errors"
	"runtime"
	"log"
	"time"
)
var fibs [50]int
func main() {
	where := func(){
		_,file,line,_ := runtime.Caller(1)
		log.Print(file,":",line)
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
	fmt.Println(MultiPly3Num(1,2,3))
	fmt.Println(add(1,2))
	fmt.Println(add_3(1,2))
	fmt.Println(sub(1,2))
	fmt.Println(sub_3(1,2))
	fmt.Println(multi(1,2))
	fmt.Println(multi_3(1,2))
	fmt.Println(mysqrts_3(123))
	min , max := MinMax(1,0)
	fmt.Printf("Min is :%d,Max is %d\n", min ,max)
	n :=0
	reply := &n
	Multiply(3,3,reply)
	fmt.Println(n)

	fmt.Println(MinList(3,1,6,4,-6,9))

	F1("a","b","c")
	typecheck("123",123,true)
	A()
	doDB()
	bbb :=new(string)//new 返回构造的变量的地址，指向分配好的一块空间
	fmt.Println(len(*bbb))
	fmt.Println(getFib_n(5))
	myprint(10)
	printStepList(9)
	fmt.Println(replaceNoASSIC("HELLOJLJ$#@HJK刘家伟"))
	fun := func(x,y int)(int){return x*y}
	fmt.Println(fun(1,2))
	fmt.Println(f())
	fmt.Println("闭包")
	num:=10

	now := time.Now()
	for i:=0;i<num;i++{
    fmt.Println(fibonacci(i))
	}
	end:=time.Now()
	delta:=end.Sub(now)
	fmt.Printf("recurisive fib took this amount of time: %s\n", delta)
	now = time.Now()
	for i:=0;i<num;i++{
    fmt.Println(fibonacci_cache(i))
	}
	end =time.Now()
	delta =end.Sub(now)
	fmt.Printf("cache fib took this amount of time: %s\n", delta)

	a_c:=fibonacciC()
	now = time.Now()
	for i:=0;i<num;i++{
		fmt.Println(a_c())
	}
	end = time.Now()
	delta=end.Sub(now)
	fmt.Printf("close package fid took this amount of time: %s\n", delta)

	where()
	array()

	as:=new([3]int)
	ff(*as)
	ffp(as)
	var  lists [5]int = [...]int{1,2,3,4,5}
	list1 := [5]int{3:1233,4:234}
	list2 :=[]int{2,3,4}[:2]
	ss:=[3]int{1,2,3}
	s:=ss[:]
	list4 :=[6]int{2,3,4,5,6,7}
	fmt.Println(lists,list1,len(list2),cap(list2),list4,list4[:2])
	fmt.Println(sumArray(list4[:]))
	fmt.Println(s,list4)
	pfList(list4[:],"string")
	list5 :=  []string{"old","old","old"}[:]
	pfListIx(list5)
	arrayOp()
	list6 := []float32{1.1,1.2,1.3,1.4,1.5}
	fmt.Println(Sum(list6[:3]))

	fmt.Println("the min is ",minSlice(list4[:]))
	fmt.Println("the max is ",maxSlice(list4[:]))
	fmt.Println(SumAndAverage(2,2.4))
}
func minSlice(slice []int)(min int){
	min = slice[0]
	for i:=1;i<len(slice);i++{
		if min > slice[i]{
			min = slice[i]
	    }
	}
	return
}
func maxSlice(slice []int)(max int){
	max = slice[0]
	for i:=1;i<len(slice);i++{
		if max < slice[i]{
			max = slice[i]
		}
	}
	return
}
func SumAndAverage(a int,b float32)(sum, ave float32){
	sum = float32(a)+b
	ave = (sum/2)
	return
}	
func Sum(arr []float32)(re float32){
	for ix :=range arr{
	re+=arr[ix]
	}
	return
} 

func arrayOp(){
		items :=[...]int{10,20,30,40,50}
		for ix := range items{
			items[ix] *= 2
			}
		fmt.Println(items)
	}

func pfListIx(list []string){
	for ix := range list{
		fmt.Println("the ",ix,"th num",list[ix])
		list[ix]="New line"
	}
	fmt.Println(list)
}

func pfList(list ... interface{}){
	for ix ,value := range list{
		fmt.Println("the ",ix,"th num is",value)
	}
}

func sumArray(arr  []int)(s int){
		for i:=0;i<len(arr);i++{
			s+=arr[i]
		}
		return 
	}

func ff(a [3]int) { fmt.Println(a) }
func ffp(a *[3]int) { fmt.Println(*(a)) }

func array(){
	var list [20]int
	for j:=0;j<len(list);j++{
		fmt.Println(j,":",list[j])
	}
}

func f()(ret int){
	defer func(){
		ret++
		fmt.Println("ret",ret)
	}()
	return 1
}
//close len cap new make copy append panic recover print println complex real imag 

func print(){
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

func MultiPly3Num(a,b,c int)(multi int ){//multi 被初始化为0
multi = a * b * c
return
}

func add(a,b int)(re int){
re = a + b
return
}

func add_3(a,b int)(re int){
return a + b
}

func sub(a,b int)(re int){
re = a - b
return
}

func sub_3(a,b int)(re int){
return a - b
}

func multi(a,b int)(re int){
re = a * b
return
}

func multi_3(a,b int)(re int){
return a * b
}

func mysqrts(a float64)(re float64, err error){
if a>0 {
	re = math.Sqrt(a)
	return 
}else{
	err = errors.New("Not a positive num")
	return
}
}

func mysqrts_3(a float64)(re float64, err error){
if a>0 {
	re = math.Sqrt(a)
}else{
	err = errors.New("Not a positive num")
}
	return re ,err
}

func ThreeValues()(int, int, float32){
return 5, 6, 7.5
}

func MinMax(a int, b int)(min int, max int){
	if a < b {
		min = a
		max = b
	}else{
		min = b
		max = a
	}
	return
}

func Multiply(a ,b int ,reply *int){
	*reply = a * b
}

func MinList(list...int)(min int){
	if(len(list))==0{
		return 
	}
	min = list[0]
	for i,s := range list{
		if s < min{
			min = s
			fmt.Println(i,"th is litter")
		} 

	}
	return 
}

func F1(s ...string){
	F2(s...)
	F3(s)
}
func F2(s ...string){
	if len(s)==0{
		return
	}
	for _,v := range s{
		fmt.Println(v)
	}
}
func F3(s []string){
	if len(s)==0{
		return
	}
	for _,v := range s{
		fmt.Println(v)
	}
}

func typecheck(values ... interface{}){
	for _, value := range values{
		switch v:= value.(type){
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
func trace(s string) (string) {
	fmt.Println("entering", s)
	return s
}

func un(s string){
	fmt.Println("leaving",s)
}

func A(){
	defer un(trace("a"))
	fmt.Println("in a")
	b()
}
func b(){
	defer un(trace("b"))
	fmt.Println("in b")
}

func connectToDB(){
	fmt.Println("connect")
}

func disconnectDB(){
	fmt.Println("disconnect")
}

func doDB(){
	connectToDB()
	defer disconnectDB()
	fmt.Println("do something")
}

func fibonacci(n int )(res int){
	if n<=1{
		res =1
	}else{
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}


func fibonacci_cache(n int )(res int){
	if fibs[n] != 0 { 
		res= fibs[n]
		return
	}
	if n<=1{
		res = 1
	}else{
		res = fibonacci_cache(n-1) + fibonacci_cache(n-2)
	}
	fibs[n] = res
	return
}

func getFib_n(n int)(int, int){
	return n,fibonacci(n)
	
}
func myprint(n int){
	if n==0{
		return 
	}
	 fmt.Println(n)
	myprint(n-1)
}

func step(n int)(re int){
	if n==1{
		re = 1
		return
	}
	if n==0{
		re = 1
		return
	}
	re = n*step(n-1) 
	return
}

func printStepList(n int){
	for i:=n;i>0;i--{
		fmt.Println(i,step(i),",")
	}
}
func IsAssic(c rune)bool{
	if int(c)>255{
		return false
	}
	return true
}

func replaceNoASSIC(s string)string{
	strs := ""
	for i,v:=range s{
		if IsAssic(v)==false {
			fmt.Println(s[i])
		     strs+="?"
		}else{
			strs+=string(s[i])
		}
	} 
    return strs
}

func fibonacciC() func() int{
	var result_0 int
	result:=1
	a := func()int{
		temp:= result
		result += result_0
		result_0 = temp
		return result
	}
	return a
}
