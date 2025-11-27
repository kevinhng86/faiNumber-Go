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
import "strconv"

/* Note: 
 *
 * BenchSink declaration is save in fnDecimalUtil_test.go file.
 * 
 **/

/* Testing Struct Declaration Start */
type int32ToStringTest struct {
	in  int32
	out string
}

type uint32ToStringTest struct {
	in  uint32
	out string
}
/* Testing Struct Declaration End */

/* Test Cases Start */
var int32ToDecTests = []int32ToStringTest{
	{0 ,"0"},
	{-2147483648 ,"-2147483648"},
	{-2147483647 ,"-2147483647"},
	{2147483647 ,"2147483647"},
	{9 ,"9"},
	{76 ,"76"},
	{100 ,"100"},
	{125 ,"125"},
	{-9 ,"-9"},
	{-76 ,"-76"},
	{-100 ,"-100"},
	{-125 ,"-125"},
}

var uint32ToDecTests = []uint32ToStringTest{
	{0 ,"0"},
	{4294967295 ,"4294967295"},
	{9 ,"9"},
	{76 ,"76"},
	{100 ,"100"},
	{125 ,"125"},
}

var int32ToBinaryTests = []int32ToStringTest{
	{0 ,"0"},
	{-2147483648 ,"-10000000000000000000000000000000"},
	{-2147483647 ,"-1111111111111111111111111111111"},
	{-2147221503 ,"-1111111111110111111111111111111"},
	{2147483647 ,"1111111111111111111111111111111"},
	{2147221503 ,"1111111111110111111111111111111"},
}

var uint32ToBinaryTests = []uint32ToStringTest{
	{0 ,"0"},
	{4294967295 ,"11111111111111111111111111111111"},
	{4294967294 ,"11111111111111111111111111111110"},
	{4294705151 ,"11111111111110111111111111111111"},
}

var int32ToBinaryAsUnsignedTests = []int32ToStringTest{
	{0 ,"0"},
	{-1 ,"11111111111111111111111111111111"},
	{-2 ,"11111111111111111111111111111110"},
}

var int32ToOctalTests = []int32ToStringTest{
	{0 ,"0"},
	{-2147483648 ,"-20000000000"},
	{-2147483647 ,"-17777777777"},
	{2147483647 ,"17777777777"},
}

var uint32ToOctalTests = []uint32ToStringTest{
	{0 ,"0"},
	{4294967295 ,"37777777777"},
}

var int32ToHexTests = []int32ToStringTest{
	{0 ,"0"},
	{-2147483648 ,"-80000000"},
	{-2147483647 ,"-7FFFFFFF"},
	{2147483647 ,"7FFFFFFF"},
}

var uint32ToHexTests = []uint32ToStringTest{
	{0 ,"0"},
	{4294967295 ,"FFFFFFFF"},
}
/* Test Cases End */

/* Test Functions Start */
func TestInt32ToDec(t *testing.T) {
	for i := range int32ToDecTests {
		test := &int32ToDecTests[i]
		out := Int32ToDec(test.in)
		if test.out != out {
			t.Errorf("Int32ToDec(%d) = %v, want %v",
				test.in, []byte(out), []byte(test.out))
		}
	}
}

func TestUInt32ToDec(t *testing.T) {
	for i := range uint32ToDecTests {
		test := &uint32ToDecTests[i]
		out := UInt32ToDec(test.in)
		if test.out != out {
			t.Errorf("UInt32ToDec(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt32ToBinary(t *testing.T) {
	for i := range int32ToBinaryTests {
		test := &int32ToBinaryTests[i]
		out := Int32ToBinary(test.in)
		if test.out != out {
			t.Errorf("Int32ToBinary(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt32ToBinary(t *testing.T) {
	for i := range uint32ToBinaryTests {
		test := &uint32ToBinaryTests[i]
		out := UInt32ToBinary(test.in)
		if test.out != out {
			t.Errorf("UInt32ToBinary(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt32ToBinaryAsUnsigned(t *testing.T) {
	for i := range int32ToBinaryAsUnsignedTests {
		test := &int32ToBinaryAsUnsignedTests[i]
		out := Int32ToBinaryAsUnsigned(test.in)
		if test.out != out {
			t.Errorf("Int32ToBinaryAsUnsigned(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt32ToOctal(t *testing.T) {
	for i := range int32ToOctalTests {
		test := &int32ToOctalTests[i]
		out := Int32ToOctal(test.in)
		if test.out != out {
			t.Errorf("Int32ToOctal(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt32ToOctal(t *testing.T) {
	for i := range uint32ToOctalTests {
		test := &uint32ToOctalTests[i]
		out := UInt32ToOctal(test.in)
		if test.out != out {
			t.Errorf("UInt32ToOctal(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestInt32ToHex(t *testing.T) {
	for i := range int32ToHexTests {
		test := &int32ToHexTests[i]
		out := Int32ToHex(test.in)
		if test.out != out {
			t.Errorf("Int32ToHex(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}

func TestUInt32ToHex(t *testing.T) {
	for i := range uint32ToHexTests {
		test := &uint32ToHexTests[i]
		out := UInt32ToHex(test.in)
		if test.out != out {
			t.Errorf("UInt32ToHex(%d) = %v, want %v",
				test.in, out, test.out)
		}
	}
}
/* Test Functions End */

/* strconv Benchmark Functions Start */
func benchmarkFormatInt(b *testing.B, name string, input int64, base int) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(strconv.FormatInt(input, base))
		}
	})
}

func benchmarkFormatUint(b *testing.B, name string, input uint64, base int) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(strconv.FormatUint(input, base))
		}
	})
}
/* strconv Benchmark Functions End */

/* Benchmark Functions Start*/
func BenchmarkInt32ToDecVSstrconv(b *testing.B) {
	var case1 int32 = 2147483647
	var case1str string = "2147483647"
	var case2 int32 = -2147483647
	var case2str string = "-2147483647"
	b.Run("Int32ToDec", func(b *testing.B) {
		benchmarkInt32ToDec(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, int64(case1), 10)
	})
	b.Run("Int32ToDec", func(b *testing.B) {
		benchmarkInt32ToDec(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, int64(case2), 10)
	})
}

func benchmarkInt32ToDec(b *testing.B, name string, input int32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int32ToDec(input))
		}
	})
}

func BenchmarkUInt32ToDecVSstrconv(b *testing.B) {
	var case1 uint32 = 4294967295
	var case1str string = "4294967295"
	b.Run("UInt32ToDec", func(b *testing.B) {
		benchmarkUInt32ToDec(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, uint64(case1), 10)
	})
}

func benchmarkUInt32ToDec(b *testing.B, name string, input uint32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt32ToDec(input))
		}
	})
}

func BenchmarkInt32ToBinaryVSstrconv(b *testing.B) {
	var case1 int32 = 2147483647
	var case1str string = "2147483647"
	var case2 int32 = -2147483647
	var case2str string = "-2147483647"
	b.Run("Int32ToBinary", func(b *testing.B) {
		benchmarkInt32ToBinary(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, int64(case1), 2)
	})
	b.Run("Int32ToBinary", func(b *testing.B) {
		benchmarkInt32ToBinary(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, int64(case2), 2)
	})
}

func benchmarkInt32ToBinary(b *testing.B, name string, input int32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int32ToBinary(input))
		}
	})
}

func BenchmarkUInt32ToBinaryVSstrconv(b *testing.B) {
	var case1 uint32 = 4294967295
	var case1str string = "4294967295"
	b.Run("UInt32ToBinary", func(b *testing.B) {
		benchmarkUInt32ToBinary(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, uint64(case1), 2)
	})
}

func benchmarkUInt32ToBinary(b *testing.B, name string, input uint32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt32ToBinary(input))
		}
	})
}

func BenchmarkInt32ToOctalVSstrconv(b *testing.B) {
	var case1 int32 = 2147483647
	var case1str string = "2147483647"
	var case2 int32 = -2147483647
	var case2str string = "-2147483647"
	b.Run("Int32ToOctal", func(b *testing.B) {
		benchmarkInt32ToOctal(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, int64(case1), 8)
	})
	b.Run("Int32ToOctal", func(b *testing.B) {
		benchmarkInt32ToOctal(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, int64(case2), 8)
	})
}

func benchmarkInt32ToOctal(b *testing.B, name string, input int32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int32ToOctal(input))
		}
	})
}

func BenchmarkUInt32ToOctalVSstrconv(b *testing.B) {
	var case1 uint32 = 4294967295
	var case1str string = "4294967295"
	b.Run("UInt32ToOctal", func(b *testing.B) {
		benchmarkUInt32ToOctal(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, uint64(case1), 8)
	})
}

func benchmarkUInt32ToOctal(b *testing.B, name string, input uint32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt32ToOctal(input))
		}
	})
}

func BenchmarkInt32ToHexVSstrconv(b *testing.B) {
	var case1 int32 = 2147483647
	var case1str string = "2147483647"
	var case2 int32 = -2147483647
	var case2str string = "-2147483647"
	b.Run("Int32ToHex", func(b *testing.B) {
		benchmarkInt32ToHex(b, case1str, case1)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case1str, int64(case1), 16)
	})
	b.Run("Int32ToHex", func(b *testing.B) {
		benchmarkInt32ToHex(b, case2str, case2)
	})
	b.Run("FormatInt", func(b *testing.B) {
		benchmarkFormatInt(b, case2str, int64(case2), 16)
	})
}

func benchmarkInt32ToHex(b *testing.B, name string, input int32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(Int32ToHex(input))
		}
	})
}

func BenchmarkUInt32ToHexVSstrconv(b *testing.B) {
	var case1 uint32 = 4294967295
	var case1str string = "4294967295"
	b.Run("UInt32ToHex", func(b *testing.B) {
		benchmarkUInt32ToHex(b, case1str, case1)
	})
	b.Run("FormatUint", func(b *testing.B) {
		benchmarkFormatUint(b, case1str, uint64(case1), 16)
	})
}

func benchmarkUInt32ToHex(b *testing.B, name string, input uint32) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BenchSink = len(UInt32ToHex(input))
		}
	})
}
/* Benchmark Functions End*/

