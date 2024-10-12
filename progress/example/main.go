package main

import (
	"log"
	"time"

	"github.com/dGilli/terminal-ui/progress"
)

func main() {
	p := progress.New(progress.Config{})

	log.Println("Starting the progress bar")

	p.Update(50)

	time.Sleep(time.Second * 1)
	p.Update(100)

	log.Println("Progress complete")
}
