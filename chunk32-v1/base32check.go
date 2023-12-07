// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2017 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/big"
)

var checkAlphabet = "*~$=U"

// // ErrChecksum indicates that the checksum of a check-encoded string does not verify against
// // the checksum.
// var ErrChecksum = errors.New("checksum error")

// // ErrInvalidFormat indicates that the check-encoded string has an invalid format.
// var ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")

func getCheckByte(mod uint8) byte {
	if mod < 32 {
		return alphabet[int(mod)]
	}
	return checkAlphabet[int(mod-32)]
}

func getCheckValue(b byte) uint8 {
	switch b {
	case '*':
		return 32
	case '~':
		return 33
	case '$':
		return 34
	case '=':
		return 35
	case 'u', 'U':
		return 36
	default:
		return b32[b]
	}
}

func CheckEncode(b []byte) string {
	answer, mod := encode(b, true)
	if mod >= 37 {
		panic(fmt.Sprintf("invalid check mod=%v", mod))
	}
	checkStr := string(getCheckByte(mod))
	return answer + checkStr
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the check byte
func CheckDecode(s string) ([]byte, error) {
	n := len(s)
	s, lastByte := s[:n-1], s[n-1]

	output, bigVal, err := decode(s)
	if err != nil {
		return output, err
	}

	lastByteVal := getCheckValue(lastByte)
	if lastByteVal == 255 {
		return output, fmt.Errorf("invalid check character %#v", string(lastByte))
	}

	bigMod := new(big.Int)
	bigMod.Mod(bigVal, big.NewInt(37))
	bidMod64 := bigMod.Int64()
	if bidMod64 != int64(lastByteVal) {
		return output, fmt.Errorf("check failed, %v != %v (%v)", bidMod64, lastByteVal, string(lastByte))
	}

	return output, err
}
