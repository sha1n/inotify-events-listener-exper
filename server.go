package main

import (
	"fmt"
	"github.com/sha1n/inotify-events-listener-exper/handlers"
	"net"
)

const ServerPort = 8081
const EventBufferSize = 1014
const HandlerCount = 2

func main() {
	fmt.Println("Launching TCP server...")

	eventQueue := make(chan handlers.FileChangeEvent, EventBufferSize)
	startEventHandlers(eventQueue)
	startAcceptingEvents(eventQueue)
}

func startAcceptingEvents(eventQueue chan handlers.FileChangeEvent) {
	localServerAddress := fmt.Sprintf(":%d", ServerPort)
	ln, err := net.Listen("tcp", localServerAddress)
	if err != nil {
		fmt.Println("Failed to open server socket on port", ServerPort, err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept incoming connection:", err)
			panic(err)
		}
		go handlers.HandleClientConnection(conn, eventQueue)
	}
}

func startEventHandlers(eventQueue chan handlers.FileChangeEvent) {
	for i := 0; i < HandlerCount; i++ {
		go handlers.StartDispatchEvents(eventQueue)
	}
}
