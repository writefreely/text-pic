package main

import (
	"os"
	"time"

	"github.com/writeas/text-pic"
	"github.com/writeas/web-core/log"
)

func main() {
	log.Info("Starting...")
	start := time.Now()
	err := textpic.Run()
	if err != nil {
		log.Error("%s", err)
		os.Exit(1)
	}
	log.Info("Completed in %s", time.Since(start))
}
