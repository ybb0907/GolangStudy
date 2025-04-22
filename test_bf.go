package main

import (
	"fmt"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	fmt.Printf("...跑起来了")
	wg.Done()
}

func main() {
	c1 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c1)
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go Run(&wg)
	// wg.Wait()
}
