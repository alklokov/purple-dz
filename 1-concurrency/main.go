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

func main() {
	nums := make([]int, 0, limit)
	for num := range proceed(generateNumbers()) {
		nums = append(nums, num)
	}
	for _, num := range nums {
		fmt.Print(num, " ")
	}
}

func generateNumbers() <-chan int {
	nums := make([]int, limit)
	for i := 0; i < limit; i++ {
		nums[i] = rand.Intn(maxNum + 1)
	}
	ch := make(chan int, limit)
	for _, num := range nums {
		ch <- num
	}
	close(ch)
	return ch
}

func proceed(ch <-chan int) <-chan int {
	ch2 := make(chan int, limit)
	for num := range ch {
		ch2 <- num * num
	}
	close(ch2)
	return ch2
}
