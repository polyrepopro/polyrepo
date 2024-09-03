package main

import (
	"context"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		if err := exec.CommandContext(ctx, "sleep", "59999").Run(); err != nil {
			// This will fail after 100 milliseconds. The 5 second sleep
			// will be interrupted.
		}
		time.Sleep(100 * time.Second)
	}()
	println("done")
	select {}
}
