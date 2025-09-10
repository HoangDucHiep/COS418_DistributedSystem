package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args struct{ Amount int }

func callWithRetry(amount int) {
	for i := 0; i < 2; i++ { // thử 2 lần
		conn, err := net.DialTimeout("tcp", "localhost:1234", 2*time.Second)
		if err != nil {
			log.Printf("Dial error: %v", err)
			continue
		}
		client := rpc.NewClient(conn)
		defer client.Close()

		args := Args{Amount: amount}
		var reply int

		err = client.Call("Bank.Debit", args, &reply)
		if err != nil {
			log.Printf("Call error: %v", err)
		} else {
			fmt.Printf("Debit %d thành công, số dư còn: %d\n", amount, reply)
			return
		}
	}
	log.Println("Gọi RPC thất bại sau khi retry")
}

func main() {
	callWithRetry(10)
}
