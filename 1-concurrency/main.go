package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ch := make(chan int, 10)
	res := make([]int, 10)

	go func(in chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- rand.Int() % 100
		}
		close(in)
	}(ch)

	go func(out <-chan int) {
		i := 0
		for num := range out {
			res[i] = num * num
			i++
		}
	}(ch)

	fmt.Print(res)
}
