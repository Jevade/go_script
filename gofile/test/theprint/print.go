package theprint

import (
	"os"
	"strconv"
)

// Strings is type of string
//
//
//

type Strings string

type Myfloat struct {
	value float64
}

type Day int

var days = []string{"Monday", "Thusday", "Wednesday", "Thirsday", "Friday", "satusday", "Sunday"}

func (day Day) String() string {
	return days[day]
}

func (mf *Myfloat) Set(vlue float64) {
	mf.value = vlue
}
func (mf *Myfloat) Get() float64 {
	return mf.value
}

func (mf *Myfloat) String() string {
	// return "2222"
	return strconv.FormatFloat(mf.value, 'E', -1, 32)
}

func (ms *Strings) String() string {
	return string(*ms)
}

// PrintValue ...PrintValu
//:w

//
//
func PrintValue(args ...interface{}) {
	for arg := range args {
		if arg > 0 {
			os.Stdout.WriteString(" ")
		}
		switch c := args[arg].(type) {
		case Strings:
			os.Stdout.WriteString(c.String())
		case Day:
			os.Stdout.WriteString(c.String())
		case Myfloat:
			os.Stdout.WriteString(c.String())
		}
	}

}
