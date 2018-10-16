package myserver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type work struct {
	method string
	param  []string
}

func server() chan work {
	works := make(chan work)
	go func() {
		for job := range works {
			go safelyDo(&job)
		}
	}()
	return works
}
func safelyDo(job *work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("work field with %s in %v", err, job)
		}

	}()
	do(job)
}

func do(work *work) {
	fmt.Println(work.method)
	fmt.Println(work.param)
}

// Launch is launch
func Launch() {
	works := server()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("send data")
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			fmt.Printf("err:%s,Please try again", err)
			continue
		}
		switch line {
		case "p":
			works <- work{"post", []string{"1", "2", "3", line}}
		case "f":
			works <- work{"fetch", []string{"1", "2", "3", line}}
		case "h":
			works <- work{"head", []string{"1", "2", "3", line}}
			return
		default:
			works <- work{"get", []string{"1", "2", "3", line}}

		}
	}
}

//MyServer func as Server
func MyServer() {
	Launch()
	// frontend()
}
func backend() (chan string, chan string, chan string, *time.Timer) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	chStop := make(chan string)
	ratePerSec := 1
	var dur = time.Duration(1e9 / ratePerSec)
	ticker := time.Tick(dur)
	boom := time.NewTimer(3 * dur)

	go func() {
		for {
			fmt.Println("get data from chan")
			select {
			case cmd := <-ch1:
				fmt.Printf("Do job %s from ch1\n", cmd)
			case cmd := <-ch2:
				fmt.Printf("Do job %s from ch2\n", cmd)
			case cmd := <-chStop:
				fmt.Printf("Do job %s from chStop\n", cmd)
				return
			case c := <-ticker:
				log.Println(c)
			}
		}
	}()
	return ch1, ch2, chStop, boom
}

func frontend() {
	ch1, ch2, chStop, boom := backend()
	defer func() {
		close(ch1)
		fmt.Println("close ch1")
	}()
	defer func() {
		close(ch2)
		fmt.Println("close ch2")
	}()
	defer func() {
		close(chStop)
		fmt.Println("close chStop")
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("send data")
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			fmt.Printf("err:%s,Please try again", err)
			continue
		}
		switch line {
		case "1":
			ch1 <- line
			fmt.Println("TO ch1")
		case "2":
			ch2 <- line
			fmt.Println("TO ch2")
		case "q":
			chStop <- line
			fmt.Println("TO chStop")
			return
		default:
			select {
			case <-boom.C:
				fmt.Println("**********time is comming")
				return
			case ch1 <- line:
				fmt.Println("TO ch1")
			case ch2 <- line:
				fmt.Println("TO ch2")
			}
		}
	}
}
