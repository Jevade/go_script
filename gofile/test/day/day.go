package day

type Day int

var strings []string =[]string{"Mon","Tus","Wen","Thu","Fri","Sta","Sun"}

func (day Day)String()string{
    return strings[int(day)]
}
