// Copyright (C) 2015 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package colorhash

import "testing"

var data = []byte("Correct Horse Battery Staple")

func TestBytesAsANSI(t *testing.T) {
	var (
		expected = []uint8{191, 35, 175, 178, 148, 70, 82, 37}
		got      = BytesAsANSI(data)
	)

	for i, el := range got {
		if el != expected[i] {
			t.Errorf("BytesAsANSI(%v) == %v, want %v", data, got, expected)
		}
	}
}

func TestStringAsANSI(t *testing.T) {
	var (
		expected = []uint8{191, 35, 175, 178, 148, 70, 82, 37}
		got      = StringAsANSI(string(data))
	)

	for i, el := range got {
		if el != expected[i] {
			t.Errorf("BytesAsANSI(%v) == %v, want %v", data, got, expected)
		}
	}
}

func TestBytesAsRGB(t *testing.T) {
	var (
		expected = [][]uint8{
			[]uint8{204, 255, 51},
			[]uint8{0, 153, 51},
			[]uint8{204, 102, 153},
			[]uint8{204, 153, 0},
			[]uint8{153, 204, 0},
			[]uint8{51, 153, 0},
			[]uint8{51, 255, 0},
			[]uint8{0, 153, 153},
		}
		got = BytesAsRGB(data)
	)

	for i, part := range got {
		for j, el := range part {
			if el != expected[i][j] {
				t.Errorf("BytesAsANSI(%v) == %v, want %v", data, got, expected)
			}
		}
	}
}

func TestStringAsRGB(t *testing.T) {
	var (
		expected = [][]uint8{
			[]uint8{204, 255, 51},
			[]uint8{0, 153, 51},
			[]uint8{204, 102, 153},
			[]uint8{204, 153, 0},
			[]uint8{153, 204, 0},
			[]uint8{51, 153, 0},
			[]uint8{51, 255, 0},
			[]uint8{0, 153, 153},
		}
		got = StringAsRGB(string(data))
	)

	for i, part := range got {
		for j, el := range part {
			if el != expected[i][j] {
				t.Errorf("BytesAsANSI(%v) == %v, want %v", data, got, expected)
			}
		}
	}
}

func BenchmarkBytesAsANSI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesAsANSI(data)
	}
}

func BenchmarkBytesAsRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesAsRGB(data)
	}
}

func BenchmarkStringAsANSI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAsANSI(string(data))
	}
}

func BenchmarkStringAsRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAsRGB(string(data))
	}
}
