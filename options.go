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

var fonts = map[string]string{
	"norm":  "Lora",
	"serif": "Lora",
	"sans":  "OpenSans",
	"mono":  "Hack",
	"wrap":  "Hack",
}

type ContentOptions struct {
	// Author information
	Instance string
	Username string

	// Write.as-only option
	IsSubdomain bool // UNIMPLEMENTED

	// Content
	UserFont string
	Content  string
}

func NewContentOptions(instance, username string, isSubdomain bool, font, content string) *ContentOptions {
	opt := &ContentOptions{
		Instance:    instance,
		Username:    username,
		IsSubdomain: isSubdomain,
		UserFont:    font,
		Content:     content,
	}
	if opt.Instance == "" {
		opt.Instance = "write.as"
	}
	if opt.Content == "" {
		opt.Content = "Hello, world!"
	}
	return opt
}

func IsValidFont(f string) bool {
	_, valid := fonts[f]
	return valid
}
