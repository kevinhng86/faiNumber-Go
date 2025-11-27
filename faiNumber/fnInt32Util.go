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

// This is a number value to digit map that use for converting int to hexadecimal string.
const fnNumberToDigitMap = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// This is a dual digit map use for converting int to a decimal string 2 digits at a time algorithm.
const fnDualDigitMap string =
        "0001020304050607080910111213141516171819" +
        "2021222324252627282930313233343536373839" +
        "4041424344454647484950515253545556575859" +
        "6061626364656667686970717273747576777879" +
        "8081828384858687888990919293949596979899"

// Int32ToDec returns a decimal string representation of the input value of the int32 data type.
func Int32ToDec (number int32) (str string) {
    // Min value -2147483648. Length 10 + negative sign is 11
    var isNeg bool
	var t [11]byte 
	i := 10

    if (number < 0){
        isNeg = true
        if ( number == -2147483648 ){
            str = "-2147483648"
            return
        }
        number = -number
    }

	for number >= 100 {
		q100 := (number % 100) << 1
		t[i] = fnDualDigitMap[q100 + 1]
		t[i - 1] = fnDualDigitMap[q100]
		i = i - 2
		number = number / 100
	}

	if ( number >= 10 ){
		t[i] = fnDualDigitMap[(number << 1) + 1]
		i = i - 1
		t[i] = fnDualDigitMap[number << 1]
	} else {
		t[i] = byte(number ^ 48)
    }
    
	if isNeg {
		i--
		t[i] = '-'
	}

	str = string(t[i:])
	return
}

// UInt32ToDec returns a decimal string representation of the input value of the uint32 data type.
func UInt32ToDec (number uint32) (str string) {
    //  Max value 4294967295. Length 10.
    var t [10]byte 
	i := 9

	for number >= 100 {
		q100 := (number % 100) << 1
		t[i] = fnDualDigitMap[q100 + 1]
		t[i - 1] = fnDualDigitMap[q100]
		i = i - 2
		number = number / 100
	}

	if ( number >= 10 ){
		t[i] = fnDualDigitMap[(number << 1) + 1]
		i = i - 1
		t[i] = fnDualDigitMap[number << 1]
	} else {
		t[i] = byte(number ^ 48)
    }

	str = string(t[i:])
	return
}

// Int32ToBinary returns a binary string representation of the input value of the int32 data type.
func Int32ToBinary (number int32) (str string) {
    // Min value -10000000000000000000000000000000 in binary. Length 32 + negative sign is 33.
    // Min value is manually return so the longest value to be return automatically is 31 +
    // negative sign is 32.
    var isNeg bool
	var t [32]byte 
	i := 32

    if (number < 0){
        isNeg = true
        if ( number == -2147483648 ){
            str = "-10000000000000000000000000000000"
            return
        }
        number = -number
    }

	for number >= 2 {
		i--
		t[i] = byte((number & 1) ^ 48)
		number >>= 1
	}

	i--
	t[i] = byte(number ^ 48)

	if isNeg {
		i--
		t[i] = '-'
	}

	str = string(t[i:])
	return
}

// UInt32ToBinary returns a binary string representation of the input value of the uint32 data type.
func UInt32ToBinary (number uint32) (str string) {
    // Max value 11111111111111111111111111111111. Length 32.
	var t [32]byte
	i := 32

	for number >= 2 {
		i--
		t[i] = byte((number & 1) ^ 48)
		number >>= 1
	}

	i--
	t[i] = byte(number ^ 48)

	str = string(t[i:])
	return
}

// Int32ToBinaryAsUnsigned returns a binary string representation of the input value of the int32 type,
// This function treat the input value of type int32 as unsigned. This function returns the raw bit value
// of the input value of the int32 data type and it doesn't interprete negative values as negative values.
func Int32ToBinaryAsUnsigned (number int32) (str string) {
    str = UInt32ToBinary(uint32(number))
    return
}

// Int32ToOctal returns an octal string representation of the input value of the int32 data type.
func Int32ToOctal (number int32) (str string) {
    // Min value -20000000000 in octal. Length is 11 + negative sign is 12.
    // Min value is manually return. The longest value to be return automatically is -17777777777.
    // Length is 11 + negative sign is 12.
    var isNeg bool
	var t [12]byte 
	i := 12

    if (number < 0){
        isNeg = true
        if ( number == -2147483648 ){
            str = "-20000000000"
            return
        }
        number = -number
    }

	for number >= 8 {
		i--
		t[i] = byte((number & 7) ^ 48)
		number >>= 3
	}

	i--
	t[i] = byte(number ^ 48)

	if isNeg {
		i--
		t[i] = '-'
	}

	str = string(t[i:])
	return
}

// UInt32ToOctal returns an octal string representation of the input value of the uint32 data type.
func UInt32ToOctal (number uint32) (str string) {
    // Max value 37777777777. Length 11.
	var t [11]byte
	i := 11

	for number >= 8 {
		i--
		t[i] = byte((number & 7) ^ 48)
		number >>= 3
	}

	i--
	t[i] = byte(number ^ 48)

	str = string(t[i:])
	return
}

// Int32ToHex returns a hexadecimal string representation of the input value of the int32 data type.
// This function uses uppercase letters 'A' to 'F' for digit values >= 10.
func Int32ToHex (number int32) (str string) {
    // Min value -80000000 in hex. Length is 8 + negative sign is 9.
    // Min value is manually return. The longest value to be return automatically is -7FFFFFFF.
    // Length is 8 + negative sign is 9.
    var isNeg bool
	var t [9]byte 
	i := 9

    if (number < 0){
        isNeg = true
        if ( number == -2147483648 ){
            str = "-80000000"
            return
        }
        number = -number
    }

	for number >= 16 {
		i--
        t[i] = fnNumberToDigitMap[number & 15]    
		number >>= 4
	}

	i--
    t[i] = fnNumberToDigitMap[number]

	if isNeg {
		i--
		t[i] = '-'
	}

	str = string(t[i:])
	return
}

// UInt32ToHex returns a hexadecimal string representation of the input value of the uint32 data type.
// This function uses uppercase letters 'A' to 'F' for digit values >= 10.
func UInt32ToHex (number uint32) (str string) {
    // Max value FFFFFFFF. Length 8.
	var t [8]byte
	i := 8

	for number >= 16 {
		i--
        t[i] = fnNumberToDigitMap[number & 15]
		number >>= 4
	}

	i--
    t[i] = fnNumberToDigitMap[number]
     
	str = string(t[i:])
	return
}
