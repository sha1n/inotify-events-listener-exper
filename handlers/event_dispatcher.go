package handlers

import (
	"fmt"
	"os"
)

func StartDispatchEvents(queue chan FileChangeEvent) {
	syncHandler := NewSyncHandler()
	syncHandler.Start()
	for event := range queue {
		//if strings.HasSuffix(event.File(), ".txt") ||
		//	strings.HasSuffix(event.File(), ".jar") ||
		//	strings.HasSuffix(event.File(), ".srcjar") ||
		//	strings.HasSuffix(event.File(), ".xml") {
			if event.HasType(EtDelete) || event.HasType(EtDeleteSelf) || event.HasType(EtMovedFrom) {
				err := remove(event.LocalFilePath())
				if err != nil {
					fmt.Println("ERROR:", err)
				}
			} else {
				syncHandler.RegisterRelativePath(event.RelativeFilePath())

			}
		//}
	}
}

func remove(path string) error {
	return os.RemoveAll(path)
}
