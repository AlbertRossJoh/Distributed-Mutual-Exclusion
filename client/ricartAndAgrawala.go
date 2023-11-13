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
	go someCritFunc()
	state = WANTED
	multicast()
	for replies < N-1 {
		log.Println(replies)
		time.Sleep(time.Duration(5) * time.Second)
	}
	state = HELD
}

func multicast() {
	for _, str := range GetFileContents() {
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
	log.Printf("other %d : own %d", req.LamportTs, lamport)
	if state == HELD || (state == WANTED && lamport < req.LamportTs) {
		replyQueue.Enqueue(req)
		lamport = req.LamportTs + 1
	} else {
		lamport++
		replyTo(req.Id)
	}
}

// On exit do
// state := RELEASED reply to all in queue
// End on
func exit() {
	state = RELEASED
	for _, req := range replyQueue.Items() {
		replyTo(req.Id)
	}
	replyQueue.Clear()
}

func someCritFunc() {
	for {
		if state == HELD {
			replies = 0
			log.Println("Ohh no critical function")
			time.Sleep(time.Duration(5) * time.Second)
			exit()
		}
	}
}
