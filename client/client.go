package client

import (
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

var (
	state        State = RELEASED
	lamport      int64 = 0
	messageQueue       = queue.NewQueue[int64](1024)
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

func received(reqLamport int64) {
	if state == HELD || (state == WANTED && lamport < reqLamport) {
		messageQueue.Enqueue(1)
	} else {
		// reply
	}
}

func exit() {
	state = RELEASED
	// Loop over queue and reply
}

func critical_function() {
	for {
		for !messageQueue.IsEmpty() {

		}
		log.Println("Ohh no critical function")
	}
}
