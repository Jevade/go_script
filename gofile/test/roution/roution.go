package roution

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"
)

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 1
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func test(ch1, ch2 chan int) {
	for i := 0; ; i++ {
		select {
		case x := <-ch1:
			fmt.Println("suck num from ch1,", i, x)
		case x := <-ch2:
			fmt.Println("suck num from ch2,", i, x)
		default:
			fmt.Println("not get")
		}
	}
}

//RunRoution is runroution
func RunRoution() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump2(ch2)
	go pump1(ch1)
	go test(ch1, ch2)
	testChan()
	time.Sleep(5 * 1e9)
}

func testChan() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	sendDataToChan(ch, "Washington")
	sendDataToChan(ch, "Tripoli")
	sendDataToChan(ch, "London")
	sendDataToChan(ch, "Beijing")
	sendDataToChan(ch, "Tokyo")
	close(ch)
}
func sendDataToChan(ch chan string, str string) {
	ch <- str
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Println("this str is: ", input)
	}
}
func Log(num float64) float64 {
	return math.Log(num)
}

func Mutil(num float64) float64 {
	return num * num
}

func Sqrt(num float64) float64 {
	return math.Sqrt(num)
}
func MutilCump(nums []float64, f func(float64) float64) []float64 {
	res := make([]float64, len(nums))
	ch := make(chan float64, len(nums))
	for i, xi := range nums {
		go func(i int, ix float64) {
			res[i] = f(ix)
			ch <- 0
		}(i, xi)
	}
	for i := 0; i < len(nums); i++ {
		<-ch
	}
	return res
}

type Empty interface{}
type semaphore chan Empty

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func Sum(x, y float64, ch chan float64) {
	ch <- x + y
}

func numGen(start, step, count int, out chan<- int) {
	for num, i := start, 0; i < count; i, num = i+1, num+step {
		out <- num
	}
	close(out)
}

// <-chan readonly chan
// chan<- sendonly chan
// chan send and read chan
func numCustom(in <-chan int, done chan<- bool) {
	for {
		num, ok := <-in
		if !ok {
			fmt.Println("not sucked successfully")
			break
		}
		fmt.Println(num)
	}
	// for num := range in {
	// 	fmt.Println(num)
	// }
	done <- true
}

func PC() {
	nums := make(chan int)
	done := make(chan bool)
	go numGen(10, 10, 9, nums)
	go numCustom(nums, done)
	<-done
}
func TestLoop1() {
	nums := make(chan int)
	done := make(chan bool)
	go addNum1(nums, done)
	for {
		select {
		case x := <-nums:
			fmt.Println("it is ", x)
		case <-done:
			os.Exit(1)
		}
	}
}

func TestLoop() {
	nums := make(chan int)
	go addNum(nums)
	for {
		x, ok := <-nums
		if !ok {
			fmt.Println("pipe closed ")
			break
		}
		fmt.Println("it is ", x)
	}
}

func addNum1(ch chan<- int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("begin add num")
		ch <- i //如果接受者没有准备好则，开始阻塞
	}
	done <- true
}

func addNum(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("begin add num")
		ch <- i //如果接受者没有准备好则，开始阻塞
	}
	close(ch)
}
