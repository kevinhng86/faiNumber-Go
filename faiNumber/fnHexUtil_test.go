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
var hexToInt32NoErrorTests = []noErrorInt32Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-000abcdef", -11259375}, // This is to check whether fnDigitMap working correctly for digit 'a' - 'f'
	{"-abcdef", -11259375}, // This is to check whether fnDigitMap working correctly for digit 'a' - 'f'
	{"-ABCDEF", -11259375}, // This is to check whether fnDigitMap working correctly for digit 'A' - 'F'
	{"-80000000", -2147483648},
	{"-7FFFFFFF", -2147483647},
	{"-7FFFFFF", -134217727},
	{"0007FFFFFFF", 2147483647},
	{"7FFFFFFF", 2147483647},
	{"7FFFFFF", 134217727},
	{"+0007FFFFFFF", 2147483647},
	{"+7FFFFFFF", 2147483647},
	{"+7FFFFFF", 134217727},	
}

var hexToInt32ErrorTests = []errorInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.HexToInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.HexToInt32"},
	{"+", 0, 2, "faiNumber.HexToInt32"},
    /* Checking fnDigitMap Digit Larger Than 'F' or 'f' Start */
	{"G", 0, 2, "faiNumber.HexToInt32"},
	{"H", 0, 2, "faiNumber.HexToInt32"},
	{"I", 0, 2, "faiNumber.HexToInt32"},
	{"J", 0, 2, "faiNumber.HexToInt32"},
	{"K", 0, 2, "faiNumber.HexToInt32"},
	{"L", 0, 2, "faiNumber.HexToInt32"},
	{"M", 0, 2, "faiNumber.HexToInt32"},
	{"N", 0, 2, "faiNumber.HexToInt32"},
	{"O", 0, 2, "faiNumber.HexToInt32"},
	{"P", 0, 2, "faiNumber.HexToInt32"},
	{"Q", 0, 2, "faiNumber.HexToInt32"},
	{"R", 0, 2, "faiNumber.HexToInt32"},
	{"S", 0, 2, "faiNumber.HexToInt32"},
	{"T", 0, 2, "faiNumber.HexToInt32"},
	{"U", 0, 2, "faiNumber.HexToInt32"},
	{"V", 0, 2, "faiNumber.HexToInt32"},
	{"W", 0, 2, "faiNumber.HexToInt32"},
	{"X", 0, 2, "faiNumber.HexToInt32"},
	{"Y", 0, 2, "faiNumber.HexToInt32"},
	{"Z", 0, 2, "faiNumber.HexToInt32"},
	{"g", 0, 2, "faiNumber.HexToInt32"},
	{"h", 0, 2, "faiNumber.HexToInt32"},
	{"i", 0, 2, "faiNumber.HexToInt32"},
	{"j", 0, 2, "faiNumber.HexToInt32"},
	{"k", 0, 2, "faiNumber.HexToInt32"},
	{"l", 0, 2, "faiNumber.HexToInt32"},
	{"m", 0, 2, "faiNumber.HexToInt32"},
	{"n", 0, 2, "faiNumber.HexToInt32"},
	{"o", 0, 2, "faiNumber.HexToInt32"},
	{"p", 0, 2, "faiNumber.HexToInt32"},
	{"q", 0, 2, "faiNumber.HexToInt32"},
	{"r", 0, 2, "faiNumber.HexToInt32"},
	{"s", 0, 2, "faiNumber.HexToInt32"},
	{"t", 0, 2, "faiNumber.HexToInt32"},
	{"u", 0, 2, "faiNumber.HexToInt32"},
	{"v", 0, 2, "faiNumber.HexToInt32"},
	{"w", 0, 2, "faiNumber.HexToInt32"},
	{"x", 0, 2, "faiNumber.HexToInt32"},
	{"y", 0, 2, "faiNumber.HexToInt32"},
	{"z", 0, 2, "faiNumber.HexToInt32"},
    /* Checking fnDigitMap Digit Larger Than 'F' or 'f' End */
	{"-81G00000", 0, 2, "faiNumber.HexToInt32"},
	{"-90G00000", 0, 2, "faiNumber.HexToInt32"},
	{"-7FFFFFFFG", 0, 2, "faiNumber.HexToInt32"},
	{"-7FFFFFG", 0, 2, "faiNumber.HexToInt32"},
	{"8000000G", 0, 2, "faiNumber.HexToInt32"},
	{"7FFFFFFFG", 0, 2, "faiNumber.HexToInt32"},
	{"7FFFFFG", 0, 2, "faiNumber.HexToInt32"},
	{"+8000000G", 0, 2, "faiNumber.HexToInt32"},
	{"+7FFFFFFFG", 0, 2, "faiNumber.HexToInt32"},
	{"+7FFFFFG", 0, 2, "faiNumber.HexToInt32"},
	// Underflow Errors
	{"-80000001", 0, 3, "faiNumber.HexToInt32"},
	{"-90000000", 0, 3, "faiNumber.HexToInt32"},
	{"-7FFFFFFF1", 0, 3, "faiNumber.HexToInt32"},
	// Overflow Errors
	{"80000000", 0, 4, "faiNumber.HexToInt32"},
	{"7FFFFFFF1", 0, 4, "faiNumber.HexToInt32"},
}

var hexToUInt32NoErrorTests = []noErrorUInt32Test{
	{"000", 0},
	{"000FFFFFFFF", 4294967295},
	{"FFFFFFFF", 4294967295},
}

var hexToUInt32ErrorTests = []errorUInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.HexToUInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.HexToUInt32"},
	{"+", 0, 2, "faiNumber.HexToUInt32"},
	{"FFFFFFFg", 0, 2, "faiNumber.HexToUInt32"},
	{"FFFFFFFFg", 0, 2, "faiNumber.HexToUInt32"},
	// Overflow Errors
	{"FFFFFFFF1", 0, 4, "faiNumber.HexToUInt32"},
}

var hexToInt64NoErrorTests = []noErrorInt64Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-0008000000000000000", -9223372036854775808},
	{"-8000000000000000", -9223372036854775808},
	{"-7FFFFFFFFFFFFFFF", -9223372036854775807},
	{"-7FFFFFFFFFFFFFF", -576460752303423487},
	{"0007FFFFFFFFFFFFFFF", 9223372036854775807},
	{"7FFFFFFFFFFFFFFF", 9223372036854775807},
	{"7FFFFFFFFFFFFFF", 576460752303423487},
	{"+0007FFFFFFFFFFFFFFF", 9223372036854775807},
	{"+7FFFFFFFFFFFFFFF", 9223372036854775807},
	{"+7FFFFFFFFFFFFFF", 576460752303423487},
}

var hexToInt64ErrorTests = []errorInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.HexToInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.HexToInt64"},
	{"+", 0, 2, "faiNumber.HexToInt64"},
	{"-81g0000000000000", 0, 2, "faiNumber.HexToInt64"},
	{"-900000000000000g", 0, 2, "faiNumber.HexToInt64"},
	{"-7FFFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	{"-7FFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	{"800000000000000g", 0, 2, "faiNumber.HexToInt64"},
	{"7FFFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	{"7FFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	{"+800000000000000g", 0, 2, "faiNumber.HexToInt64"},
	{"+7FFFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	{"+7FFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToInt64"},
	// Underflow Errors
	{"-8000000000000001", 0, 3, "faiNumber.HexToInt64"},
	{"-80000000000000001", 0, 3, "faiNumber.HexToInt64"},
	{"-9000000000000000", 0, 3, "faiNumber.HexToInt64"},
	// Overflow Errors
	{"8000000000000000", 0, 4, "faiNumber.HexToInt64"},
	{"7FFFFFFFFFFFFFFF1", 0, 4, "faiNumber.HexToInt64"},
}

var hexToUInt64NoErrorTests = []noErrorUInt64Test{
	{"000", 0},
	{"000FFFFFFFFFFFFFFFF", 18446744073709551615},
	{"FFFFFFFFFFFFFFFF", 18446744073709551615},
}

var hexToUInt64ErrorTests = []errorUInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.HexToUInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.HexToUInt64"},
	{"+", 0, 2, "faiNumber.HexToUInt64"},
	{"FFFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToUInt64"},
	{"FFFFFFFFFFFFFFFFg", 0, 2, "faiNumber.HexToUInt64"},
	// Overflow Errors
	{"FFFFFFFFFFFFFFFF1", 0, 4, "faiNumber.HexToUInt64"},
}
/* Test Cases End */

/* Test Functions Start */
func TestHexToInt32NoError(t *testing.T) {
	for i := range hexToInt32NoErrorTests {
		test := &hexToInt32NoErrorTests[i]
		out, err := HexToInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("HexToInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestHexToInt32Error(t *testing.T) {
	for i := range hexToInt32ErrorTests {
		test := &hexToInt32ErrorTests[i]
		out, err := HexToInt32(test.in)
		if err == nil {
			t.Errorf("HexToInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("HexToInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestHexToUInt32NoError(t *testing.T) {
	for i := range hexToUInt32NoErrorTests {
		test := &hexToUInt32NoErrorTests[i]
		out, err := HexToUInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("HexToUInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestHexToUInt32Error(t *testing.T) {
	for i := range hexToUInt32ErrorTests {
		test := &hexToUInt32ErrorTests[i]
		out, err := HexToUInt32(test.in)
		if err == nil {
			t.Errorf("HexToUInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("HexToUInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestHexToInt64NoError(t *testing.T) {
	for i := range hexToInt64NoErrorTests {
		test := &hexToInt64NoErrorTests[i]
		out, err := HexToInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("HexToInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestHexToInt64Error(t *testing.T) {
	for i := range hexToInt64ErrorTests {
		test := &hexToInt64ErrorTests[i]
		out, err := HexToInt64(test.in)
		if err == nil {
			t.Errorf("HexToInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("HexToInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestHexToUInt64NoError(t *testing.T) {
	for i := range hexToUInt64NoErrorTests {
		test := &hexToUInt64NoErrorTests[i]
		out, err := HexToUInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("HexToUInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestHexToUInt64Error(t *testing.T) {
	for i := range hexToUInt64ErrorTests {
		test := &hexToUInt64ErrorTests[i]
		out, err := HexToUInt64(test.in)
		if err == nil {
			t.Errorf("HexToUInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("HexToUInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}
/* Test Functions End */

/* Benchmark Functions Start*/
func BenchmarkHexToInt32VSstrconv(b *testing.B) {
	case1 := "ABCDEF"
	case2 := "-ABCDEF"
	b.Run("HexToInt32", func(b *testing.B) {
		benchmarkHexToInt32(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 16, 32)
	})
	b.Run("HexToInt32", func(b *testing.B) {
		benchmarkHexToInt32(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 16, 32)
	})
}

func benchmarkHexToInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := HexToInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkHexToUInt32VSstrconv(b *testing.B) {
	case1 := "1A2B3C4D"
	b.Run("HexToUInt32", func(b *testing.B) {
		benchmarkHexToUInt32(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 16, 32)
	})
}

func benchmarkHexToUInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := HexToUInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkHexToInt64VSstrconv(b *testing.B) {
	case1 := "ABCDEFABCDEF"
	case2 := "-ABCDEFABCDEF"
	b.Run("HexToInt64", func(b *testing.B) {
		benchmarkHexToInt64(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 16, 64)
	})
	b.Run("HexToInt64", func(b *testing.B) {
		benchmarkHexToInt64(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 16, 64)
	})
}

func benchmarkHexToInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := HexToInt64(input)
			BenchSink = out
		}
	})
}

func BenchmarkHexToUInt64VSstrconv(b *testing.B) {
	case1 := "1A2B3C4D5E6F"
	b.Run("HexToUInt64", func(b *testing.B) {
		benchmarkHexToUInt64(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 16, 64)
	})
}

func benchmarkHexToUInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := HexToUInt64(input)
			BenchSink = out
		}
	})
}
/* Benchmark Functions End*/
