package handlers

import (
	"bufio"
	"net"
)

func HandleClientConnection(conn net.Conn, queue chan FileChangeEvent) {
	for {
		scanner := bufio.NewScanner(bufio.NewReader(conn))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			eventParser := NewParser()
			event, e := eventParser.Parse(scanner.Text())
			if e != nil {
				panic(e)
			} else {
				queue <- event
			}
		}
	}
}

