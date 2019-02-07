package handlers

import "fmt"

func StartDispatchEvents(queue chan FileChangeEvent) {
	for event := range queue {
		fmt.Println(event) // todo: replay locally...
		fmt.Println()
	}
}
