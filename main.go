// Copyright (c) 2017 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import "flag"

func main() {
	// TODO: parse flags, support -c (code) and -d (decode)
	decodeFlag := flag.Bool(
		"d",
		false,
		"chunk32 -d",
	)

	flag.Parse()

	if decodeFlag != nil && *decodeFlag {
		decodeFromStdin()
	} else {
		encodeFromStdin()
	}
}
