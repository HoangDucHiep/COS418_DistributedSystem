package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

func main() {
	// Dial với timeout để tránh treo vô hạn nếu server không phản hồi
	conn, err := net.DialTimeout("tcp", "localhost:1234", 5*time.Second)
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := Args{7, 8}
	var reply int
	// Gọi đồng bộ. Nếu muốn gọi bất đồng bộ dùng Go routines hoặc Client.Go.
	if err := client.Call("Arith.Multiply", args, &reply); err != nil {
		log.Fatalf("RPC call error: %v", err)
	}
	fmt.Printf("Result: %d\n", reply) // mong đợi 56
}
