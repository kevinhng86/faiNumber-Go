/*
###########################################################################################
#                                                                                         #
#    Copyright 2025 Khang Hoang Nguyen (kevinhg86)                                        #
#    E-mail: kevinhng86@gmail.com | Web: https://fai.host                                 #
#                                                                                         #
#    Permission is hereby granted, free of charge, to any person obtaining a copy         #
#    of this software and associated documentation files (the "Software"), to deal        #
#    in the Software without restriction, including without limitation the rights         #
#    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies     #
#    of the Software, and to permit persons to whom the Software is furnished             #
#    to do so, subject to the following conditions:                                       #
#                                                                                         #
#    The above copyright notice and this permission notice shall be  included in all      #
#    copies or substantial portions of the Software.                                      #
#                                                                                         #
#    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR           #
#    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,             #
#    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL              #
#    THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER           #
#    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,        #
#    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN            #
#    THE SOFTWARE.                                                                        #
#                                                                                         #
###########################################################################################
*/

package faiNumber

import "testing"

/* Note: 
 *
 * BenchSink declaration is save in fnDecimalUtil_test.go file.
 * Benchmark functions for strconv.FormatInt and strconv.FormatUint is save in fnInt32Util_test.go file.
 * 
 **/

/* Testing Struct Declaration Start */
type int64ToStringTest struct {
	in  int64
	out string
}

type uint64ToStringTest struct {
	in  uint64
	out string
}
/* Testing Struct Declaration End */

/* Test Cases Start */
var int64ToDecTests = []int64ToStringTest{
	{0 ,"0"},
	{-9223372036854775808 ,"-9223372036854775808"},
	{-9223372036854775807 ,"-9223372036854775807"},
	{9223372036854775807 ,"9223372036854775807"},
	{9 ,"9"},
	{76 ,"76"},
	{100 ,"100"},
	{125 ,"125"},
	{-9 ,"-9"},
	{-76 ,"-76"},
	{-100 ,"-100"},
	{-125 ,"-125"},
}

var uint64ToDecTests = []uint64ToStringTest{
	{0 ,"0"},
	{18446744073709551615 ,"18446744073709551615"},
	{9 ,"9"},
	{76 ,"76"},
	{100 ,"100"},
	{125 ,"125"},
}

var int64ToBinaryTests = []int64ToStringTest{
	{0 ,"0"},
	{-9223372036854775808 ,"-1000000000000000000000000000000000000000000000000000000000000000"},
	{-9223372036854775807 ,"-111111111111111111111111111111111111111111111111111111111111111"},
	{-9223372036854775806 ,"-111111111111111111111111111111111111111111111111111111111111110"},
	{-9223336852482686975 ,"-111111111111111110111111111111111111111111111111111111111111111"},
	{9223372036854775807 ,"111111111111111111111111111111111111111111111111111111111111111"},
	{9223372036854775806 ,"111111111111111111111111111111111111111111111111111111111111110"},
	{9223336852482686975 ,"111111111111111110111111111111111111111111111111111111111111111"},
}

var uint64ToBinaryTests = []uint64ToStringTest{
	{0 ,"0"},
	{18446744073709551615 ,"1111111111111111111111111111111111111111111111111111111111111111"},
	{18446744073709551614 ,"1111111111111111111111111111111111111111111111111111111111111110"},
	{18446744073692774398 ,"1111111111111111111111111111111111111110111111111111111111111110"},
}

var int64ToBinaryAsUnsignedTests = []int64ToStringTest{
	{0 ,"0"},
	{-1 ,"1111111111111111111111111111111111111111111111111111111111111111"},
	{-2 ,"1111111111111111111111111111111111111111111111111111111111111110"},
}

var int64ToOctalTests = []int64ToStringTest{
	{0 ,"0"},
	{-9223372036854775808 ,"-1000000000000000000000"},
	{-9223372036854775807 ,"-777777777777777777777"},
	{9223372036854775807 ,"777777777777777777777"},
}

var uint64ToOctalTests = []uint64ToStringTest{
	{0 ,"0"},
	{18446744073709551615 ,"1777777777777777777777"},
}

var int64ToHexTests = []int64ToStringTest{
	{0 ,"0"},
	{-9223372036854775808 ,"-8000000000000000"},
	{-9223372036854775807 ,"-7FFFFFFFFFFFFFFF"},
	{9223372036854775807 ,"7FFFFFFFFFFFFFFF"},
}

var uint64ToHexTests = []uint64ToStringTest{
	{0 ,"0"},
	{18446744073709551615 ,"FFFFFFFFFFFFFFFF"},
}
/* Test Cases End */


/* Test Functions Start */
func TestInt64ToDec(t *testing.T) {
	for i := range int64ToDecTests {
		test := &int64ToDecTests[i]
		out := Int64ToDec(test.in)
		if test.out != out {
			t.Errorf("Int64ToDec(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt64ToDec(t *testing.T) {
	for i := range uint64ToDecTests {
		test := &uint64ToDecTests[i]
		out := UInt64ToDec(test.in)
		if test.out != out {
			t.Errorf("UInt64ToDec(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt64ToBinary(t *testing.T) {
	for i := range int64ToBinaryTests {
		test := &int64ToBinaryTests[i]
		out := Int64ToBinary(test.in)
		if test.out != out {
			t.Errorf("Int64ToBinary(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt64ToBinary(t *testing.T) {
	for i := range uint64ToBinaryTests {
		test := &uint64ToBinaryTests[i]
		out := UInt64ToBinary(test.in)
		if test.out != out {
			t.Errorf("UInt64ToBinary(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt64ToBinaryAsUnsigned(t *testing.T) {
	for i := range int64ToBinaryAsUnsignedTests {
		test := &int64ToBinaryAsUnsignedTests[i]
		out := Int64ToBinaryAsUnsigned(test.in)
		if test.out != out {
			t.Errorf("Int64ToBinaryAsUnsigned(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt64ToOctal(t *testing.T) {
	for i := range int64ToOctalTests {
		test := &int64ToOctalTests[i]
		out := Int64ToOctal(test.in)
		if test.out != out {
			t.Errorf("Int64ToOctal(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt64ToOctal(t *testing.T) {
	for i := range uint64ToOctalTests {
		test := &uint64ToOctalTests[i]
		out := UInt64ToOctal(test.in)
		if test.out != out {
			t.Errorf("UInt64ToOctal(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt64ToHex(t *testing.T) {
	for i := range int64ToHexTests {
		test := &int64ToHexTests[i]
		out := Int64ToHex(test.in)
		if test.out != out {
			t.Errorf("Int64ToHex(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt64ToHex(t *testing.T) {
	for i := range uint64ToHexTests {
		test := &uint64ToHexTests[i]
		out := UInt64ToHex(test.in)
		if test.out != out {
			t.Errorf("UInt64ToHex(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}
/* Test Functions End */

/* Benchmark Functions Start*/
func BenchmarkInt64ToDecVSstrconv(b *testing.B) {
	var case1 int64 = 9223372036854775807
	var case1str string = "9223372036854775807"
	var case2 int64 = -9223372036854775807
	var case2str string = "-9223372036854775807"
	b.Run("Int64ToDec", func(b *testing.B) {
		benchmarkInt64ToDec(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, case1, 10)
	})
	b.Run("Int64ToDec", func(b *testing.B) {
		benchmarkInt64ToDec(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, case2, 10)
	})
}

func benchmarkInt64ToDec(b *testing.B, name string, input int64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int64ToDec(input))
		}
	})
}

func BenchmarkUInt64ToDecVSstrconv(b *testing.B) {
	var case1 uint64 = 18446744073709551615
	var case1str string = "18446744073709551615"
	b.Run("UInt64ToDec", func(b *testing.B) {
		benchmarkUInt64ToDec(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, case1, 10)
	})
}

func benchmarkUInt64ToDec(b *testing.B, name string, input uint64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt64ToDec(input))
		}
	})
}

func BenchmarkInt64ToBinaryVSstrconv(b *testing.B) {
	var case1 int64 = 9223372036854775807
	var case1str string = "9223372036854775807"
	var case2 int64 = -9223372036854775807
	var case2str string = "-9223372036854775807"
	b.Run("Int64ToBinary", func(b *testing.B) {
		benchmarkInt64ToBinary(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, case1, 2)
	})
	b.Run("Int64ToBinary", func(b *testing.B) {
		benchmarkInt64ToBinary(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, case2, 2)
	})
}

func benchmarkInt64ToBinary(b *testing.B, name string, input int64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int64ToBinary(input))
		}
	})
}

func BenchmarkUInt64ToBinaryVSstrconv(b *testing.B) {
	var case1 uint64 = 18446744073709551615
	var case1str string = "18446744073709551615"
	b.Run("UInt64ToBinary", func(b *testing.B) {
		benchmarkUInt64ToBinary(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, case1, 2)
	})
}

func benchmarkUInt64ToBinary(b *testing.B, name string, input uint64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt64ToBinary(input))
		}
	})
}

func BenchmarkInt64ToOctalVSstrconv(b *testing.B) {
	var case1 int64 = 9223372036854775807
	var case1str string = "9223372036854775807"
	var case2 int64 = -9223372036854775807
	var case2str string = "-9223372036854775807"
	b.Run("Int64ToOctal", func(b *testing.B) {
		benchmarkInt64ToOctal(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, case1, 8)
	})
	b.Run("Int64ToOctal", func(b *testing.B) {
		benchmarkInt64ToOctal(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, case2, 8)
	})
}

func benchmarkInt64ToOctal(b *testing.B, name string, input int64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int64ToOctal(input))
		}
	})
}

func BenchmarkUInt64ToOctalVSstrconv(b *testing.B) {
	var case1 uint64 = 18446744073709551615
	var case1str string = "18446744073709551615"
	b.Run("UInt64ToOctal", func(b *testing.B) {
		benchmarkUInt64ToOctal(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, case1, 8)
	})
}

func benchmarkUInt64ToOctal(b *testing.B, name string, input uint64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt64ToOctal(input))
		}
	})
}

func BenchmarkInt64ToHexVSstrconv(b *testing.B) {
	var case1 int64 = 9223372036854775807
	var case1str string = "9223372036854775807"
	var case2 int64 = -9223372036854775807
	var case2str string = "-9223372036854775807"
	b.Run("Int64ToHex", func(b *testing.B) {
		benchmarkInt64ToHex(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, case1, 16)
	})
	b.Run("Int64ToHex", func(b *testing.B) {
		benchmarkInt64ToHex(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, case2, 16)
	})
}

func benchmarkInt64ToHex(b *testing.B, name string, input int64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int64ToHex(input))
		}
	})
}

func BenchmarkUInt64ToHexVSstrconv(b *testing.B) {
	var case1 uint64 = 18446744073709551615
	var case1str string = "18446744073709551615"
	b.Run("UInt64ToHex", func(b *testing.B) {
		benchmarkUInt64ToHex(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, case1, 16)
	})
}

func benchmarkUInt64ToHex(b *testing.B, name string, input uint64) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt64ToHex(input))
		}
	})
}
/* Benchmark Functions End*/
