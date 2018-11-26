package pool

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

//Pool is type define a resoure a pool
type Pool struct {
	m         sync.Mutex                "mutex"
	resources chan io.Closer            "resouce chan"
	factory   func() (io.Closer, error) "factory function return resource"
	closed    bool                      "closed flag"
}

var (
	ErrNewPoolFailed = errors.New("Pool Size error")
	ErrPoolCloseed   = errors.New("Pool closed")
	ErrNoResource    = errors.New("Pool no resource")
)

func Get(url string) (int, error) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("test.html")
	if err != nil {
		log.Fatalln("Create file failed", err)
	}
	defer file.Close()
	io.Copy(file, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
	return r.StatusCode, err
}

//New is factory and return a resource pool
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 1 {
		log.Println("Pool Size error")
		return nil, ErrNewPoolFailed
	}
	pool := Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
		closed:    false,
	}
	return &pool, nil
}

//Acquire a resouce to handler
func (p *Pool) Acquire() (*io.Closer, error) {
	//Thread safety
	p.m.Lock()
	defer p.m.Unlock()

	//resource pool closed,return error
	if p.closed {
		return nil, ErrPoolCloseed
	}

	select {
	case closer, ok := <-p.resources: //pool
		log.Println("Acquired:", "Shared Resource")
		if !ok {
			return nil, ErrPoolCloseed
		}
		return &closer, nil
	default: //create new resource
		log.Println("Acquired:", "New Resource")
		closer, err := p.factory()
		return &closer, err
	}
}

//Release will release acquired resource
func (p *Pool) Release(r io.Closer) {
	//thread safety
	p.m.Lock()
	defer p.m.Unlock()

	//pool closed,destorey resource
	if p.closed {
		r.Close()
		return
	}
	// release resouce
	select {
	//return resource to chan
	case p.resources <- r:
		log.Println("Release:", "IN Queue")
		return
	// resource pool is full,destory resource
	default:
		log.Println("Release:", "Closeing")
		r.Close()
		return
	}
}

//Close will close the resource pool
func (p *Pool) Close() {
	//sure the thread safety
	p.m.Lock()
	defer p.m.Unlock()

	//if closed,do nothing
	if p.closed {
		return
	}
	//close resource pool
	p.closed = true

	//close the resouce chan,should done first to avoid dead lock
	close(p.resources)

	//close all resource in chan
	for r := range p.resources {
		r.Close()
	}
}
