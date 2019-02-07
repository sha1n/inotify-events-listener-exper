package main

import (
	"bufio"
	"fmt"
	"github.com/sha1n/inotify-events-listener-exper/parser"
	"net"
)

const ServerPort = 8081
const EventBufferSize = 1014
const HandlerCount = 2

func main() {
	fmt.Println("Launching TCP server...")

	eventQueue := make(chan *parser.Event, EventBufferSize)
	startEventHandlers(eventQueue)
	acceptConnections(eventQueue)
}

func acceptConnections(eventQueue chan *parser.Event) {
	localServerAddress := fmt.Sprintf(":%d", ServerPort)
	ln, err := net.Listen("tcp", localServerAddress)
	if err != nil {
		fmt.Println( "Failed to open server socket on port", ServerPort, err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept incoming connection:", err)
			panic(err)
		}
		go handleClientConnection(conn, eventQueue)
	}
}

func startEventHandlers(eventQueue chan *parser.Event) {
	for i := 0; i < HandlerCount; i++ {
		go handleEvents(eventQueue)
	}
}

func handleEvents(queue chan *parser.Event) {
	for event := range queue {
		fmt.Println(event) // todo: replay locally...
		fmt.Println()
	}
}

func handleClientConnection(conn net.Conn, queue chan *parser.Event) {
	for {
		scanner := bufio.NewScanner(bufio.NewReader(conn))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			eventParser := parser.NewParser()
			event, e := eventParser.Parse(scanner.Text())
			if e != nil {
				panic(e)
			} else {
				queue<-event
			}
		}
	}
}
