 package main

import (
            "fmt"
	    "./typeString"
	    "./day"
	    "./tz"
	    "./stackArr"
)
func main(){
	fmt.Println(124)
	Slicetest()
}
type T typeString.T
func Slicetest(){
	var sli1 = []byte {'1','2'}
	sli1 = append(sli1,'1','2','3')

	fmt.Println(sli1)
	sliss:= []int{1,2,3}
	sliss = append(sliss,2,3)
	fmt.Println(sliss)
        t := &typeString.T{7,-2.35,"abc\tdef"}
	da := day.Day(int(6))
	fmt.Println(da)
	fmt.Printf("%v\n",t)
	tzz := tz.TZ(0)
	fmt.Println("tz", int(tzz) ,"is:",tzz)
	stack := new(stackArr.StackArr)
	stack.Push(1)
	fmt.Println(stack)
	stack.Push(2)
	fmt.Println(stack)
	fmt.Println(stack.Pop(),stack)
}
