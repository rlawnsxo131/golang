package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// Consumer struct
type Consumer struct {
	in   *chan int
	jobs chan int
}

// Work func
func (c Consumer) Work(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range c.jobs {
		time.Sleep(time.Second) // Do something
		fmt.Printf("%dth job finished\n", job)
	}
}

// Consume func
func (c Consumer) Consume(ctx context.Context) {
	for {
		select {
		case job := <-*c.in:
			c.jobs <- job
		case <-ctx.Done():
			close(c.jobs)
			return
		}
	}
}

// Producer struct
type Producer struct {
	in *chan int
}

// Produce func
func (p Producer) Produce() {
	task := 1
	for {
		*p.in <- task
		task++
	}
}

func main() {
	const nConsumers int = 10
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := make(chan int, 1)
	p := Producer{&in}
	c := Consumer{&in, make(chan int, nConsumers)}
	go p.Produce()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go c.Consume(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(nConsumers)
	for i := 0; i < nConsumers; i++ {
		go c.Work(wg)
	}
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancelFunc()
	wg.Wait()
}