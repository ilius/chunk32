// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>

package main

func splitBytesIntoChunks(data []byte, size int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(data)/size+1)
	for len(data) >= size {
		chunk, data = data[:size], data[size:]
		chunks = append(chunks, chunk)
	}
	if len(data) > 0 {
		chunks = append(chunks, data[:])
	}
	return chunks
}

func splitStringIntoChunks(data string, size int) []string {
	var chunk string
	chunks := make([]string, 0, len(data)/size+1)
	for len(data) >= size {
		chunk, data = data[:size], data[size:]
		chunks = append(chunks, chunk)
	}
	if len(data) > 0 {
		chunks = append(chunks, data[:])
	}
	return chunks
}
