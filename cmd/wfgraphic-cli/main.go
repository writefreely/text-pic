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
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/writeas/text-pic"
	"github.com/writeas/web-core/log"
)

var (
	outputFile = flag.String("o", "out.png", "Image output filename")
	font       = flag.String("font", "serif", "Post font (options: \"serif\", \"sans\", \"mono\")")
	instance   = flag.String("i", "write.as", "WriteFreely instance hostname (e.g. pencil.writefree.ly)")
	author     = flag.String("u", "", "WriteFreely author username (for multi-user instances)")
	size       = flag.String("size", "1024", "Image size, either a single number for a square (e.g. \"900\") or a combined width and height (e.g. \"1080x1920\")")
)

func main() {
	log.Info("Starting...")
	flag.Parse()

	// Validate input
	if !textpic.IsValidFont(*font) {
		log.Info("Invalid font given. Options: \"serif\", \"sans\", \"mono\"")
		os.Exit(1)
	}
	// Parse image dimensions and validate
	var w, h int
	if strings.ContainsRune(*size, 'x') {
		parts := strings.Split(*size, "x")
		if len(parts) != 2 {
			log.Info("Invalid --size given. Must be a single number for a square (e.g. \"900\") or a combined width and height (e.g. \"1080x1920\")")
			os.Exit(1)
		}
		var err error
		w, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Info("Unable to parse --size dimension '%s': %s", parts[0], err)
			os.Exit(1)
		}
		h, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Info("Unable to parse --size dimension '%s': %s", parts[1], err)
			os.Exit(1)
		}
	} else {
		sq, err := strconv.Atoi(*size)
		if err != nil {
			log.Info("Unable to parse --size: %s", err)
			os.Exit(1)
		}
		w = sq
		h = sq
	}

	log.Info("Reading input...")
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Error("read: %s", err)
		os.Exit(1)
	}

	log.Info("Generating image...")
	start := time.Now()
	err = textpic.GenerateImage(textpic.NewContentOptions(*instance, *author, false, *font, string(in)), w, h, *outputFile)
	if err != nil {
		log.Error("%s", err)
		os.Exit(1)
	}
	log.Info("Completed in %s", time.Since(start))
}
