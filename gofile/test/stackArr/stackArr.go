package stackArr

import (
	"fmt"
	"strconv"
)

//StackArr is a stack
type StackArr struct {
	value   []interface{}
	top     int
	thesize int
}

//NewStack will return a new stack
func NewStack(size int) *StackArr {
	return &StackArr{
		value:   make([]interface{}, size),
		thesize: size,
		top:     -1,
	}
}

//IsFull determine if stack full
func (T *StackArr) IsFull() bool {
	return bool(T.top == T.thesize)
}

//IsEmpty determine if stack empty
func (T *StackArr) IsEmpty() bool {
	return bool(T.Len() == 0)
}

//Push will push item into stack
func (T *StackArr) Push(v interface{}) {
	if ok := T.IsFull(); ok {
		fmt.Println("full")
		return
	}
	T.top++
	T.value[T.top] = v
}

//Pop will pop item out of stack
func (T *StackArr) Pop() (re interface{}) {
	if ok := T.IsEmpty(); ok {
		fmt.Println("Empty")
		return
	}
	re = T.value[T.top]
	T.top--
	return
}

func (T *StackArr) String() string {
	var str string
	for ix := 0; ix <= T.top; ix++ {
		switch T.value[ix].(type) {
		case int:
			str += strconv.Itoa(T.value[ix].(int))
		case string:
			str += T.value[ix].(string)
		default:
			str += fmt.Sprintf("%v", T.value[ix])
		}
		str += ","
	}
	return str
}

//Len return length og stack
func (T *StackArr) Len() int {
	return T.top + 1
}
