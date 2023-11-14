package main

import "sync"

var (
	lampertMu        sync.Mutex
	repliesMu        sync.Mutex
	logicalTimeStamp int64 = 0
	amountOfReplies        = 0
)

func ConditionalUpdateLogicalTimeStamp(other int64) {
	lampertMu.Lock()
	if logicalTimeStamp < other {
		logicalTimeStamp = other
	}
	logicalTimeStamp++
	lampertMu.Unlock()
}

func UpdateLogicalTimeStamp() {
	lampertMu.Lock()
	logicalTimeStamp++
	lampertMu.Unlock()
}

func GetLogicalTimeStamp() int64 {
	lampertMu.Lock()
	ts := logicalTimeStamp
	lampertMu.Unlock()
	return ts
}

func UpdateReplies() {
	repliesMu.Lock()
	amountOfReplies++
	repliesMu.Unlock()
}

func GetReplies() int {
	repliesMu.Lock()
	ts := amountOfReplies
	repliesMu.Unlock()
	return ts
}

func SetReplies(a int) {
	repliesMu.Lock()
	amountOfReplies = a
	repliesMu.Unlock()
}
