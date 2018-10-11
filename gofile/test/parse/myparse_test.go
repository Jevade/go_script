package myparse

import (
	"fmt"
	"testing"
)

func TestMyParse(t *testing.T) {
	num, err := MyParse("1 2 3 4 5 2 3 4 5 6")
	if err != nil {
		t.Log("trans fails")
		t.Fail()
	}
	fmt.Println(num)
}

func BenchmarkMyParse(b *testing.B){
	num, err := MyParse("1 2 3 4 5 2 3 4 5 6")
	if err != nil {
		b.Log("trans fails")
		b.Fail()
	}
	fmt.Println(num)
}
