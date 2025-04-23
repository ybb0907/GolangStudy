package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClass() {
	o := sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Printf("once")
		})
	}
}

func test_waitgroup() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Printf("1")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Printf("2")
		wg.Done()
	}()

	wg.Wait()
}

func test_syncMap() {
	mp := &sync.Map{}

	go func() {
		for i := 0; i < 10; i++ {
			mp.Store(i, i)
		}
	}()

	time.Sleep(time.Second * 1)
	go func() {
		mp.Range(func(key, value interface{}) bool {
			fmt.Printf("key: %d, value: %d \n", key, value)
			return true
		})
	}()

}

func main() {
	test_syncMap()
	time.Sleep(time.Second * 100)
}
