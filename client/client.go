package main

import (
	proto "Distributed-Mutual-Exclusion/grpc"
	"context"
	"fmt"
	"log"
	"math/rand"
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
	state           State = RELEASED
	lamport         int64 = 0
	messageQueue          = queue.NewQueue[*proto.Request](1024)
	clientIpAddr          = os.Getenv("HOSTNAME")
	canUse                = false
	amountOfClients       = 0
	replies               = 0
)

func main() {

	client := Client{
		id:   clientIpAddr,
		port: SERVER_PORT,
	}

	WriteToSharedFile()

	go startClient(client)

	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		enter()
	}
	// state = HELD
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
	return &proto.Response{
		Status: 200,
	}, nil
}

func (c *Client) Reply(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	replies++
	return &proto.Response{
		Status: 200,
	}, nil
}

// func askForCrit() {
// 	replies := 0
// 	i := 0
// 	for _, str := range GetFileContents() {
// 		res, _ := makeCritRequest(str)
// 		if res.Status == 200 {
// 			replies++
// 		}
// 		i++
// 	}
// 	log.Println(replies)
// 	log.Println(amountOfClients)
// 	log.Println(i)
// 	if i-1 > amountOfClients {
// 		amountOfClients = i - 1
// 	}
// 	if replies == amountOfClients {
// 		canUse = true
// 	}
// }

func makeCritRequest(str string) (*proto.Response, error) {
	log.Println("hello")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", str, strconv.Itoa(SERVER_PORT)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Could not connect to client: ", str)
	}
	client := proto.NewClientServiceClient(conn)
	lamport++
	return client.MakeRequest(context.Background(), &proto.Request{
		State:     proto.State_WANTED,
		LamportTs: lamport,
		Id:        clientIpAddr,
	})
}

func replyTo(clientIp string) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", clientIp, strconv.Itoa(SERVER_PORT)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Could not connect to client: ", clientIp)
	}
	client := proto.NewClientServiceClient(conn)
	lamport++
	client.Reply(context.Background(), &proto.Request{
		State:     proto.State_RELEASED,
		LamportTs: lamport,
		Id:        clientIpAddr,
	})
}

// func clientRelease() {
// 	for _, req := range messageQueue.Items() {
// 		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", req.Id, strconv.Itoa(SERVER_PORT)), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Println("Could not connect to client: ", req.Id)
// 		}
// 		client := proto.NewClientServiceClient(conn)
// 		_, err = client.Reply(context.Background(), &proto.Request{
// 			State:     proto.State_RELEASED,
// 			LamportTs: lamport,
// 			Id:        clientIpAddr,
// 		})
// 	}
// }

// func received(req *proto.Request) *proto.Response {
// 	log.Printf("Incoming ts: %d\n", req.LamportTs)
// 	log.Printf("Own ts: %d\n", lamport)
// 	if state == HELD || (state == WANTED && lamport < req.LamportTs) {
// 		messageQueue.Enqueue(req)
// 		lamport = req.LamportTs + 1
// 		return &proto.Response{
// 			Status: 400,
// 		}
// 	} else {
// 		lamport++
// 		return &proto.Response{
// 			Status: 200,
// 		}
// 	}
// }

// func exit() {
// 	state = RELEASED
// 	// Loop over queue and reply
// }

// func criticalFunction() {
// 	for {
// 		if state == WANTED && canUse {
// 			log.Println("Ohh no critical function")
// 			time.Sleep(time.Duration(5) * time.Second)
// 			canUse = false
// 			exit()
// 		}
// 	}
// }
