// Copyright (C) 2015 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package colorhash

import "crypto/sha256"

func BytesAsANSI(data []byte) []uint8 {
	var (
		hash   = sha256.Sum256(data)
		colors = make([]uint8, 8)
	)

	for i := 0; i < 8; i++ {
		colors[i] = 16 + uint8(hash[i])%216
	}

	return colors
}

func StringAsANSI(data string) []uint8 {
	return BytesAsANSI([]byte(data))
}

// BytesAsRGB accepts slice of bytes and returns slice of slices of bytes, each slice of bytes stands for three 0-255 RGB values
func BytesAsRGB(data []byte) [][]uint8 {
	var (
		hash   = BytesAsANSI(data)
		colors = make([][]uint8, 8)
		x      uint8
	)

	for i := 0; i < 8; i++ {
		x = hash[i] - 16

		colors[i] = []uint8{
			uint8(x / 36 * 51),
			uint8(x % 36 / 6 * 51),
			uint8(x % 36 % 6 * 51),
		}
	}

	return colors
}

// StringAsRGB accepts string and returns slice of slices of bytes, each slice of bytes stands for three 0-255 RGB values
func StringAsRGB(data string) [][]uint8 {
	return BytesAsRGB([]byte(data))
}
