package main

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type WordCountServer struct {
	addr string
}

type WordCountRequest struct {
	Input string
}

type WordCountReply struct {
	Counts map[string]int
}
