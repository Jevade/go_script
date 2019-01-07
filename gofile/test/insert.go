package main

import (
	"flag"
	"fmt"
	"os"

    "errors"
	"./day"
	"./sort"
	"./stackArr"
	"./typeString"
	"./tz"
)

var name string

func init() {
	flag.StringVar(&name, "name", "every", "the greeting object.")
	// flag.String()
}
func main(){
 testInOperate()
}
func Maini() {
	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "Usage of %s/n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	n, ok := fmt.Printf("Hello,%s!\n", name)
	if ok == nil {
		fmt.Println("have println successful")
		fmt.Println("the length of name is: ", n)
	}
	fmt.Println(124)
	Slicetest()
	Ints()
}

//Ints will test Ints sort function
func Ints() {
	intlist := []int{123, 45, -34, 67, -89, 899, -3422}
	a := sort.IntArray(intlist)
	fmt.Printf("the origin list of int is %v\n", a)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("falied")
	}
	fmt.Printf("the sorted list of int is %v\n", a)

}

type operate func(x,y int) int

type calcFunc func(x,y int)(int,error)

func genCalc(op operate)calcFunc{
    return func(x,y int)(int,error){
        if op==nil{
            return 0, errors.New("invalid operation")
        }
        return op(x,y),nil
    }
}

func testInOperate(){
    f:= genCalc(func(x,y int)int{return x*y})
    v,err:= f(1,2)
    if err!=nil{
    fmt.Println("functions missing")
     return
    }
     fmt.Println(v)
}

type T typeString.T

func Slicetest() {
	var sli1 = []byte{'1', '2'}
	sli1 = append(sli1, '1', '2', '3')

	fmt.Println(sli1)
	sliss := []int{1, 2, 3}
	sliss = append(sliss, 2, 3)
	fmt.Println(sliss)
	t := &typeString.T{7, -2.35, "abc\tdef"}
	da := day.Day(int(6))
	fmt.Println(da)
	fmt.Printf("%v\n", t)
	tzz := tz.TZ(0)
	fmt.Println("tz", int(tzz), "is:", tzz)
	stack := new(stackArr.StackArr)
	stack.Push(1)
	fmt.Println(stack)
	stack.Push(2)
	fmt.Println(stack)
	fmt.Println(stack.Pop(), stack)
}
