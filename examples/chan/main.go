package main

import (
	"math/rand"
	"time"

	"github.com/albekov/progress"
)

func main() {
	ch := createChan()
	progress := progress.ProgressChan(ch, progress.WithTotal(100))
	for value := range progress {
		doSomething(value)
	}
}

func createChan() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			t := rand.Intn(100) + 5
			time.Sleep(time.Duration(t) * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func doSomething(value int) {
	//
}
