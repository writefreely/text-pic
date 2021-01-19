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
	"time"

	"github.com/writeas/text-pic"
	"github.com/writeas/web-core/log"
)

var (
	outputFile = flag.String("o", "out.png", "Image output filename")
	font       = flag.String("font", "serif", "Post font (options: \"serif\", \"sans\", \"mono\") - NOT IMPLEMENTED YET")
	instance   = flag.String("i", "write.as", "WriteFreely instance hostname (e.g. pencil.writefree.ly)")
	author     = flag.String("u", "", "WriteFreely author username (for multi-user instances)")
)

func main() {
	log.Info("Starting...")
	flag.Parse()

	log.Info("Reading input...")
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Error("read: %s", err)
		os.Exit(1)
	}

	log.Info("Generating image...")
	start := time.Now()
	err = textpic.GenerateImage(textpic.NewContentOptions(*instance, *author, false, *font, string(in)), *outputFile)
	if err != nil {
		log.Error("%s", err)
		os.Exit(1)
	}
	log.Info("Completed in %s", time.Since(start))
}
