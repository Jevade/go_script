package integer

import (
	"fmt"
)

type Integer int

func (p Integer)get()int{
	return int(p)
}

func f(i int){
	fmt.Println(i+2)
}

func Test(){
	var v Integer
	f(int(v))
}

type Integer2 struct{
	n int
}

func (p Integer2) get()int{
	return p.n
}

func Test2(){
	var v Integer2
	f(v.n)
}