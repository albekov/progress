package main

import (
	"math/rand"
	"time"

	"github.com/albekov/progress"
)

func main() {
	progress := progress.New(progress.WithTotal(100))
	for i := 0; i < 100; i++ {
		t := rand.Intn(100) + 5
		time.Sleep(time.Duration(t) * time.Millisecond)
		progress.Inc()
	}
	progress.Done()
}
