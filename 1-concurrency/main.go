package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type counter struct {
	value int
	mutex sync.Mutex
}

const (
	maxNum = 100
	limit  = 10
)

var (
	ch1 chan int
	ch2 chan int
)

func main() {
	ch1 = make(chan int, limit)
	ch2 = make(chan int, limit)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		generateNumbers(ch1)
	}()
	go func() {
		defer wg.Done()
		proceed(ch1, ch2)
	}()
	wg.Wait()
	for num := range ch2 {
		fmt.Print(num, " ")
	}
}

func generateNumbers(ch chan<- int) {
	nums := make([]int, limit)
	for i := 0; i < limit; i++ {
		nums[i] = rand.Intn(maxNum + 1)
	}
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

func proceed(ch <-chan int, resCh chan<- int) {
	for num := range ch {
		resCh <- num * num
	}
	close(resCh)
}
