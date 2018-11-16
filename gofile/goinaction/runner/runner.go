package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	task      chan func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error, 2),
		timeout:   time.After(d),
		task:      make(chan func(int)),
	}
}

func (r *Runner) Close() {
	close(r.task)
}
func (r *Runner) Add(task func(int)) {
	r.task <- task
}
func (r *Runner) run() (err error) {
	var id int
	go func() {
		for task := range r.task {

			task(id)
			id++
			r.complete <- err
		}
	}()
	return
}

// func (r *Runner) gotInterrupt() bool {
// 	select {

// 		return true
// 	default:
// 		return false
// 	}

// }
func (r *Runner) Start() {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.run()
	}()

}
func (r *Runner) Wait(lenth int) error {
	for i := 0; i < lenth; i++ {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		case <-r.timeout:
			return ErrTimeout
		default:
		}
		<-r.complete
	}
	return nil
}
