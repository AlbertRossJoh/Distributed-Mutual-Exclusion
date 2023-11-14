package main

import (
	proto "Distributed-Mutual-Exclusion/grpc"
	"log"
	"time"
)

const N = 3

// INIT
// STATE = RELEASED

// On enter do
// state := WANTED;
// “multicast ‘req(T,p)’”, where T := LAMPORT time of ‘req’ at p wait for N-1 replies
// state := HELD;
// End on

func enter() {
	state = WANTED
	multicast()
	for {
		currRep := <-replies
		if currRep >= N-1 {
			break
		}
		replies <- currRep
	}
	state = HELD
	someCritFunc()
	replies <- 0
}

func multicast() {
	res := GetFileContents()
	for _, str := range res {
		if str != clientIpAddr {
			makeCritRequest(str)
		}
	}
}

// On receive ‘req (Ti,pi)’do
//
//	if(state == HELD || (state == WANTED && (T,pme) < (Ti,pi)))
//	then queue req
//	else reply to req
//
// End on
func receive(req *proto.Request) {
	curr := <-lamport
	if curr < req.LamportTs {
		curr = req.LamportTs
	}
	curr++
	lamport <- curr
	if state == HELD || (state == WANTED && curr < req.LamportTs) {
		replyQueue.Enqueue(req)
	} else {
		replyTo(req.Id)
	}
}

// On exit do
// state := RELEASED reply to all in queue
// End on
func exit() {
	state = RELEASED
	for i := 0; i < replyQueue.Size(); i++ {
		res, _ := replyQueue.Dequeue()
		replyTo(res.Id)
	}
}

func someCritFunc() {
	curr := <-lamport
	lamport <- curr
	log.Println("______________________________________")
	log.Println("Critical function")
	log.Printf("Current critical timestamp is: %d", curr)
	log.Println("______________________________________")
	time.Sleep(time.Duration(3) * time.Second)
	exit()
}
