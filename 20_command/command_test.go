package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	// 事件通道
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}

	}()
	defer close(eventChan)

	// 命令缓冲通道
	commands := make(chan Command, 1000)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = StartCommandFunc()
			case "archive":
				command = ArchiveCommandFunc()
			}
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			_ = c()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}