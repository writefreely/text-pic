/*
 * Copyright © 2021 A Bunch Tell LLC.
 *
 * This file is part of text-pic.
 *
 * text-pic is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

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
