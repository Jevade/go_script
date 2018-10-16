package myparse

import (
	"fmt"
	"strconv"
	"strings"
)

//ParseError is self defined error type
type ParseError struct {
	Index int
	Word  string
	Err   error
}

//String will convert Parseerror to string
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse:error parsing %q as int", e.Word)
}

//Parse will convter stings to ints
func MyParse(input string) (number []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg:%v", r)
			}
		}
	}()
	fields := strings.Fields(input)
	number, err = Fields2number(fields)
	return
}

func Fields2number(strslice []string) (numbers []int, err error) {
	if len(strslice) < 1 {
		panic("empty strings")
	}
	for idx, value := range strslice {
		number, err := strconv.Atoi(strslice[idx])
		if err != nil {
			panic(&ParseError{idx, value, err})
		}
		numbers = append(numbers, number)
	}
	return
}
