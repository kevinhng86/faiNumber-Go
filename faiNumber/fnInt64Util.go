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

// Int64ToDec returns a decimal string representation of the input value of the int64 data type.
func Int64ToDec (number int64) (str string) {
    // Min value -9223372036854775808. Length 19 + negative sign is 20.
    var isNeg bool
	var t [20]byte 
	i := 19

    if (number < 0){
        isNeg = true
        if ( number == -9223372036854775808 ){
            str = "-9223372036854775808"
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

// UInt64ToDec returns a decimal string representation of the input value of the uint64 data type.
func UInt64ToDec (number uint64) (str string) {
    //  Max value 18446744073709551615. Length 20.
    var t [20]byte 
	i := 19

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

// Int64ToBinary returns a binary string representation of the input value of the int64 data type.
func Int64ToBinary (number int64) (str string) {
    // Min value -1000000000000000000000000000000000000000000000000000000000000000 in binary.
    // Length 64 + negative sign is 65. Min value is manually return the length of the longest
    // value is 63 + negative sign is 64.
    var isNeg bool
	var t [64]byte 
	i := 64

    if (number < 0){
        isNeg = true
        if ( number == -9223372036854775808 ){
            str = "-1000000000000000000000000000000000000000000000000000000000000000"
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

// UInt64ToBinary returns a binary string representation of the input value of the uint64 data type.
func UInt64ToBinary (number uint64) (str string) {
    // Max value 1111111111111111111111111111111111111111111111111111111111111111.
    // Length 64.
	var t [64]byte
	i := 64

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

// Int64ToBinaryAsUnsigned returns a binary string representation of the input value of the int64 type,
// This function treat the input value of type int64 as unsigned. This function returns the raw bit value
// of the input value of the int64 data type and it doesn't interprete negative values as negative values.
func Int64ToBinaryAsUnsigned (number int64) (str string) {
    str = UInt64ToBinary(uint64(number))
    return
}

// Int64ToOctal returns an octal string representation of the input value of the int64 data type.
func Int64ToOctal (number int64) (str string) {
    // Min value -1000000000000000000000 in octal. Length is 22 + negative sign is 23.
    // Min value is manually return. The longest value to be return automatically is -777777777777777777777.
    // Length is 21 + negative sign is 22.
    var isNeg bool
	var t [22]byte 
	i := 22

    if (number < 0){
        isNeg = true
        if ( number == -9223372036854775808 ){
            str = "-1000000000000000000000"
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

// UInt64ToOctal returns an octal string representation of the input value of the uint64 data type.
func UInt64ToOctal (number uint64) (str string) {
    // Max value 1777777777777777777777. Length 22.
	var t [22]byte
	i := 22

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

// Int64ToHex returns a hexadecimal string representation of the input value of the int64 data type.
// This function uses uppercase letters 'A' to 'F' for digit values >= 10.
func Int64ToHex (number int64) (str string) {
    // Min value -8000000000000000 in hex. Length is 16 + negative sign is 17.
    // Min value is manually return. The longest value to be return automatically is -7FFFFFFFFFFFFFFF.
    // Length is 16 + negative sign is 17.
    var isNeg bool
	var t [17]byte 
	i := 17

    if (number < 0){
        isNeg = true
        if ( number == -9223372036854775808 ){
            str = "-8000000000000000"
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

// UInt64ToHex returns a hexadecimal string representation of the input value of the uint64 data type.
// This function uses uppercase letters 'A' to 'F' for digit values >= 10.
func UInt64ToHex (number uint64) (str string) {
    // Max value FFFFFFFFFFFFFFFF. Length 16.
	var t [16]byte
	i := 16

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
