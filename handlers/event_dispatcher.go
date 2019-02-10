package handlers

import (
	"fmt"
	"os"
	"strings"
)

func StartDispatchEvents(queue chan FileChangeEvent) {
	for event := range queue {
		//fmt.Println(event) // todo: replay locally...
		if strings.HasSuffix(event.File(), ".txt") ||
			strings.HasSuffix(event.File(), ".jar") ||
			strings.HasSuffix(event.File(), ".srcjar") ||
			strings.HasSuffix(event.File(), ".xml") {
			if event.HasType(EtDelete) || event.HasType(EtMovedFrom) {
				err := remove(event.LocalFilePath())
				if err != nil {
					fmt.Println("ERROR:", err)
				}
			} else {
				err := markForSync(event.RelativeFilePath())
				if err != nil {
					fmt.Println("ERROR:", err)
				}

			}
			//fmt.Println()
		}
	}
}

func remove(path string) error {
	return os.RemoveAll(path)
}

func markForSync(path string) error {
	f, err := os.OpenFile("/tmp/rsync.spec", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\r\n", path)); err != nil {
		return err
	}

	return nil
}
