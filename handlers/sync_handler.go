package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type SyncHandler struct {
	changesChan chan string
	outFileName string
}

func NewSyncHandler() *SyncHandler {
	return &SyncHandler{
		changesChan: make(chan string, 1000),
		outFileName: "/tmp/rsync_files",
	}
}

func (h *SyncHandler) RegisterRelativePath(fileName string) {
	h.changesChan <- fileName
}

func (h *SyncHandler) Start() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case file := <-h.changesChan:
				fmt.Println("mod: ", file)
				err := h.append(file)
				if err != nil {
					panic(err)
				}
			case <-ticker.C:
				if _, err := os.Stat(h.outFileName); os.IsNotExist(err) {
					//fmt.Println("Skipping rsync - no changes reported.")
				} else {
					syncErr := sync("/bazel-rsynced-outputs/outputs", h.outFileName)
					if syncErr != nil {
						fmt.Println("Rsync failed:", syncErr)
					} else {
						fmt.Println("Rsync success!")
						delErr := os.Remove(h.outFileName)
						if delErr != nil {
							fmt.Println(delErr, "Failed to delete output file...")
						}
					}
				}
			}
		}
	}()
}

func (h *SyncHandler) append(fileName string) error {
	f, err := os.Create(h.outFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\r\n", fileName)); err != nil {
		return err
	}

	return nil
}

func sync(target string, filesFrom string) error {
	// "/bazel-rsynced-outputs/outputs"
	cmd := exec.Command("rsync", "--port",
		"873",
		"--files-from="+filesFrom,
		"--delete",
		"--delete-excluded",
		"--archive",
		"bazel@127.0.0.1::output",
		target,
		"--password-file",
		"/bazel-nfs-volume/.wazeltmp/bazelPass", // fixme shai: to config
	)

	return cmd.Run()
}
