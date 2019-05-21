// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import "flag"

func main() {
	// TODO: parse flags, support -c (code) and -d (decode)
	decodeFlag := flag.Bool(
		"d",
		false,
		"Decode:\nchunk32 -d",
	)

	upperFlag := flag.Bool(
		"u",
		false,
		"Use uppercase characters (when encoding)\nchunk32 -u",
	)

	checkFlag := flag.Bool(
		"check",
		false,
		"Add check sybmol (when decoding)\nchunk32 -check",
	)

	noNewlineFlag := flag.Bool(
		"n",
		false,
		"Do not print newline at the end (mostly useful for decode)\nchunk32 -d -n",
	)

	flag.Parse()

	noNewline := noNewlineFlag != nil && *noNewlineFlag
	if decodeFlag != nil && *decodeFlag {
		decodeFromStdin(noNewline)
	} else {
		upper := upperFlag != nil && *upperFlag
		check := checkFlag != nil && *checkFlag
		encodeFromStdin(upper, check, noNewline)
	}
}
