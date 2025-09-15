package main

import (
	"log"
	"net"
	"net/rpc"
)

// Args phải export field (chữ hoa) để encoding/gob truy cập được.
type Args struct {
	A, B int
}

type Arith int

// Phương thức phải có dạng: (t *T) Method(args *Args, reply *Type) error
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	if err := rpc.Register(arith); err != nil {
		log.Fatalf("rpc.Register error: %v", err)
	}

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	defer l.Close()
	log.Println("RPC server listening on :1234")

	for {
		conn, err := l.Accept()
		if err != nil {
			// nếu Accept lỗi, in log rồi tiếp tục vòng lặp
			log.Println("Accept error:", err)
			continue
		}
		// xử lý connection trong goroutine để server có thể phục vụ nhiều client
		go func(c net.Conn) {
			defer c.Close()
			log.Printf("Accepted connection from %s\n", c.RemoteAddr())
			rpc.ServeConn(c)
		}(conn)
	}
}
