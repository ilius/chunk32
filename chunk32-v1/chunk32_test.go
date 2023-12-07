package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var testFileNames = []string{
	"test1",
	"test2",
	"test3",
	"test4",
	"test5",
}

func gzipString(s string) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write([]byte(s))
	if err != nil {
		log.Println(err)
	}
	w.Close()
	return b.Bytes()
}

func gunzipBytes(b []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func testChunk32(t *testing.T, withCheck bool) {
	for _, fname := range testFileNames {
		text := string(MustAsset(fname))
		textGz := gzipString(text)
		textChunk32 := Chunk32Encode(textGz, withCheck)
		textGz2, err := Chunk32Decode(textChunk32)
		if err != nil {
			t.Fatalf("err=%v, fname=%v, withCheck=%v", err, fname, withCheck)
		}
		text2Bytes, err := gunzipBytes(textGz2)
		if err != nil {
			t.Fatalf("err=%v, fname=%v, withCheck=%v", err, fname, withCheck)
		}
		text2 := string(text2Bytes)
		if text2 != text {
			tmpFilePath := filepath.Join(os.TempDir(), "chunk32-"+fname)
			err := ioutil.WriteFile(tmpFilePath, text2Bytes, 0o777)
			if err != nil {
				t.Log(err)
			}
			t.Fatalf("Final text does not match, run: diff 'assets-test/%v' '%v'", fname, tmpFilePath)
		}
	}
}

func TestChunk32_WithCheck(t *testing.T) {
	testChunk32(t, true)
}

func TestChunk32_NoCheck(t *testing.T) {
	testChunk32(t, false)
}
