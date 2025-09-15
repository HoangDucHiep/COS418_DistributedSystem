package main

import (
	"net"
	"net/rpc"
	"strings"
)

// Stub function for word count computation
// It takes a WordCountRequest and fills in the WordCountReply
// with the word counts.
// This function can be called remotely via RPC.
// The input string is split into words based on whitespace.
// The counts of each word are stored in a map and returned in the reply.
// Example: input "hello world hello" results in map{"hello": 2, "world": 1}
// This function does not modify the server state, hence it has a pointer receiver.
// It returns an error if any issues occur during processing.
// This function is thread-safe as it does not modify any shared state.
// It can be called concurrently by multiple clients.
func (*WordCountServer) Compute(request WordCountRequest, reply *WordCountReply) error {

	counts := make(map[string]int)
	input := request.Input
	tokens := strings.Fields(input)
	for _, t := range tokens {
		counts[t]++
	}

	reply.Counts = counts
	return nil
}

func (server *WordCountServer) Listen() {
	rpc.Register(server)
	listener, err := net.Listen("tcp", server.addr)
	checkError(err)

	go func() {
		rpc.Accept(listener)
	}()
}
