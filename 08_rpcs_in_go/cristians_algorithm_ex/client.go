package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"
)

func SyncTime() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// T1: Client sends request

	args := TimeRequest{}
	res := TimeResponse{}

	T1 := time.Now().UnixNano() / int64(time.Millisecond)
	err = client.Call("TimeServer.GetTime", args, &res)

	if err != nil {
		panic(err)
	}

	// T4: Client receives response
	T4 := time.Now().UnixNano() / int64(time.Millisecond)
	RTT := (T4 - T1) - (res.T3 - res.T2)
	Offset := ((res.T2 - T1) + (res.T3 - T4)) / 2

	adjustedClientTime := res.T3 + Offset

	fmt.Printf("T1=%d, T2=%d, T3=%d, T4=%d\n", T1, res.T2, res.T3, T4)
	fmt.Println("Round Trip Time (RTT):", RTT, "ms")
	fmt.Println("Clock Offset:", Offset, "ms")
	fmt.Println("Adjusted Client Time:", adjustedClientTime, "ms since epoch")
}

func main() {
	server := TimeServer{}

	rpc.Register(&server)
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	go func() {
		rpc.Accept(listener)
	}()

	time.Sleep(2 * time.Second) // Give the server a moment to start

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		SyncTime()
	}()

	wg.Wait()
	fmt.Println("Time synchronization complete.")
}
