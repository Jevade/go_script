package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"./pool"
	"./runner"
)

type user struct {
	name string
	age  int64
}
type admin struct {
	user
	level string
}
type notifer interface {
	notify()
}

func (u *user) notify() {
	fmt.Println(u.name, u.age)
}
func InterfaceNotify(n notifer) {
	n.notify()
}

func DisplayUser(u user) {
	fmt.Println(u.name, u.age)
}
func (u user) ChangeName1(name string) {
	u.name = name
}
func (u *user) ChangeName2(name string) {
	u.name = name
}
func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor :Task #%d.", id)
		fmt.Println()
		// time.Sleep(time.Duration(id) * time.Second)
	}
}

const (
	maxGoroutines  = 25
	pooledResource = 8
	length         = 5
)

type dbConnection struct {
	ID uint
}

func main() {
	// CurlTest()
	PoolTest()
	// RunnerTest()
}

var idCounter uint

func (conn dbConnection) Close() error {
	log.Printf("dbconection #%d closed", conn.ID)
	return nil
}
func createConnection() (io.Closer, error) {
	defer func() { idCounter++ }()
	return dbConnection{
		ID: idCounter,
	}, nil
}
func PoolTest() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResource)
	if err != nil {
		log.Println("Create pool failed")
		return
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQuerys(p, q)
			wg.Done()
		}(query)
	}
	wg.Wait()
	p.Close()

}
func performQuerys(p *pool.Pool, query int) {
	r, err := p.Acquire()
	if err != nil {
		log.Println("Acquire resource failed")
	}
	defer p.Release(*r)
	time.Sleep(time.Duration(rand.Intn(100000)) * time.Millisecond)
	log.Printf("QID[%d], CID[%d]\n", query, (*r).(dbConnection).ID)

}

func RunnerTest() {
	newRunner := runner.New(time.Second * 200)
	go func() {
		for id := 0; id < length; id++ {
			log.Printf("Create Task #%d.", id)
			newRunner.Add(createTask())
		}
		newRunner.Close()
	}()
	newRunner.Start()
	if err := newRunner.Wait(length); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminal due to  interrupt")
			os.Exit(1)
		case runner.ErrTimeout:
			log.Println("Terminal due to time out")
			os.Exit(2)

		}
	}
	log.Println("Process end")
	u := user{"lili", 20}
	InterfaceNotify(&u)
	var n notifer
	n = &u
	InterfaceNotify(n)
	DisplayUser(u)
	u.ChangeName1("jack") //传值传递的是副本，因此不会影响原本的值，需要完整拷贝副本。
	DisplayUser(u)
	(&u).ChangeName2("Jev") //传地址能直接影响原本的值，并且地址传递速度更快，节省资源
	DisplayUser(u)
	u.ChangeName2("Wy") //解引用，实际操作的是地址
	DisplayUser(u)
	// main is the entry point for the application.
	//值类型 string bool int/float
	//引用类型 map slice chan interface func 本身存在标头 header,传递的时候传递的就是指针，本身就是用于共享的，因此不需要特地传递引用类型的指针
	//使用的时候直接传递就可以。
	a := admin{
		user{"jjj", 20},
		"super",
	}
	InterfaceNotify(&a)
	var array1 [8]int
	array := array1[1:3]
	array[0] = 20
	fmt.Println(array1, array)
	// 限定了容量，没有超过容量的时候，和原来指针一样，超过容量会新开辟内存，不影响原来的数组，所以只会影响定义的时候设定的容量方位
	arry2 := array1[0:3:3]
	arry2[2] = 7
	arry2 = append(arry2, array...)
	fmt.Println(array, array1, arry2)
}
