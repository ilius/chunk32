// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"strings"
)

var CHUNK_SIZE = 10 // should be a factor of 5

func addTrailingDashes(line string, size int) string {
	n := len(line)
	m := ((n-1)/size + 1) * size
	if m > n {
		line = line + strings.Repeat("-", m-n)
	}
	return line
}

func addSpacesToPlainTextLine(line string, size int) string {
	return strings.Join(splitStringIntoChunks(line, size), " ")
}

func encodeChunk(chunk []byte) string {
	line := Encode(chunk)
	line = addTrailingDashes(line, 4)
	line = addSpacesToPlainTextLine(line, 4)
	return line
}

func encodeChunkWithCheck(chunk []byte) string {
	line := CheckEncode(chunk)
	n := len(line)
	line, lastByte := line[:n-1], line[n-1]
	line = addTrailingDashes(line, 4)
	line = addSpacesToPlainTextLine(line, 4)
	return line + " " + string(lastByte)
}

func Chunk32Encode(data []byte, check bool) string {
	chunks := splitBytesIntoChunks(data, CHUNK_SIZE)
	// TODO: add zeros and random bytes to the last (smaller) chunk
	lines := make([]string, len(chunks))
	for i, chunk := range chunks {
		if check {
			lines[i] = encodeChunkWithCheck(chunk)
		} else {
			lines[i] = encodeChunk(chunk)
		}
	}
	return strings.Join(lines, "\n")
}

// `line` must not have spaces
func decodeChunk(line string, allowCheck bool) (bool, []byte, error) {
	if allowCheck && len(line)%4 == 1 {
		line = strings.Replace(line, "-", "", -1)
		chunk, err := CheckDecode(line)
		return true, chunk, err
	} else {
		line = strings.Replace(line, "-", "", -1)
		chunk, err := Decode(line)
		return false, chunk, err
	}
}

func Chunk32Decode(text string) ([]byte, error) {
	lines := strings.Split(text, "\n")
	chunks := make([][]byte, 0, len(lines))
	allowCheck := true
	for lineIndex, line := range lines {
		line = strings.Replace(line, " ", "", -1)
		hadCheck, chunk, err := decodeChunk(line, allowCheck)
		if err != nil {
			return nil, err
		}
		if lineIndex == 0 {
			allowCheck = hadCheck
		}
		log.Println(line, "->", hex.EncodeToString(chunk))
		chunks = append(chunks, chunk)
	}
	return bytes.Join(chunks, []byte{}), nil
}
