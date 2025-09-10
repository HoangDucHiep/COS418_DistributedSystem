package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start:", id)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Done :", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go work(i, &wg) // chạy song song
	}
	wg.Wait() // chờ tất cả goroutine xong
	fmt.Println("All done")
}
