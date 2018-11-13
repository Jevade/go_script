package main

import (
	"fmt"
	"log"
	"os"

	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Xijin")
	fmt.Println("12333")
	log.Println(123)
}
