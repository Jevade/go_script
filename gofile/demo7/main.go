package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"

	"../bitmap"
)

//LENGTH is const value
const LENGTH uint = 3000000

//Mystring is new type
type Mystring string

func main() {
	var t = "123"
	xstr := Mystring(t)
	fmt.Println(xstr)
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS)
	fmt.Println(strconv.IntSize)
	log.Println("Start")
	theBitmap := bitmap.NewBitmap(LENGTH)
	for i := 0; i < int(LENGTH); i += 3 {

		theBitmap.Set(uint(i))
	}
	for i := 0; i < int(LENGTH); i++ {
		_ = theBitmap.IsSet(uint(i))
		// if i%50000000 == 0 {
		// fmt.Println(flag)
		// }
	}
	log.Println("End")
	name := getTheFlag()
	flag.Parse()
	fmt.Println("Hello, ", *name)
	var err error
	n, err := io.WriteString(os.Stdout, "Hello everyone!")
	if err != nil {
		fmt.Println("error", err, n)
	}

	fmt.Println("Success", n)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}
