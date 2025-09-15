package main

import (
	"fmt"
	"sync"
)

func main() {
	serverAddr := "localhost:1234"
	server := WordCountServer{addr: serverAddr}
	server.Listen()

	input1 := "hello I am good hello bye bye bye bye good night hello"
	input2 := "bye bye hello good night"
	ch1 := makeRequest(input1, serverAddr)
	ch2 := makeRequest(input2, serverAddr)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		result1 := <-ch1
		checkError(result1.Err)
		fmt.Println("Result 1:", result1.Counts)
	}()

	go func() {
		defer wg.Done()
		result2 := <-ch2
		checkError(result2.Err)
		fmt.Println("Result 2:", result2.Counts)
	}()

	fmt.Println("Requests sent, waiting for results...")

	wg.Wait()
	fmt.Println("All results received.")
}
