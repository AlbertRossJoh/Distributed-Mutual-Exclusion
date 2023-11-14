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
		if GetReplies() >= N-1 {
			break
		}
	}
	state = HELD
	someCritFunc()
	SetReplies(0)
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
	if state == HELD || (state == WANTED && GetLogicalTimeStamp() < req.LamportTs) {
		replyQueue.Enqueue(req)
	} else {
		replyTo(req.Id, req.LamportTs)
	}
}

// On exit do
// state := RELEASED reply to all in queue
// End on
func exit() {
	state = RELEASED
	for i := 0; i < replyQueue.Size(); i++ {
		res, _ := replyQueue.Dequeue()
		replyTo(res.Id, res.LamportTs)
	}
}

func someCritFunc() {
	log.Println("______________________________________")
	log.Println("Critical function")
	log.Printf("Current critical timestamp is: %d", GetLogicalTimeStamp())
	log.Println("______________________________________")
	time.Sleep(time.Duration(3) * time.Second)
	exit()
}
