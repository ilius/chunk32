// Copyright (c) 2017 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ilius/crock32"
)

func encodeFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	// fmt.Println("------------------ Hex Data ------------------")
	// fmt.Println(hex.EncodeToString(input))
	// fmt.Println("----------------- Croc32 ------------------")
	text := crock32.Encode(input)
	// TODO: add a flag to keep it uppercase
	text = strings.ToLower(text)
	fmt.Println(text)
	// fmt.Println("--------------------------------------------")
}

func decodeFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	input = bytes.TrimSpace(input)
	output, err := crock32.Decode(string(input))
	if err != nil {
		panic(err)
	}
	// TODO: add a flag to print hex-encoded
	fmt.Println(string(output))
	// fmt.Println(hex.EncodeToString(output))
}

func main() {
	// TODO: parse flags, support -c (code) and -d (decode)
	decodeFlag := flag.Bool(
		"d",
		false,
		"crock32-cli -d",
	)

	flag.Parse()

	if decodeFlag != nil && *decodeFlag {
		decodeFromStdin()
	} else {
		encodeFromStdin()
	}
}
