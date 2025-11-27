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
var octalToInt32NoErrorTests = []noErrorInt32Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-00020000000000", -2147483648},	
	{"-20000000000", -2147483648},
	{"-17777777777", -2147483647},
	{"-1777777777", -268435455},
	{"00017777777777", 2147483647},
	{"17777777777", 2147483647},
	{"1777777777", 268435455},
	{"+00017777777777", 2147483647},
	{"+17777777777", 2147483647},
	{"+1777777777", 268435455},
}

var octalToInt32ErrorTests = []errorInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.OctalToInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.OctalToInt32"},
	{"+", 0, 2, "faiNumber.OctalToInt32"},
	{"-21800000000", 0, 2, "faiNumber.OctalToInt32"},
	{"-200000000008", 0, 2, "faiNumber.OctalToInt32"},
	{"-17777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"-1777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"17777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"177777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"1777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"+17777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"+177777777778", 0, 2, "faiNumber.OctalToInt32"},
	{"+1777777778", 0, 2, "faiNumber.OctalToInt32"},	
	// Underflow Errors
	{"-20000000001", 0, 3, "faiNumber.OctalToInt32"},
	{"-200000000001", 0, 3, "faiNumber.OctalToInt32"},
	// Overflow Errors
	{"20000000000", 0, 4, "faiNumber.OctalToInt32"},
	{"177777777771", 0, 4, "faiNumber.OctalToInt32"},
}

var octalToUInt32NoErrorTests = []noErrorUInt32Test{
	{"000", 0},
	{"00037777777777", 4294967295},
	{"37777777777", 4294967295},
	{"3777777777", 536870911},
}

var octalToUInt32ErrorTests = []errorUInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.OctalToUInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.OctalToUInt32"},
	{"+", 0, 2, "faiNumber.OctalToUInt32"},
	{"40000000008", 0, 2, "faiNumber.OctalToUInt32"},
	{"377777777778", 0, 2, "faiNumber.OctalToUInt32"},
	{"37777777778", 0, 2, "faiNumber.OctalToUInt32"},
	{"3777777778", 0, 2, "faiNumber.OctalToUInt32"},
	// Overflow Errors
	{"40000000000", 0, 4, "faiNumber.OctalToUInt32"},
	{"377777777777", 0, 4, "faiNumber.OctalToUInt32"},
}

var octalToInt64NoErrorTests = []noErrorInt64Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-0001000000000000000000000", -9223372036854775808},
	{"-1000000000000000000000", -9223372036854775808},
	{"-777777777777777777777", -9223372036854775807},
	{"000777777777777777777777", 9223372036854775807},
	{"777777777777777777777", 9223372036854775807},
	{"+000777777777777777777777", 9223372036854775807},
	{"+777777777777777777777", 9223372036854775807},	
}

var octalToInt64ErrorTests = []errorInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.OctalToInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.OctalToInt64"},
	{"+", 0, 2, "faiNumber.OctalToInt64"},
	{"-1180000000000000000000", 0, 2, "faiNumber.OctalToInt64"},
	{"-10000000000000000000008", 0, 2, "faiNumber.OctalToInt64"},
	{"-777777777777777777778", 0, 2, "faiNumber.OctalToInt64"},
	{"7777777777777777777778", 0, 2, "faiNumber.OctalToInt64"},
	{"777777777777777777778", 0, 2, "faiNumber.OctalToInt64"},
	{"+7777777777777777777778", 0, 2, "faiNumber.OctalToInt64"},
	{"+777777777777777777778", 0, 2, "faiNumber.OctalToInt64"},
	// Underflow Errors
	{"-1000000000000000000001", 0, 3, "faiNumber.OctalToInt64"},
	{"-10000000000000000000000", 0, 3, "faiNumber.OctalToInt64"},
	// Overflow Errors
	{"1000000000000000000000", 0, 4, "faiNumber.OctalToInt64"},
	{"+1000000000000000000000", 0, 4, "faiNumber.OctalToInt64"},
}

var octalToUInt64NoErrorTests = []noErrorUInt64Test{
	{"000", 0},
	{"0001777777777777777777777", 18446744073709551615},
	{"1777777777777777777777", 18446744073709551615},
	{"177777777777777777777", 2305843009213693951},
}

var octalToUInt64ErrorTests = []errorUInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.OctalToUInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.OctalToUInt64"},
	{"+", 0, 2, "faiNumber.OctalToUInt64"},
	{"1777777777777777777778", 0, 2, "faiNumber.OctalToUInt64"},
	{"2777777777777777777778", 0, 2, "faiNumber.OctalToUInt64"},
	{"17777777777777777777778", 0, 2, "faiNumber.OctalToUInt64"},
	{"177777777777777777778", 0, 2, "faiNumber.OctalToUInt64"},
	// Overflow Errors
	{"2777777777777777777777", 0, 4, "faiNumber.OctalToUInt64"},
	{"17777777777777777777777", 0, 4, "faiNumber.OctalToUInt64"},
}
/* Test Cases End */

/* Test Functions Start */
func TestOctalToInt32NoError(t *testing.T) {
	for i := range octalToInt32NoErrorTests {
		test := &octalToInt32NoErrorTests[i]
		out, err := OctalToInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("OctalToInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestOctalToInt32Error(t *testing.T) {
	for i := range octalToInt32ErrorTests {
		test := &octalToInt32ErrorTests[i]
		out, err := OctalToInt32(test.in)
		if err == nil {
			t.Errorf("OctalToInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("OctalToInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestOctalToUInt32NoError(t *testing.T) {
	for i := range octalToUInt32NoErrorTests {
		test := &octalToUInt32NoErrorTests[i]
		out, err := OctalToUInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("OctalToUInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestOctalToUInt32Error(t *testing.T) {
	for i := range octalToUInt32ErrorTests {
		test := &octalToUInt32ErrorTests[i]
		out, err := OctalToUInt32(test.in)
		if err == nil {
			t.Errorf("OctalToUInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("OctalToUInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestOctalToInt64NoError(t *testing.T) {
	for i := range octalToInt64NoErrorTests {
		test := &octalToInt64NoErrorTests[i]
		out, err := OctalToInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("OctalToInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestOctalToInt64Error(t *testing.T) {
	for i := range octalToInt64ErrorTests {
		test := &octalToInt64ErrorTests[i]
		out, err := OctalToInt64(test.in)
		if err == nil {
			t.Errorf("OctalToInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("OctalToInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestOctalToUInt64NoError(t *testing.T) {
	for i := range octalToUInt64NoErrorTests {
		test := &octalToUInt64NoErrorTests[i]
		out, err := OctalToUInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("OctalToUInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestOctalToUInt64Error(t *testing.T) {
	for i := range octalToUInt64ErrorTests {
		test := &octalToUInt64ErrorTests[i]
		out, err := OctalToUInt64(test.in)
		if err == nil {
			t.Errorf("OctalToUInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("OctalToUInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}
/* Test Functions End */

/* Benchmark Functions Start*/
func BenchmarkOctalToInt32VSstrconv(b *testing.B) {
	case1 := "17777777777"
	case2 := "-17777777777"
	b.Run("OctalToInt32", func(b *testing.B) {
		benchmarkOctalToInt32(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 8, 32)
	})
	b.Run("OctalToInt32", func(b *testing.B) {
		benchmarkOctalToInt32(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 8, 32)
	})
}

func benchmarkOctalToInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := OctalToInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkOctalToUInt32VSstrconv(b *testing.B) {
	case1 := "37777777777"
	b.Run("OctalToUInt32", func(b *testing.B) {
		benchmarkOctalToUInt32(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 8, 32)
	})
}

func benchmarkOctalToUInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := OctalToUInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkOctalToInt64VSstrconv(b *testing.B) {
	case1 := "777777777777777777777"
	case2 := "-777777777777777777777"
	b.Run("OctalToInt64", func(b *testing.B) {
		benchmarkOctalToInt64(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 8, 64)
	})
	b.Run("OctalToInt64", func(b *testing.B) {
		benchmarkOctalToInt64(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 8, 64)
	})
}

func benchmarkOctalToInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := OctalToInt64(input)
			BenchSink = out
		}
	})
}

func BenchmarkOctalToUInt64VSstrconv(b *testing.B) {
	case1 := "1777777777777777777777"
	b.Run("OctalToUInt64", func(b *testing.B) {
		benchmarkOctalToUInt64(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 8, 64)
	})
}

func benchmarkOctalToUInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := OctalToUInt64(input)
			BenchSink = out
		}
	})
}
/* Benchmark Functions End*/
