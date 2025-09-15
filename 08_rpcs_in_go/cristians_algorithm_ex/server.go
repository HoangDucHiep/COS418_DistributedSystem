package main

import (
	"time"
)

/*

T1: Client sends request (T1 is recorded by client)
T2: Server receives request (T2 is recorded by server)
T3: Server sends response (T3 is recorded by server)
T4: Client receives response (T4 is recorded by client)

The round-trip time (RTT) is calculated as:
RTT = (T4 - T1) - (T3 - T2)

The offset is calculated as:
Offset = ((T2 - T1) + (T3 - T4)) / 2

Client's clock can be adjusted by adding the calculated offset to its current time.
ClientTime = ClientTime + Offset

*/

type TimeServer struct{}

type TimeRequest struct{}

type TimeResponse struct {
	T2 int64
	T3 int64
}

func (t *TimeServer) GetTime(req TimeRequest, res *TimeResponse) error {
	// T2 : Server receives request

	res.T2 = time.Now().UnixNano() / int64(time.Millisecond)

	// T3 : Server sends response
	res.T3 = time.Now().UnixNano() / int64(time.Millisecond)

	return nil
}
