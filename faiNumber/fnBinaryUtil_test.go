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
 * Test structs declaration for string to int is save in fnDecimalUtil_test.go file.
 * BenchSink declaration is save in fnDecimalUtil_test.go file.
 * Benchmark functions for strconv.ParseInt and strconv.ParseUint is save in fnDecimalUtil_test.go file.
 * 
 **/

/* Test Cases Start */
var binaryToInt32NoErrorTests = []noErrorInt32Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-00010000000000000000000000000000000", -2147483648},
	{"-10000000000000000000000000000000", -2147483648},
	{"-1111111111111111111111111111111", -2147483647},
	{"-1111111111111111111111111111110", -2147483646},
	{"0001111111111111111111111111111111", 2147483647},
	{"1111111111111111111111111111111", 2147483647},
	{"1111111111111111111111111111110", 2147483646},
	{"+0001111111111111111111111111111111", 2147483647},
	{"+1111111111111111111111111111111", 2147483647},
	{"+1111111111111111111111111111110", 2147483646},
}

var binaryToInt32ErrorTests = []errorInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToInt32"},
	{"+", 0, 2, "faiNumber.BinaryToInt32"},
	{"-1000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"-11a00000000000000000000000000000", 0, 2, "faiNumber.BinaryToInt32"},
	{"-10000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"-100000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"1000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"100000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"+1000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	{"+100000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt32"},
	// Underflow Errors
	{"-10000000000000000000000000000001", 0, 3, "faiNumber.BinaryToInt32"},
	{"-100000000000000000000000000000001", 0, 3, "faiNumber.BinaryToInt32"},
	// Overflow Errors
	{"10000000000000000000000000000000", 0, 4, "faiNumber.BinaryToInt32"},
	{"+10000000000000000000000000000000", 0, 4, "faiNumber.BinaryToInt32"},
}

var binaryToUInt32NoErrorTests = []noErrorUInt32Test{
	{"000", 0},
	{"00011111111111111111111111111111111", 4294967295},
	{"11111111111111111111111111111111", 4294967295},
	{"11111111111111111111111111111110", 4294967294},
}

var binaryToUInt32ErrorTests = []errorUInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToUInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToUInt32"},
	{"+", 0, 2, "faiNumber.BinaryToUInt32"},
	{"11111111111111111111111111111111a", 0, 2, "faiNumber.BinaryToUInt32"},
	{"1111111111111111111111111111111a", 0, 2, "faiNumber.BinaryToUInt32"},
	// Overflow Errors
	{"111111111111111111111111111111110", 0, 4, "faiNumber.BinaryToUInt32"},
}

var binaryToInt32AsUnsignedNoErrorTests = []noErrorInt32Test{
	{"000", 0},
	{"00011111111111111111111111111111111", -1},
	{"11111111111111111111111111111111", -1},
	{"11111111111111111111111111111110", -2},
}

var binaryToInt32AsUnsignedErrorTests = []errorInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToInt32AsUnsigned"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToInt32AsUnsigned"},
	{"+", 0, 2, "faiNumber.BinaryToInt32AsUnsigned"},
	{"11111111111111111111111111111111a", 0, 2, "faiNumber.BinaryToInt32AsUnsigned"},
	{"1111111111111111111111111111111a", 0, 2, "faiNumber.BinaryToInt32AsUnsigned"},
	// Overflow Errors
	{"111111111111111111111111111111110", 0, 4, "faiNumber.BinaryToInt32AsUnsigned"},
}

var binaryToInt64NoErrorTests = []noErrorInt64Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-0001000000000000000000000000000000000000000000000000000000000000000", -9223372036854775808},
	{"-1000000000000000000000000000000000000000000000000000000000000000", -9223372036854775808},
	{"-111111111111111111111111111111111111111111111111111111111111111", -9223372036854775807},
	{"-111111111111111111111111111111111111111111111111111111111111110", -9223372036854775806},
	{"000111111111111111111111111111111111111111111111111111111111111111", 9223372036854775807},
	{"111111111111111111111111111111111111111111111111111111111111111", 9223372036854775807},
	{"111111111111111111111111111111111111111111111111111111111111110", 9223372036854775806},
	{"+000111111111111111111111111111111111111111111111111111111111111111", 9223372036854775807},
	{"+111111111111111111111111111111111111111111111111111111111111111", 9223372036854775807},
	{"+111111111111111111111111111111111111111111111111111111111111110", 9223372036854775806},
}

var binaryToInt64ErrorTests = []errorInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToInt64"},
	{"+", 0, 2, "faiNumber.BinaryToInt64"},
	{"-100000000000000000000000000000000000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt64"},
	{"-11a0000000000000000000000000000000000000000000000000000000000000", 0, 2, "faiNumber.BinaryToInt64"},
	{"-1000000000000000000000000000000000000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt64"},
	{"100000000000000000000000000000000000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt64"},
	{"+100000000000000000000000000000000000000000000000000000000000000a", 0, 2, "faiNumber.BinaryToInt64"},
	// Underflow Errors
	{"-1000000000000000000000000000000000000000000000000000000000000001", 0, 3, "faiNumber.BinaryToInt64"},
	{"-10000000000000000000000000000000000000000000000000000000000000001", 0, 3, "faiNumber.BinaryToInt64"},
	// Overflow Errors
	{"1000000000000000000000000000000000000000000000000000000000000000", 0, 4, "faiNumber.BinaryToInt64"},
	{"+1000000000000000000000000000000000000000000000000000000000000000", 0, 4, "faiNumber.BinaryToInt64"},
}

var binaryToUInt64NoErrorTests = []noErrorUInt64Test{
	{"000", 0},
	{"0001111111111111111111111111111111111111111111111111111111111111111", 18446744073709551615},
	{"1111111111111111111111111111111111111111111111111111111111111111", 18446744073709551615},
	{"1111111111111111111111111111111111111111111111111111111111111110", 18446744073709551614},
}

var binaryToUInt64ErrorTests = []errorUInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToUInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToUInt64"},
	{"+", 0, 2, "faiNumber.BinaryToUInt64"},
	{"11111111111111111111111111111111111111111111111111111111111111112", 0, 2, "faiNumber.BinaryToUInt64"},
	{"1111111111111111111111111111111111111111111111111111111111111112", 0, 2, "faiNumber.BinaryToUInt64"},
	// Overflow Errors
	{"11111111111111111111111111111111111111111111111111111111111111111", 0, 4, "faiNumber.BinaryToUInt64"},
}

var binaryToInt64AsUnsignedNoErrorTests = []noErrorInt64Test{
	{"000", 0},
	{"0001111111111111111111111111111111111111111111111111111111111111111", -1},
	{"1111111111111111111111111111111111111111111111111111111111111111", -1},
	{"1111111111111111111111111111111111111111111111111111111111111110", -2},
}

var binaryToInt64AsUnsignedErrorTests = []errorInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.BinaryToInt64AsUnsigned"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.BinaryToInt64AsUnsigned"},
	{"+", 0, 2, "faiNumber.BinaryToInt64AsUnsigned"},
	{"11111111111111111111111111111111111111111111111111111111111111112", 0, 2, "faiNumber.BinaryToInt64AsUnsigned"},
	{"1111111111111111111111111111111111111111111111111111111111111112", 0, 2, "faiNumber.BinaryToInt64AsUnsigned"},
	// Overflow Errors
	{"11111111111111111111111111111111111111111111111111111111111111110", 0, 4, "faiNumber.BinaryToInt64AsUnsigned"},
}
/* Test Cases End */

/* Test Functions Start */
func TestBinaryToInt32NoError(t *testing.T) {
	for i := range binaryToInt32NoErrorTests {
		test := &binaryToInt32NoErrorTests[i]
		out, err := BinaryToInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToInt32Error(t *testing.T) {
	for i := range binaryToInt32ErrorTests {
		test := &binaryToInt32ErrorTests[i]
		out, err := BinaryToInt32(test.in)
		if err == nil {
			t.Errorf("BinaryToInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestBinaryToUInt32NoError(t *testing.T) {
	for i := range binaryToUInt32NoErrorTests {
		test := &binaryToUInt32NoErrorTests[i]
		out, err := BinaryToUInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToUInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToUInt32Error(t *testing.T) {
	for i := range binaryToUInt32ErrorTests {
		test := &binaryToUInt32ErrorTests[i]
		out, err := BinaryToUInt32(test.in)
		if err == nil {
			t.Errorf("BinaryToUInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToUInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestBinaryToInt32AsUnsignedNoError(t *testing.T) {
	for i := range binaryToInt32AsUnsignedNoErrorTests {
		test := &binaryToInt32AsUnsignedNoErrorTests[i]
		out, err := BinaryToInt32AsUnsigned(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToInt32AsUnsigned(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToInt32AsUnsignedError(t *testing.T) {
	for i := range binaryToInt32AsUnsignedErrorTests {
		test := &binaryToInt32AsUnsignedErrorTests[i]
		out, err := BinaryToInt32AsUnsigned(test.in)
		if err == nil {
			t.Errorf("BinaryToInt32AsUnsigned(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToInt32AsUnsigned(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestBinaryToInt64NoError(t *testing.T) {
	for i := range binaryToInt64NoErrorTests {
		test := &binaryToInt64NoErrorTests[i]
		out, err := BinaryToInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToInt64Error(t *testing.T) {
	for i := range binaryToInt64ErrorTests {
		test := &binaryToInt64ErrorTests[i]
		out, err := BinaryToInt64(test.in)
		if err == nil {
			t.Errorf("BinaryToInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestBinaryToUInt64NoError(t *testing.T) {
	for i := range binaryToUInt64NoErrorTests {
		test := &binaryToUInt64NoErrorTests[i]
		out, err := BinaryToUInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToUInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToUInt64Error(t *testing.T) {
	for i := range binaryToUInt64ErrorTests {
		test := &binaryToUInt64ErrorTests[i]
		out, err := BinaryToUInt64(test.in)
		if err == nil {
			t.Errorf("BinaryToUInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToUInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestBinaryToInt64AsUnsignedNoError(t *testing.T) {
	for i := range binaryToInt64AsUnsignedNoErrorTests {
		test := &binaryToInt64AsUnsignedNoErrorTests[i]
		out, err := BinaryToInt64AsUnsigned(test.in)
		if test.out != out || err != nil {
			t.Errorf("BinaryToInt64AsUnsigned(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestBinaryToInt64AsUnsignedError(t *testing.T) {
	for i := range binaryToInt64AsUnsignedErrorTests {
		test := &binaryToInt64AsUnsignedErrorTests[i]
		out, err := BinaryToInt64AsUnsigned(test.in)
		if err == nil {
			t.Errorf("BinaryToInt64AsUnsigned(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("BinaryToInt64AsUnsigned(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}
/* Test Functions End */

/* Benchmark Functions Start*/
func BenchmarkBinaryToInt32VSstrconv(b *testing.B) {
	case1 := "10101110101111010101"
	case2 := "-10101110101111010101"
	b.Run("BinaryToInt32", func(b *testing.B) {
		benchmarkBinaryToInt32(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 2, 32)
	})
	b.Run("BinaryToInt32", func(b *testing.B) {
		benchmarkBinaryToInt32(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 2, 32)
	})
}

func benchmarkBinaryToInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := BinaryToInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkBinaryToUInt32VSstrconv(b *testing.B) {
	case1 := "101010101010101001110"
	b.Run("BinaryToUInt32", func(b *testing.B) {
		benchmarkBinaryToUInt32(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 2, 32)
	})
}

func benchmarkBinaryToUInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := BinaryToUInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkBinaryToInt64VSstrconv(b *testing.B) {
	case1 := "111111000000000000000111111110"
	case2 := "-111111000000000000000111111110"
	b.Run("BinaryToInt64", func(b *testing.B) {
		benchmarkBinaryToInt64(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 2, 64)
	})
	b.Run("BinaryToInt64", func(b *testing.B) {
		benchmarkBinaryToInt64(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 2, 64)
	})
}

func benchmarkBinaryToInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := BinaryToInt64(input)
			BenchSink = out
		}
	})
}

func BenchmarkBinaryToUInt64VSstrconv(b *testing.B) {
	case1 := "10101010101010101010101010"
	b.Run("BinaryToUInt64", func(b *testing.B) {
		benchmarkBinaryToUInt64(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 2, 64)
	})
}

func benchmarkBinaryToUInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := BinaryToUInt64(input)
			BenchSink = out
		}
	})
}
/* Benchmark Functions End*/
