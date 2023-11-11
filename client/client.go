package main

import (
	proto "Distributed-Mutual-Exclusion/grpc"
	"context"
	"log"
	"os"

	"github.com/AlbertRossJoh/itualgs_go/fundamentals/queue"
)

type State int64

const (
	HELD     State = 0
	WANTED   State = 1
	RELEASED State = 2
)

type Client struct {
	proto.UnimplementedClientServiceServer
	id string
}

var (
	state        State = RELEASED
	lamport      int64 = 0
	messageQueue       = queue.NewQueue[*proto.Request](1024)
)

func main() {
	clientIpAddr := os.Getenv("HOSTNAME")
	log.Printf("The client's ip address is: ", clientIpAddr)

	f, err := os.OpenFile("/var/nas/shared", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Couldn't open the shared file! On client ", clientIpAddr)
	}
	defer f.Close()
	if _, err := f.WriteString(clientIpAddr + "\n"); err != nil {
		log.Println("Couldn't write to shared file! On client ", clientIpAddr)
	}

	state = WANTED
	// Do stuff
	state = HELD
}

func received(req *proto.Request) *proto.Response {
	if state == HELD || (state == WANTED && lamport < req.LamportTs) {
		messageQueue.Enqueue(req)
		return nil
	} else {
		return &proto.Response{
			Status: 200,
		}
	}
}

func exit() {
	state = RELEASED
	// Loop over queue and reply
}

func critical_function() {
	for {
		if state == WANTED {
			for !messageQueue.IsEmpty() {

			}
			log.Println("Ohh no critical function")
		}
	}
}

func (c *Client) MakeRequest(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	return received(in), nil
}
