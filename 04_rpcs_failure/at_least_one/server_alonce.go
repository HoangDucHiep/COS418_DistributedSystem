package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args struct {
	Amount int
}

type Bank int

var balance = 100

// Debit trừ vào balance
func (b *Bank) Debit(args *Args, reply *int) error {
	log.Printf((">>> Debit: %d"), args.Amount)
	balance -= args.Amount
	*reply = balance
	return nil
}

func main() {
	bank := new(Bank)
	rpc.Register(bank)

	l, _ := net.Listen("tcp", ":1234")
	log.Println("Server at-least-once listening on :1234")

	for {
		conn, _ := l.Accept()

		// giả lập server chậm để dễ thấy timeout
		go func(c net.Conn) {
			defer c.Close()
			log.Println("Client connected:", c.RemoteAddr())
			time.Sleep(10 * time.Second) // delay để client dễ timeout
			rpc.ServeConn(c)
		}(conn)
	}

}
