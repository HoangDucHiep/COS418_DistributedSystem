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

	wg := sync.WaitGroup{}
	wg.Add(2)
	// Gửi request với callback
	call1 := makeRequest(input1, serverAddr)
	call2 := makeRequest(input2, serverAddr)

	go func() {
		defer wg.Done()
		call := <-call1.Done
		if call.Error != nil {
			checkError(call.Error)
			return
		}
		reply := call.Reply.(*WordCountReply)
		fmt.Println("Result 1:", reply.Counts)
	}()

	go func() {
		defer wg.Done()
		call := <-call2.Done
		if call.Error != nil {
			checkError(call.Error)
			return
		}
		reply := call.Reply.(*WordCountReply)
		fmt.Println("Result 2:", reply.Counts)
	}()

	fmt.Println("Requests sent, waiting for results...")

	wg.Wait()
	fmt.Println("All results received.")
}
