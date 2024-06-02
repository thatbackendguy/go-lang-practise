package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a new ZeroMQ context
	context, _ := zmq.NewContext()
	defer context.Term()

	// Create a new PUB socket
	//socket, _ := context.NewSocket(zmq.PUB)
	socket, _ := context.NewSocket(zmq.PULL)
	defer socket.Close()

	// Bind to port 9090
	socket.Bind("tcp://10.20.40.227:9090")

	// Publish messages
	//for {
	//	message := "Hello from Go" + time.Now().Format("2006-01-02 15:04:05")
	//	socket.Send(message, 0)
	//	fmt.Println("Published message:", message)
	//	time.Sleep(10 * time.Second) // Sleep for a second before sending the next message
	//}

	for {
		message, err := socket.Recv(0)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			break
		}
		fmt.Println("Received message:", message)
	}
}
