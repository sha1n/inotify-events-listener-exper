package handlers

import "fmt"

type FileChangeEvent interface {
	Dir() string
	File() string
	Types() []EventType
	String() string
}

type EventType int

const (
	EtAccess       EventType = 0
	EtModify       EventType = 1
	EtAttrib       EventType = 2
	EtCloseWrite   EventType = 3
	EtCloseNoWrite EventType = 4
	EtClose        EventType = 5
	EtOpen         EventType = 6
	EtMovedTo      EventType = 7
	EtMovedFrom    EventType = 8
	EtMove         EventType = 9
	EtMoveSelf     EventType = 10
	EtCreate       EventType = 11
	EtDelete       EventType = 12
	EtDeleteSelf   EventType = 13
	EtUnmount      EventType = 14
)

var names = []string {
	"ACCESS",
	"MODIFY",
	"ATTRIB",
	"CLOSE_WRITE",
	"CLOSE_NOWRITE",
	"CLOSE",
	"OPEN",
	"MOVED_TO",
	"MOVED_FROM",
	"MOVE",
	"MOVE_SELF",
	"CREATE",
	"DELETE",
	"DELETE_SELF",
	"UNMOUNT",
}


type INotifyEvent struct {
	dir   string
	file  string
	types []EventType
}

func (e *INotifyEvent) Dir() string {
	return e.dir
}

func (e *INotifyEvent) File() string {
	return e.file
}

func (e *INotifyEvent) Types() []EventType {
	return e.types
}

func (e *INotifyEvent) String() string {
	return fmt.Sprintf(
		`event :
dir   : %s
file  : %s
types : %s`,
		e.dir,
		e.file,
		e.types)
}

func (t EventType) String() string {
	if t < EtAccess || t > EtUnmount {
		return "BUG: Unknown type!"
	}

	return names[t]
}
