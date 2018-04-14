// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func encodeFromStdin(upper bool) {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	// fmt.Println("------------------ Hex Data ------------------")
	// fmt.Println(hex.EncodeToString(input))
	// fmt.Println("----------------- Chunk58 ------------------")
	text := Chunk32Encode(input)
	if !upper {
		text = strings.ToLower(text)
	}
	fmt.Println(text)
	// fmt.Println("--------------------------------------------")
}

func decodeFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	input = bytes.TrimSpace(input)
	output, err := Chunk32Decode(string(input))
	if err != nil {
		panic(err)
	}
	// TODO: add a flag to print hex-encoded
	fmt.Println(string(output))
	// fmt.Println(hex.EncodeToString(output))
}
