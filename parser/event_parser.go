package parser

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type Parser interface {
	Parse(csv string) (*Event, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (parser *parser) Parse(csvLine string) (event *Event, failure error) {
	eventRecord, _ := csv.NewReader(strings.NewReader(csvLine)).Read()
	event = &Event{}

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

func parseCommonFields(event *Event, eventRecord []string) (*Event, error) {
	event.dir = eventRecord[0]
	actions, err := csv.NewReader(strings.NewReader(eventRecord[1])).Read()
	if err == nil {
		event.actions = actions
	}

	return event, err
}