package main

import "net/rpc"

type Result struct {
	Counts map[string]int
	Err    error
}

func makeRequest(input string, serverAddr string) chan Result {
	// Establish connection to the RPC server - A client
	client, err := rpc.Dial("tcp", serverAddr)
	checkError(err)

	// Prepare the arguments and reply structure
	args := WordCountRequest{Input: input}
	reply := WordCountReply{make(map[string]int)}

	// create a channel for asynchronous call
	ch := make(chan Result)

	// Make the RPC call to the Compute method on the server
	go func() {
		err = client.Call("WordCountServer.Compute", args, &reply)

		if err != nil {
			ch <- Result{nil, err}
		} else {
			ch <- Result{reply.Counts, nil}
		}
	}()

	return ch
}
