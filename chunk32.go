// Copyright (c) 2017 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"strings"

	"github.com/ilius/crock32"
)

var CHUNK_SIZE = 10 // should be a factor of 5

func addSpacesToPlainTextLine(line string, size int) string {
	return strings.Join(splitStringIntoChunks(line, 4), " ")
}

func Chunk32Encode(data []byte) string {
	chunks := splitBytesIntoChunks(data, CHUNK_SIZE)
	// TODO: add zeros and random bytes to the last (smaller) chunk
	lines := make([]string, len(chunks))
	for i, chunk := range chunks {
		line := crock32.Encode(chunk)
		line = addSpacesToPlainTextLine(line, 4)
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

func Chunk32Decode(text string) ([]byte, error) {
	lines := strings.Split(text, "\n")
	chunks := make([][]byte, 0, len(lines))
	for _, line := range lines {
		line = strings.Replace(line, " ", "", -1)
		chunk, err := crock32.Decode(line)
		if err != nil {
			return nil, err
		}
		log.Println(line, "->", hex.EncodeToString(chunk))
		chunks = append(chunks, chunk)
	}
	return bytes.Join(chunks, []byte{}), nil
}
