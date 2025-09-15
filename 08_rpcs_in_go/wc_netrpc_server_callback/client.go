package main

import "net/rpc"

type Result struct {
	Counts map[string]int
	Err    error
}

func makeRequest(input string, serverAddr string) *rpc.Call {
	// Establish connection to the RPC server - A client
	client, err := rpc.Dial("tcp", serverAddr)
	checkError(err)

	// Prepare the arguments and reply structure
	args := WordCountRequest{Input: input}
	reply := WordCountReply{make(map[string]int)}

	return client.Go("WordCountServer.Compute", args, &reply, nil)
}
