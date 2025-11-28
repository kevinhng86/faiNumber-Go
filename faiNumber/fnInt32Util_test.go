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
    /* Test to see if fnDualDigitMap working correctly Start */
	{100 ,"100"},
	{101 ,"101"},
	{102 ,"102"},
	{103 ,"103"},
	{104 ,"104"},
	{105 ,"105"},
	{106 ,"106"},
	{107 ,"107"},
	{108 ,"108"},
	{109 ,"109"},
	{110 ,"110"},
	{111 ,"111"},
	{112 ,"112"},
	{113 ,"113"},
	{114 ,"114"},
	{115 ,"115"},
	{116 ,"116"},
	{117 ,"117"},
	{118 ,"118"},
	{119 ,"119"},
	{120 ,"120"},
	{121 ,"121"},
	{122 ,"122"},
	{123 ,"123"},
	{124 ,"124"},
	{125 ,"125"},
	{126 ,"126"},
	{127 ,"127"},
	{128 ,"128"},
	{129 ,"129"},
	{130 ,"130"},
	{131 ,"131"},
	{132 ,"132"},
	{133 ,"133"},
	{134 ,"134"},
	{135 ,"135"},
	{136 ,"136"},
	{137 ,"137"},
	{138 ,"138"},
	{139 ,"139"},
	{140 ,"140"},
	{141 ,"141"},
	{142 ,"142"},
	{143 ,"143"},
	{144 ,"144"},
	{145 ,"145"},
	{146 ,"146"},
	{147 ,"147"},
	{148 ,"148"},
	{149 ,"149"},
	{150 ,"150"},
	{151 ,"151"},
	{152 ,"152"},
	{153 ,"153"},
	{154 ,"154"},
	{155 ,"155"},
	{156 ,"156"},
	{157 ,"157"},
	{158 ,"158"},
	{159 ,"159"},
	{160 ,"160"},
	{161 ,"161"},
	{162 ,"162"},
	{163 ,"163"},
	{164 ,"164"},
	{165 ,"165"},
	{166 ,"166"},
	{167 ,"167"},
	{168 ,"168"},
	{169 ,"169"},
	{170 ,"170"},
	{171 ,"171"},
	{172 ,"172"},
	{173 ,"173"},
	{174 ,"174"},
	{175 ,"175"},
	{176 ,"176"},
	{177 ,"177"},
	{178 ,"178"},
	{179 ,"179"},
	{180 ,"180"},
	{181 ,"181"},
	{182 ,"182"},
	{183 ,"183"},
	{184 ,"184"},
	{185 ,"185"},
	{186 ,"186"},
	{187 ,"187"},
	{188 ,"188"},
	{189 ,"189"},
	{190 ,"190"},
	{191 ,"191"},
	{192 ,"192"},
	{193 ,"193"},
	{194 ,"194"},
	{195 ,"195"},
	{196 ,"196"},
	{197 ,"197"},
	{198 ,"198"},
	{199 ,"199"},
	/* Test to see if fnDualDigitMap working correctly Start */
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
