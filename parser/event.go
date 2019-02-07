package parser

import "fmt"

type Event struct {
	dir     string
	file    string
	actions []string
}

func (e *Event) String() string {
	return fmt.Sprintf(
		`EVENT:
dir:     %s
file:    %s
actions: %s`,
		e.dir,
		e.file,
		e.actions)
}
