package main

import (
	proto "Distributed-Mutual-Exclusion/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/AlbertRossJoh/itualgs_go/fundamentals/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type State int64

const (
	HELD     State = 0
	WANTED   State = 1
	RELEASED State = 2
)

type Client struct {
	proto.UnimplementedClientServiceServer
	id   string
	port int
}

const SERVER_PORT = 6969

var (
	state        State      = RELEASED
	lamport      chan int64 = make(chan int64, 1)
	replyQueue              = queue.NewBufQueue[*proto.Request](1024)
	clientIpAddr            = os.Getenv("HOSTNAME")
	replies      chan int   = make(chan int, 1)
)

func main() {
	// go printState()
	WriteToSharedFile()
	replies <- 0
	lamport <- 0
	client := Client{
		id:   clientIpAddr,
		port: SERVER_PORT,
	}
	go startClient(client)

	for {
		time.Sleep(time.Duration(5) * time.Second)
		enter()
	}
}

func startClient(client Client) {
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(client.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", client.port)

	// Register the grpc server and serve its listener
	proto.RegisterClientServiceServer(grpcServer, &client)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatal("Could not serve listener")
	}
}

func (c *Client) MakeRequest(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	receive(in)
	return &proto.Response{
		Status: 200,
	}, nil
}

func (c *Client) Reply(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	currRep := <-replies
	currRep++
	replies <- currRep
	return &proto.Response{
		Status: 200,
	}, nil
}

func makeCritRequest(str string) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", str, strconv.Itoa(SERVER_PORT)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Could not connect to client: ", str)
	}
	client := proto.NewClientServiceClient(conn)
	curr := <-lamport
	curr++
	lamport <- curr
	client.MakeRequest(context.Background(), &proto.Request{
		State:     proto.State_WANTED,
		LamportTs: curr,
		Id:        clientIpAddr,
	})
}

func replyTo(clientIp string) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", clientIp, strconv.Itoa(SERVER_PORT)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Could not connect to client: ", clientIp)
	}
	// log.Printf("Replying to: %s", clientIp)
	client := proto.NewClientServiceClient(conn)
	curr := <-lamport

	curr++
	lamport <- curr
	client.Reply(context.Background(), &proto.Request{
		State:     proto.State_RELEASED,
		LamportTs: curr,
		Id:        clientIpAddr,
	})
}

func printState() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		switch state {
		case HELD:
			log.Printf("Current state is HELD")
		case RELEASED:
			log.Printf("Current state is RELEASED")
		case WANTED:
			log.Printf("Current state is WANTED")
		}
		curr := <-lamport
		lamport <- curr
		log.Printf("Current timestamp is: %d", curr)
	}
}
