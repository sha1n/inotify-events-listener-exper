package handlers

import (
	"encoding/csv"
	"fmt"
	"strings"
)

var eventTypeMapping = make(map[string]EventType)

func init() {
	eventTypeMapping["ACCESS"] = EtAccess
	eventTypeMapping["MODIFY"] = EtModify
	eventTypeMapping["ATTRIB"] = EtAttrib
	eventTypeMapping["CLOSE_WRITE"] = EtCloseWrite
	eventTypeMapping["CLOSE_NOWRITE"] = EtCloseNoWrite
	eventTypeMapping["CLOSE"] = EtClose
	eventTypeMapping["OPEN"] = EtOpen
	eventTypeMapping["MOVED_TO"] = EtMovedTo
	eventTypeMapping["MOVED_FROM"] = EtMovedFrom
	eventTypeMapping["MOVE"] = EtMove
	eventTypeMapping["MOVE_SELF"] = EtMoveSelf
	eventTypeMapping["CREATE"] = EtCreate
	eventTypeMapping["DELETE"] = EtDelete
	eventTypeMapping["DELETE_SELF"] = EtDeleteSelf
	eventTypeMapping["UNMOUNT"] = EtUnmount
}

type Parser interface {
	Parse(csv string) (FileChangeEvent, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (parser *parser) Parse(csvLine string) (rtn FileChangeEvent, failure error) {
	eventRecord, _ := csv.NewReader(strings.NewReader(csvLine)).Read()
	event := &INotifyEvent{}

	fieldCount := len(eventRecord)
	switch fieldCount {
	case 2:
		event, failure = parseCommonFields(event, eventRecord)
	case 3:
		event, failure = parseCommonFields(event, eventRecord)
		event.file = eventRecord[2]
	default:
		failure = fmt.Errorf("WTF?! Unexpected number of fields on event: %s", csvLine)
	}

	return event, failure
}

func parseCommonFields(event *INotifyEvent, eventRecord []string) (*INotifyEvent, error) {
	event.dir = eventRecord[0]
	protocolEventTypes, err := csv.NewReader(strings.NewReader(eventRecord[1])).Read()
	if err == nil {
		event.types = toEventTypes(protocolEventTypes)
	}

	return event, err
}

func toEventTypes(protocolEventTypes []string) []EventType {
	types := make([]EventType, len(protocolEventTypes))

	for i := range protocolEventTypes {
		types[i] = eventTypeMapping[protocolEventTypes[i]]
	}

	return types
}
