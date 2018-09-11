package stackArr

import ( 
        "strconv"
	"fmt"
)
type StackArr struct{
    value [4]int
    top int
}

func(T*StackArr)IsFull()bool{
    return bool(T.top==len(T.value))
}

func(T*StackArr)IsEmpty()bool{
    return bool(T.top==0)
}

func(T*StackArr)Push(v int){
    if ok:= T.IsFull();ok{
        fmt.Println("full")
    }
    T.value[T.top] = v
    T.top += 1
}

func(T *StackArr)Pop()(re int){
    if ok:=T.IsEmpty();ok{
        fmt.Println("Empty")
	return 
    }
    T.top -=1
    re = T.value[T.top]
    return
}

func(T *StackArr)String()string{
    var str string
    for ix :=0;ix<T.top;ix++{
        str+=strconv.Itoa(T.value[ix])
        str += ","
    }
    return str
}
