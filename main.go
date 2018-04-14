// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import "flag"

func main() {
	// TODO: parse flags, support -c (code) and -d (decode)
	decodeFlag := flag.Bool(
		"d",
		false,
		"chunk32 -d",
	)

	upperFlag := flag.Bool(
		"u",
		false,
		"chunk32 -u",
	)

	flag.Parse()

	if decodeFlag != nil && *decodeFlag {
		decodeFromStdin()
	} else {
		upper := upperFlag != nil && *upperFlag
		encodeFromStdin(upper)
	}
}
