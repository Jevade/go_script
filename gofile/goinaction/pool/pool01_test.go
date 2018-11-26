package pool_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"../pool"
)

const checkmark = "\u2713"
const ballltx = "\u2717"

func TestNew(t *testing.T) {
	url := "http://163.com"
	statusCode := 200
	t.Log("Test if can get url by http ")
	{
		code, err := pool.Get(url)
		if err != nil {
			t.Fatal("Get return error")
		}
		t.Logf("\t\t Should make a successful call. %v", checkmark)
		if code == statusCode {
			t.Logf("\t\t Should receive a \"%d\" status.%v", statusCode, checkmark)
		} else {
			t.Errorf("\t\t Should receive a \"%d\",but receive \"%d\"  status.%v", statusCode, code, ballltx)
		}

	}
}

func ExampleGet() {
	url := "http://163.com"
	code, err := pool.Get(url)
	if err != nil {
		log.Println("Err occue when get", err)
	}
	fmt.Println(code)
	//Output:
	//200
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}

}

func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}

}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}

}
