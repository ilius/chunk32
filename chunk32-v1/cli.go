// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func encodeFromStdin(upper bool, check bool, noNewline bool) {
	input, _ := io.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	input = bytes.TrimRight(input, "\n")
	// fmt.Println("------------------ Hex Data ------------------")
	// fmt.Println(hex.EncodeToString(input))
	// fmt.Println("----------------- Chunk58 ------------------")
	text := Chunk32Encode(input, check)
	if !upper {
		text = strings.ToLower(text)
	}
	if noNewline {
		fmt.Print(text)
	} else {
		fmt.Println(text)
	}
}

func decodeFromStdin(noNewline bool) {
	input, _ := io.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	input = bytes.TrimRight(input, "\n")
	output, err := Chunk32Decode(string(input))
	if err != nil {
		panic(err)
	}
	// TODO: add a flag to print hex-encoded
	if noNewline {
		fmt.Print(string(output))
	} else {
		fmt.Println(string(output))
	}
	// fmt.Println(hex.EncodeToString(output))
}
