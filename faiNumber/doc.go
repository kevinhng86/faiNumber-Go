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

// Package faiNumber mainly deals with converting numerical strings to a value of an int
// data type, converting a value of an int data type to a numerical string. faiNumber is an extremely
// fast library for parsing string and converting int values to strings. Be it whether if it is
// an overflow, underflow, empty strings where it shouldn't be, or an invalid input,
// faiNumber functions were designed to always return an error with a code when an error
// has occurred where it is possible to have an error.
//
// faiNumber is extremely fast. faiNumber may be the fastest golang library for numerical string parsing
// and converting an int data type to a numerical string. faiNumber functions were tested and is much faster
// than strconv package. But faiNumber only support parsing decimal, binary, octal, hexadecimal string
// to int32, uint32, int64 or uint64 data types. For converting an int value to a numerical string,
// faiNumber only support convert int32, uint32, int64 or uint64 to a string of decimal, binary, octal
// or hexadecimal. 
//
// Each of faiNumber parsing function were written separately for each usage case and they were optimize
// for speed. For converting to and from each data type of int32, uint32, int64, uint64, a separate function
// was created to optimize for speed and ensure no type casting is required for using any of faiNumber function.
//
// # Benchmark
//
// faiNumber's benchmark compare to strconv can be found here:
// https://github.com/kevinhng86/faiNumber-Go/blob/main/benchmark.md
//
// # Parsing numerical strings to int data type
//
// For converting string to int faiNumber has the following functions: [DecToInt32], [DecToUInt32],
// [DecToInt64], [DecToUInt64], [BinaryToInt32], [BinaryToInt32AsUnsigned], [BinaryToUInt32], [BinaryToInt64],
// [BinaryToInt64AsUnsigned], [BinaryToUInt64], [OctalToInt32], [OctalToUInt32], [OctalToInt64], [OctalToUInt64],
// [HexToInt32], [HexToUInt32], [HexToInt64], and [HexToUInt64].
//
//	i, err := faiNumber.DecToInt32("-99")
//	i, err := faiNumber.DecToUInt32("99")
//	i, err := faiNumber.DecToInt64("-102")
//	i, err := faiNumber.DecToUInt64("102")
//	i, err := faiNumber.BinaryToInt32("-101")
//	i, err := faiNumber.BinaryToInt32AsUnsigned("10101010")
//	i, err := faiNumber.BinaryToUInt32("1010")
//	i, err := faiNumber.BinaryToInt64("-11110000")
//	i, err := faiNumber.BinaryToInt64AsUnsigned("11110000")
//	i, err := faiNumber.BinaryToUInt64("11110000")
//	i, err := faiNumber.OctalToInt32("-17")
//	i, err := faiNumber.OctalToUInt32("17")
//	i, err := faiNumber.OctalToInt64("-3000")
//	i, err := faiNumber.OctalToUInt64("3000")
//	i, err := faiNumber.HexToInt32("-FF")
//	i, err := faiNumber.HexToUInt32("FF")
//	i, err := faiNumber.HexToInt64("-AB")
//	i, err := faiNumber.HexToUInt64("AB")
//
// # Converting an int value to string
//
// For converting an int value to string faiNumber has the following functions: [Int32ToDec], [UInt32ToDec],
// [Int32ToBinary], [UInt32ToBinary], [Int32ToBinaryAsUnsigned], [Int32ToOctal], [UInt32ToOctal], [Int32ToHex],
// [UInt32ToHex], [Int64ToDec], [UInt64ToDec], [Int64ToBinary], [UInt64ToBinary], [Int64ToBinaryAsUnsigned],
// [Int64ToOctal], [UInt64ToOctal], [Int64ToHex], and [UInt64ToHex].
package faiNumber
