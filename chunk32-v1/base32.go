// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
// based on https://github.com/btcsuite/btcd/tree/master/btcutil/base58

package main

import (
	"fmt"
	"math/big"
)

//go:generate go run genalphabet.go

var (
	bigRadix = big.NewInt(32)
	bigZero  = big.NewInt(0)
)

func decode(s string) ([]byte, *big.Int, error) {
	bigVal := big.NewInt(0)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(s) - 1; i >= 0; i-- {
		tmp := b32[s[i]]
		if tmp == 255 {
			return nil, nil, fmt.Errorf("error in Decode: invalid character %#v", string(s[i]))
		}
		scratch.SetInt64(int64(tmp))
		scratch.Mul(j, scratch)
		bigVal.Add(bigVal, scratch)
		j.Mul(j, bigRadix)
	}

	tmpBytes := bigVal.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(s); numZeros++ {
		if s[numZeros] != alphabetIdx0 {
			break
		}
	}
	flen := numZeros + len(tmpBytes)
	output := make([]byte, flen)
	copy(output[numZeros:], tmpBytes)

	return output, bigVal, nil
}

// Decode decodes a modified base32 string to a byte slice.
func Decode(b string) ([]byte, error) {
	output, _, err := decode(b)
	return output, err
}

// encode and return check number (0 to 36)
func encode(b []byte, wantCheck bool) (string, uint8) {
	x := new(big.Int)
	x.SetBytes(b)

	mod := uint8(0)
	if wantCheck {
		bigMod := new(big.Int)
		bigMod.Mod(x, big.NewInt(37))
		mod = uint8(bigMod.Int64())
	}

	output := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		output = append(output, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		output = append(output, alphabetIdx0)
	}

	// reverse
	alen := len(output)
	for i := 0; i < alen/2; i++ {
		output[i], output[alen-1-i] = output[alen-1-i], output[i]
	}

	return string(output), mod
}

// Encode encodes a byte slice to a modified base32 string.
func Encode(b []byte) string {
	output, _ := encode(b, false)
	return output
}
