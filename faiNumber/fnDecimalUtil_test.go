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

/* Testing Struct Declaration Start */
type noErrorInt32Test struct {
	in  string
	out int32
}

type errorInt32Test struct {
	in  string
	out int32
	errCode int
	errFromFunction string
}

type noErrorUInt32Test struct {
	in  string
	out uint32
}

type errorUInt32Test struct {
	in  string
	out uint32
	errCode int
	errFromFunction string
}

type noErrorInt64Test struct {
	in  string
	out int64
}

type errorInt64Test struct {
	in  string
	out int64
	errCode int
	errFromFunction string
}

type noErrorUInt64Test struct {
	in  string
	out uint64
}

type errorUInt64Test struct {
	in  string
	out uint64
	errCode int
	errFromFunction string
}
/* Testing Struct Declaration End */

/* Test Cases Start */
var decToInt32NoErrorTests = []noErrorInt32Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-0002147483648", -2147483648},
	{"-2147483648", -2147483648},
	{"-2147483637", -2147483637},
	{"-2147483647", -2147483647},
	{"-214748364", -214748364},
	{"0002147483647", 2147483647},
	{"2147483647", 2147483647},
	{"2147483637", 2147483637},
	{"2147483646", 2147483646},
	{"214748364", 214748364},
	{"+0002147483647", 2147483647},
	{"+2147483647", 2147483647},
	{"+2147483637", 2147483637},
	{"+2147483646", 2147483646},
	{"+214748364", 214748364},
}

var decToInt32ErrorTests = []errorInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.DecToInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.DecToInt32"},
	{"+", 0, 2, "faiNumber.DecToInt32"},
	{"-314748a648", 0, 2, "faiNumber.DecToInt32"},
	{"-214748364a", 0, 2, "faiNumber.DecToInt32"},
	{"-214748363a", 0, 2, "faiNumber.DecToInt32"},
	{"-214748a648", 0, 2, "faiNumber.DecToInt32"},
	{"-2147483648a", 0, 2, "faiNumber.DecToInt32"},
	{"-21474836a", 0, 2, "faiNumber.DecToInt32"},
	{"+31474a3647", 0, 2, "faiNumber.DecToInt32"},
	{"+214748364a", 0, 2, "faiNumber.DecToInt32"},
	{"+214748363a", 0, 2, "faiNumber.DecToInt32"},
	{"+21474a3647", 0, 2, "faiNumber.DecToInt32"},
	{"+21474836a", 0, 2, "faiNumber.DecToInt32"},
	{"31474a3647", 0, 2, "faiNumber.DecToInt32"},
	{"214748364a", 0, 2, "faiNumber.DecToInt32"},
	{"214748363a", 0, 2, "faiNumber.DecToInt32"},
	{"21474a3647", 0, 2, "faiNumber.DecToInt32"},
	{"21474836a", 0, 2, "faiNumber.DecToInt32"},
	// Underflow Errors
	{"-2147483649", 0, 3, "faiNumber.DecToInt32"},
	{"-21474836481", 0, 3, "faiNumber.DecToInt32"},
	{"-3147483648", 0, 3, "faiNumber.DecToInt32"},
	// Overflow Errors
	{"2147483648", 0, 4, "faiNumber.DecToInt32"},
	{"21474836471", 0, 4, "faiNumber.DecToInt32"},
	{"3147483647", 0, 4, "faiNumber.DecToInt32"},
	{"+2147483648", 0, 4, "faiNumber.DecToInt32"},
	{"+21474836471", 0, 4, "faiNumber.DecToInt32"},
	{"+3147483647", 0, 4, "faiNumber.DecToInt32"},
}

var decToUInt32NoErrorTests = []noErrorUInt32Test{
	{"000", 0},
	{"0004294967295", 4294967295},
	{"4294967295", 4294967295},
	{"4294967294", 4294967294},
	{"429496729", 429496729},
}

var decToUInt32ErrorTests = []errorUInt32Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.DecToUInt32"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.DecToUInt32"},
	{"+", 0, 2, "faiNumber.DecToUInt32"},
	{"529496729a", 0, 2, "faiNumber.DecToUInt32"},
	{"429496729a", 0, 2, "faiNumber.DecToUInt32"},
	{"4294967295a", 0, 2, "faiNumber.DecToUInt32"},
	{"42949672a", 0, 2, "faiNumber.DecToUInt32"},	
	// Overflow Errors
	{"5294967295", 0, 4, "faiNumber.DecToUInt32"},
	{"4294967296", 0, 4, "faiNumber.DecToUInt32"},
	{"42949672951", 0, 4, "faiNumber.DecToUInt32"},
}

var decToInt64NoErrorTests = []noErrorInt64Test{
	{"000", 0},
	{"-00", 0},
	{"+000", 0},
	{"-0009223372036854775808", -9223372036854775808},
	{"-9223372036854775808", -9223372036854775808},
	{"-9223372036854775798", -9223372036854775798},
	{"-9223372036854775807", -9223372036854775807},
	{"-922337203685477580", -922337203685477580},
	{"0009223372036854775807", 9223372036854775807},
	{"9223372036854775807", 9223372036854775807},
	{"9223372036854775797", 9223372036854775797},
	{"9223372036854775806", 9223372036854775806},
	{"922337203685477580", 922337203685477580},
	{"+0009223372036854775807", 9223372036854775807},
	{"+9223372036854775807", 9223372036854775807},
	{"+9223372036854775797", 9223372036854775797},
	{"+9223372036854775806", 9223372036854775806},
	{"+922337203685477580", 922337203685477580},
}

var decToInt64ErrorTests = []errorInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.DecToInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.DecToInt64"},
	{"+", 0, 2, "faiNumber.DecToInt64"},
	{"-9223372036854775808a", 0, 2, "faiNumber.DecToInt64"},
	{"-922337203685477580a", 0, 2, "faiNumber.DecToInt64"},
	{"-92233720368547758a", 0, 2, "faiNumber.DecToInt64"},
	{"-922337203685477581a", 0, 2, "faiNumber.DecToInt64"},
	{"9223372036854775807a", 0, 2, "faiNumber.DecToInt64"},
	{"922337203685477580a", 0, 2, "faiNumber.DecToInt64"},
	{"92233720368547758a", 0, 2, "faiNumber.DecToInt64"},
	{"922337203685477581a", 0, 2, "faiNumber.DecToInt64"},
	{"+9223372036854775807a", 0, 2, "faiNumber.DecToInt64"},
	{"+922337203685477580a", 0, 2, "faiNumber.DecToInt64"},
	{"+92233720368547758a", 0, 2, "faiNumber.DecToInt64"},
	{"+922337203685477581a", 0, 2, "faiNumber.DecToInt64"},
	// Underflow Errors
	{"-9223372036854775809", 0, 3, "faiNumber.DecToInt64"},
	{"-92233720368547758081", 0, 3, "faiNumber.DecToInt64"},
	{"-9223372036854775818", 0, 3, "faiNumber.DecToInt64"},
	// Overflow Errors
	{"9223372036854775808", 0, 4, "faiNumber.DecToInt64"},
	{"92233720368547758071", 0, 4, "faiNumber.DecToInt64"},
	{"9223372036854775817", 0, 4, "faiNumber.DecToInt64"},
	{"+9223372036854775808", 0, 4, "faiNumber.DecToInt64"},
	{"+92233720368547758071", 0, 4, "faiNumber.DecToInt64"},
	{"+9223372036854775817", 0, 4, "faiNumber.DecToInt64"},
}

var decToUInt64NoErrorTests = []noErrorUInt64Test{
	{"000", 0},
	{"00018446744073709551615", 18446744073709551615},
	{"18446744073709551615", 18446744073709551615},
	{"18446744073709551614", 18446744073709551614},
	{"1844674407370955161", 1844674407370955161},
}

var decToUInt64ErrorTests = []errorUInt64Test{
	// Empty String Error
	{"", 0, 1, "faiNumber.DecToUInt64"},
	// Invalid Format Error
	{"-", 0, 2, "faiNumber.DecToUInt64"},
	{"+", 0, 2, "faiNumber.DecToUInt64"},
	{"18446744073709551615a", 0, 2, "faiNumber.DecToUInt64"},
	{"1844674407370955161a", 0, 2, "faiNumber.DecToUInt64"},
	{"2844674407370955161a", 0, 2, "faiNumber.DecToUInt64"},
	{"1844674407370955162a", 0, 2, "faiNumber.DecToUInt64"},
	{"184467440737095516a", 0, 2, "faiNumber.DecToUInt64"},
	// Overflow Errors
	{"184467440737095516151", 0, 4, "faiNumber.DecToUInt64"},
	{"18446744073709551616", 0, 4, "faiNumber.DecToUInt64"},
	{"18446744073709551625", 0, 4, "faiNumber.DecToUInt64"},
	{"28446744073709551615", 0, 4, "faiNumber.DecToUInt64"},
}
/* Test Cases End */

/* Test Functions Start */
func TestDecToInt32NoError(t *testing.T) {
	for i := range decToInt32NoErrorTests {
		test := &decToInt32NoErrorTests[i]
		out, err := DecToInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("DecToInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestDecToInt32Error(t *testing.T) {
	for i := range decToInt32ErrorTests {
		test := &decToInt32ErrorTests[i]
		out, err := DecToInt32(test.in)
		if err == nil {
			t.Errorf("DecToInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("DecToInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestDecToUInt32NoError(t *testing.T) {
	for i := range decToUInt32NoErrorTests {
		test := &decToUInt32NoErrorTests[i]
		out, err := DecToUInt32(test.in)
		if test.out != out || err != nil {
			t.Errorf("DecToUInt32(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestDecToUInt32Error(t *testing.T) {
	for i := range decToUInt32ErrorTests {
		test := &decToUInt32ErrorTests[i]
		out, err := DecToUInt32(test.in)
		if err == nil {
			t.Errorf("DecToUInt32(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("DecToUInt32(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestDecToInt64NoError(t *testing.T) {
	for i := range decToInt64NoErrorTests {
		test := &decToInt64NoErrorTests[i]
		out, err := DecToInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("DecToInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestDecToInt64Error(t *testing.T) {
	for i := range decToInt64ErrorTests {
		test := &decToInt64ErrorTests[i]
		out, err := DecToInt64(test.in)
		if err == nil {
			t.Errorf("DecToInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("DecToInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}

func TestDecToUInt64NoError(t *testing.T) {
	for i := range decToUInt64NoErrorTests {
		test := &decToUInt64NoErrorTests[i]
		out, err := DecToUInt64(test.in)
		if test.out != out || err != nil {
			t.Errorf("DecToUInt64(%q) = %v, Error = %v want %v",
				test.in, out, err, test.out)
		}
	}
}

func TestDecToUInt64Error(t *testing.T) {
	for i := range decToUInt64ErrorTests {
		test := &decToUInt64ErrorTests[i]
		out, err := DecToUInt64(test.in)
		if err == nil {
			t.Errorf("DecToUInt64(%q) = %v, Error = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err, test.out, test.errCode, test.errFromFunction)
		}

		if err.(*FNError).Code != test.errCode || err.(*FNError).FromFunction != test.errFromFunction || out != test.out {
			t.Errorf("DecToUInt64(%q) = %v, Error Code = %v, Error Function = %v\nWant Value: %v\nWant Error Code: %v\nWant Function Name: %v",
				test.in, out, err.(*FNError).Code, err.(*FNError).FromFunction, test.out, test.errCode, test.errFromFunction)
		}
	}
}
/* Test Functions End */

/* Benchmark Data Type Start */
var BenchSink any
/* Benchmark Data Type End */

/* strconv Benchmark Functions Start */
func benchmarkParseInt(b *testing.B, name string, input string, base int, bitSize int) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := strconv.ParseInt(input, base, bitSize)
			BenchSink = out
		}
	})
}

func benchmarkParseUint(b *testing.B, name string, input string, base int, bitSize int) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := strconv.ParseUint(input, base, bitSize)
			BenchSink = out
		}
	})
}
/* strconv Benchmark Functions End */

/* Benchmark Functions Start*/
func BenchmarkDecToInt32VSstrconv(b *testing.B) {
	case1 := "1234567890"
	case2 := "-1234567890"
	b.Run("DecToInt32", func(b *testing.B) {
		benchmarkDecToInt32(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 10, 32)
	})
	b.Run("DecToInt32", func(b *testing.B) {
		benchmarkDecToInt32(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 10, 32)
	})
}

func benchmarkDecToInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := DecToInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkDecToUInt32VSstrconv(b *testing.B) {
	case1 := "123456789"
	b.Run("DecToUInt32", func(b *testing.B) {
		benchmarkDecToUInt32(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 10, 32)
	})
}

func benchmarkDecToUInt32(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := DecToUInt32(input)
			BenchSink = out
		}
	})
}

func BenchmarkDecToInt64VSstrconv(b *testing.B) {
	case1 := "3546874589456"
	case2 := "-3546874589456"
	b.Run("DecToInt64", func(b *testing.B) {
		benchmarkDecToInt64(b, case1, case1)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case1, case1, 10, 64)
	})
	b.Run("DecToInt64", func(b *testing.B) {
		benchmarkDecToInt64(b, case2, case2)
	})
	b.Run("ParseInt", func(b *testing.B) {
		benchmarkParseInt(b, case2, case2, 10, 64)
	})
}

func benchmarkDecToInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := DecToInt64(input)
			BenchSink = out
		}
	})
}

func BenchmarkDecToUInt64VSstrconv(b *testing.B) {
	case1 := "56478945612345"
	b.Run("DecToUInt64", func(b *testing.B) {
		benchmarkDecToUInt64(b, case1, case1)
	})
	b.Run("ParseUint", func(b *testing.B) {
		benchmarkParseUint(b, case1, case1, 10, 64)
	})
}

func benchmarkDecToUInt64(b *testing.B, name string, input string) {
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out, _ := DecToUInt64(input)
			BenchSink = out
		}
	})
}
/* Benchmark Functions End*/
