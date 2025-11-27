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

package faiNumber_test

import (
	"fmt"
	"github.com/kevinhng86/faiNumber-Go/faiNumber"
)

/* Decimal To Int Example Functions Start */
func ExampleDecToInt32() {
	str := "-5487458"

	output, err := faiNumber.DecToInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// int32, -5487458
}

func ExampleDecToUInt32() {
	str := "987"

	output, err := faiNumber.DecToUInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// uint32, 987
}

func ExampleDecToInt64() {
	str := "-987456321487564"

	output, err := faiNumber.DecToInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// int64, -987456321487564
}

func ExampleDecToUInt64() {
	str := "98745641236549"

	output, err := faiNumber.DecToUInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// uint64, 98745641236549
}
/* Decimal To Int Example Functions End */

/* Binary To Int Example Functions Start */
func ExampleBinaryToInt32() {
	str := "-10101010"

	output, err := faiNumber.BinaryToInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %b\n", output, output)
	
	// Output:
	// int32, -10101010
}

func ExampleBinaryToUInt32() {
	str := "1111100000"

	output, err := faiNumber.BinaryToUInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %b\n", output, output)
	
	// Output:
	// uint32, 1111100000
}

func ExampleBinaryToInt32AsUnsigned() {
	str := "11111111111111111111111111111111"

	output, err := faiNumber.BinaryToInt32AsUnsigned(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// int32, -1
}

func ExampleBinaryToInt64() {
	str := "-11111001101100001111110011011110101"

	output, err := faiNumber.BinaryToInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %b\n", output, output)
	
	// Output:
	// int64, -11111001101100001111110011011110101
}

func ExampleBinaryToUInt64() {
	str := "1110110110110101011001010100101010110010101"

	output, err := faiNumber.BinaryToUInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %b\n", output, output)
	
	// Output:
	// uint64, 1110110110110101011001010100101010110010101
}

func ExampleBinaryToInt64AsUnsigned() {
	str := "1111111111111111111111111111111111111111111111111111111111111111"

	output, err := faiNumber.BinaryToInt64AsUnsigned(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %v\n", output, output)
	
	// Output:
	// int64, -1
}
/* Binary To Int Example Functions End */

/* Octal To Int Example Functions Start */
func ExampleOctalToInt32() {
	str := "-7434741"

	output, err := faiNumber.OctalToInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %o\n", output, output)
	
	// Output:
	// int32, -7434741
}

func ExampleOctalToUInt32() {
	str := "3741471"

	output, err := faiNumber.OctalToUInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %o\n", output, output)
	
	// Output:
	// uint32, 3741471
}

func ExampleOctalToInt64() {
	str := "-5434741567"

	output, err := faiNumber.OctalToInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %o\n", output, output)
	
	// Output:
	// int64, -5434741567
}

func ExampleOctalToUInt64() {
	str := "7534534754"

	output, err := faiNumber.OctalToUInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %o\n", output, output)
	
	// Output:
	// uint64, 7534534754
}
/* Octal To Int Example Functions End */

/* Hex To Int Example Functions Start */
func ExampleHexToInt32() {
	str := "-F10A"

	output, err := faiNumber.HexToInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %X\n", output, output)
	
	// Output:
	// int32, -F10A
}

func ExampleHexToUInt32() {
	str := "FFFFAAAA"

	output, err := faiNumber.HexToUInt32(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %X\n", output, output)
	
	// Output:
	// uint32, FFFFAAAA
}

func ExampleHexToInt64() {
	str := "-8ABCDEFABC"

	output, err := faiNumber.HexToInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %X\n", output, output)
	
	// Output:
	// int64, -8ABCDEFABC
}

func ExampleHexToUInt64() {
	str := "FABC1234FABC"

	output, err := faiNumber.HexToUInt64(str)
	if err != nil {
		fmt.Println("Error handling")
	}
	fmt.Printf("%T, %X\n", output, output)
	
	// Output:
	// uint64, FABC1234FABC
}
/* Hex To Int Example Functions End */

/* Int To Decimal Example Functions Start */
func ExampleInt32ToDec() {
	var number int32 = -98745
	output := faiNumber.Int32ToDec(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -98745
}

func ExampleUInt32ToDec() {
	var number uint32 = 35478567
	output := faiNumber.UInt32ToDec(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 35478567
}

func ExampleInt64ToDec() {
	var number int64 = -45756787452134
	output := faiNumber.Int64ToDec(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -45756787452134
}

func ExampleUInt64ToDec() {
	var number uint64 = 547456745412
	output := faiNumber.UInt64ToDec(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 547456745412
}
/* Int To Decimal Example Functions End */

/* Int To Binary Example Functions Start */
func ExampleInt32ToBinary() {
	var number int32 = -98745
	output := faiNumber.Int32ToBinary(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -11000000110111001
}

func ExampleUInt32ToBinary() {
	var number uint32 = 35478567
	output := faiNumber.UInt32ToBinary(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 10000111010101110000100111
}

func ExampleInt32ToBinaryAsUnsigned() {
	var number int32 = -1
	output := faiNumber.Int32ToBinaryAsUnsigned(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 11111111111111111111111111111111
}

func ExampleInt64ToBinary() {
	var number int64 = -45756787452134
	output := faiNumber.Int64ToBinary(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -1010011001110110010101000011111110100011100110
}

func ExampleUInt64ToBinary() {
	var number uint64 = 547456745412
	output := faiNumber.UInt64ToBinary(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 111111101110110111101101111111111000100
}

func ExampleInt64ToBinaryAsUnsigned() {
	var number int64 = -1
	output := faiNumber.Int64ToBinaryAsUnsigned(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 1111111111111111111111111111111111111111111111111111111111111111
}
/* Int To Binary Example Functions End */

/* Int To Octal Example Functions Start */
func ExampleInt32ToOctal() {
	var number int32 = -98745
	output := faiNumber.Int32ToOctal(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -300671
}

func ExampleUInt32ToOctal() {
	var number uint32 = 35478567
	output := faiNumber.UInt32ToOctal(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 207256047
}

func ExampleInt64ToOctal() {
	var number int64 = -45756787452134
	output := faiNumber.Int64ToOctal(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -1231662503764346
}

func ExampleUInt64ToOctal() {
	var number uint64 = 547456745412
	output := faiNumber.UInt64ToOctal(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 7756675577704
}
/* Int To Octal Example Functions End */

/* Int To Hex Example Functions Start */
func ExampleInt32ToHex() {
	var number int32 = -98745
	output := faiNumber.Int32ToHex(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -181B9
}

func ExampleUInt32ToHex() {
	var number uint32 = 35478567
	output := faiNumber.UInt32ToHex(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 21D5C27
}

func ExampleInt64ToHex() {
	var number int64 = -45756787452134
	output := faiNumber.Int64ToHex(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, -299D950FE8E6
}

func ExampleUInt64ToHex() {
	var number uint64 = 547456745412
	output := faiNumber.UInt64ToHex(number)
	fmt.Printf("%T, %v\n", output, output)

	// Output:
	// string, 7F76F6FFC4
}
/* Int To Hex Example Functions End */

/* Error Example Start */
func ExampleFNError() {
	str := "1A"
	if _, err := faiNumber.DecToInt32(str); err != nil {
		e := err.(*faiNumber.FNError)
		fmt.Println("From Function:", e.FromFunction)
		fmt.Println("Input:", e.Input)
		fmt.Println("Error Message:", e.Message)
		fmt.Println("Error Code:", e.Code)
		fmt.Println(err)
	}

	// Output:
	// From Function: faiNumber.DecToInt32
	// Input: 1A
	// Error Message: Invalid format error
	// Error Code: 2
	// Error from function faiNumber.DecToInt32. Error code: 2. Input value: "1A". Message: Invalid format error.
}
/* Error Example End */
