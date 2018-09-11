package typeString

import "strconv"
type T struct{
    A int 
    B float64
    C string
}

func (t *T)String()string{
    return  strconv.Itoa(t.A)+"/"+strconv.FormatFloat(t.B,'e',7,32)+"/"+t.C
}

