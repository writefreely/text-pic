/*
 * Copyright © 2021 A Bunch Tell LLC.
 *
 * This file is part of text-pic.
 *
 * text-pic is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package textpic

import (
	"fmt"
	"image/color"
	"path/filepath"

	"github.com/fogleman/gg"
)

func loadFont(dc *gg.Context, bold bool, points float64) error {
	fontLoraBoldPath := filepath.Join("fonts", "Lora-Bold.ttf")
	fontLoraPath := filepath.Join("fonts", "Lora-Regular.ttf")
	path := fontLoraPath
	if bold {
		path = fontLoraBoldPath
	}

	err := dc.LoadFontFace(path, points)
	if err != nil {
		return fmt.Errorf("load bold font: %s", err)
	}
	return nil
}

func Run() error {
	w := 900
	h := 900
	wf := float64(w)
	hf := float64(h)
	dc := gg.NewContext(w, h)
	dc.DrawRectangle(0, 0, wf, hf)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	// Define margins for footer text
	footerFontSize := 32.0
	footerMargin := 20.0
	x := footerMargin
	y := footerMargin
	footerMarginY := 20.0

	// Content parameters
	contentFontSize := 48.0
	lineSpacing := 1.8
	contentBottomMargin := 100.0
	contentRightMargin := 50.0
	contentTopMargin := 50.0
	contentWidth := wf - contentRightMargin - contentRightMargin

	// Create bold instance name
	err := loadFont(dc, true, footerFontSize)
	if err != nil {
		return err
	}
	instance := "write.as"
	baseTextWidth, textHeight := dc.MeasureString(instance)

	// Create user path
	err = loadFont(dc, false, footerFontSize)
	if err != nil {
		return err
	}
	dc.SetColor(color.Black)

	userPath := "/matt"
	userTextWidth, _ := dc.MeasureString(userPath)
	// x = canvas halfway point - total text width halfway point
	x = wf/2 - (baseTextWidth+userTextWidth)/2
	y = hf - textHeight - footerMarginY
	err = loadFont(dc, true, footerFontSize)
	if err != nil {
		return err
	}
	dc.DrawString(instance, x, y)

	// x = original x coordinate + base text width
	x += baseTextWidth
	y = hf - textHeight - footerMarginY
	err = loadFont(dc, false, footerFontSize)
	if err != nil {
		return err
	}
	dc.DrawString(userPath, x, y)

	// Draw the content
	err = loadFont(dc, false, contentFontSize)
	if err != nil {
		return err
	}
	s := "The rest of the travelers in our flying bus napped or stared listlessly at a shiny slab in their lap and the staring yellow orb morphed into a full circle out in the blue. As we banked to the right — a nod to its awakening — it seemed to rest in acknowledgement, hanging for a moment on the invisible horizon."
	lines := dc.WordWrap(s, contentWidth)
	linesStr := ""
	for i, str := range lines {
		linesStr += str
		if i != len(lines)-1 {
			linesStr += "\n"
		}
	}
	_, contentTextHeight := dc.MeasureMultilineString(linesStr, lineSpacing)
	x = contentRightMargin
	y = contentTopMargin - contentBottomMargin + hf/2 - contentTextHeight/2
	dc.DrawStringWrapped(s, x, y, 0, 0, contentWidth, lineSpacing, gg.AlignLeft)

	err = dc.SavePNG("out.png")
	if err != nil {
		return fmt.Errorf("save png: %s", err)
	}
	return nil
}
