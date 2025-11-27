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

// This is a map of ascii code to digit value. '0' - '9' ascii will map to 0 - 9.
// 'A' - 'Z' will map to 10 - 35. 'a' to 'z' will map to 10 - 35.
// Any other ascii value will map to 255. This is for fast conversion of ascii digit value
// to a number value. This map is local.
const fnDigitMap string = "\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\x0A\x0B\x0C\x0D\x0E\x0F\x10\x11\x12\x13\x14" +
    "\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F\x20\x21\x22\x23" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\x0A\x0B\x0C\x0D\x0E\x0F\x10\x11\x12" +
    "\x13\x14\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F\x20\x21" +
    "\x22\x23\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" +
    "\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF" 

// HexToInt32 parse the input string as a signed hexadecimal integer string to an int32 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that HexToInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.HexToInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid hexadecimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int32 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func HexToInt32 (str string) (int32, error) {
    var result int32 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.HexToInt32", str)
    }

    /* HexToInt32 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The smallest hex number that can be parse for an int32 is -80000000.
        // The length is 8 without the negative sign.  
        if ( runlen > 7 ) {
            if ( runlen > 8 ) {
                for ; start < length; start++{
                    if ( fnDigitMap[str[start]] > 15 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.HexToInt32", str)
            }
            result = int32(fnDigitMap[str[start]])
            if ( result > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
            }
            start++
            // If it is length 8 and the first digit is larger than 8
            // then it is an underflow. If it is 8 than the rest of the string must be zero
            // to be the smallest number otherwise it is an underflow error. If the first digit is
            // a 7 or below than the rest of the string can be anything and it won't underflow because
            // the next smallest digit is -7FFFFFFF.           
            if ( result > 7 ) {
                if ( result > 8 ) {
                    for ; start < length; start++ {
                        if ( fnDigitMap[str[start]] > 15 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.HexToInt32", str)
                }

                var underflow bool
                // checking to see if the rest of the digit are 0 or invalid format.
                for ; start < length; start++ {
                    if ( fnDigitMap[str[start]] > 15 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                    }
                    if ( fnDigitMap[str[start]] > 0 ) {
                        underflow = true
                        start++
                        break
                    }
                }
                // if one of the digit is larger than a zero than check the rest of the digit
                // to see if it is a valid hex string or invalid format.
                if ( underflow == true ) {
                    for ; start < length; start++ {
                        if ( fnDigitMap[str[start]] > 15 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.HexToInt32", str)
                }

                // If first digit is 8 and the rest of the is 0 then return the smallest int32 value.
                return -2147483648, nil
            }

            // parsing for when length is 8 and the first digit is a 7 or below.
            for ; start < length; start++ {
                c = int32(fnDigitMap[str[start]])
                if ( c > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                }
                result = (result << 4) | c
            }

            return -result, nil
        }

        // parsing for when the length is 7 or below.
        for ; start < length; start++ {
            c = int32(fnDigitMap[str[start]])
            if ( c > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
            }
            result = (result << 4) | c
        }
        return -result, nil
    }
    /* HexToInt32 Negative Parsing End */

    /* HexToInt32 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest hex number that can fit into an int32 is 7FFFFFFF.
    // That is length 8.
    if ( runlen > 7 ) {
        // If the input number have a length of larger than 8 than it is a overflow
        // error. we just checking all the digits to see if it is a valid string
        // of hex number.
        if ( runlen > 8 ) {
            for ; start < length; start++ {
                if ( fnDigitMap[str[start]] > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.HexToInt32", str)
        }
        
        result = int32(fnDigitMap[str[start]])
        if ( result > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
        }
        start++
        // if the first digit is larger than a 7 than it is an overflow error.
        // just checking the string to see if it is a valid string of hex number.
        if ( result > 7 ) {
            for ; start < length; start++ {
                if ( fnDigitMap[str[start]] > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.HexToInt32", str)
        }

        // parsing for when the string is 8 in length and the first digit is a 7 or below.
        for ; start < length; start++ {
            c = int32(fnDigitMap[str[start]])
            if ( c > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
            }
            result = (result << 4) | c
        }

        return result, nil
    }

    // parsing for when the string's length is 7 or below.
    for ; start < length; start++ {
        c = int32(fnDigitMap[str[start]])
        if ( c > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt32", str)
        }
        result = (result << 4) | c
    }

    return result, nil
    /* HexToInt32 Positive Parsing End */
}

// HexToUInt32 parse the input string as an unsigned hexadecimal integer string to an uint32 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that HexToUInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.HexToUInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid hexadecimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func HexToUInt32 (str string) (uint32, error) {
    var result uint32 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.HexToUInt32", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;
    // The largest hex number that can fit into a uint32 is FFFFFFFF.
    // The length is 8. If length larger than 8 then it is an overflow.
    if ( runlen > 8 ) {
        for ; start < length; start++ {
            if ( fnDigitMap[str[start]] > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToUInt32", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.HexToUInt32", str)
    }

    // parsing for when the length is 8 or less.
    for ; start < length; start++ {
        c = uint32(fnDigitMap[str[start]])
        if ( c > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToUInt32", str)
        }
        result = (result << 4) | c
    }

    return result, nil
}

// HexToInt64 parse the input string as a signed hexadecimal integer string to an int64 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that HexToInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.HexToInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid hexadecimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int64 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func HexToInt64 (str string) (int64, error) {
    var result int64 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.HexToInt64", str)
    }

    /* HexToInt64 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The smallest hex number that can be parse for an int64 is -8000000000000000.
        // The length is 16 without the negative sign.  
        if ( runlen > 15 ) {
            if ( runlen > 16 ) {
                for ; start < length; start++{
                    if ( fnDigitMap[str[start]] > 15 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.HexToInt64", str)
            }
            result = int64(fnDigitMap[str[start]])
            if ( result > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
            }
            start++
            // If it is length 16 and the first digit is larger than 8
            // then it is an underflow. If it is 8 than the rest of the string must be zero
            // to be the smallest number otherwise it is an underflow error. If the first digit is
            // a 7 or below than the rest of the string can be anything and it won't underflow because
            // the next smallest digit is -7FFFFFFFFFFFFFFF.           
            if ( result > 7 ) {
                if ( result > 8 ) {
                    for ; start < length; start++ {
                        if ( fnDigitMap[str[start]] > 15 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.HexToInt64", str)
                }

                var underflow bool
                // checking to see if the rest of the digit are 0 or invalid format.
                for ; start < length; start++ {
                    if ( fnDigitMap[str[start]] > 15 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                    }
                    if ( fnDigitMap[str[start]] > 0 ) {
                        underflow = true
                        start++
                        break
                    }
                }
                // if one of the digit is larger than a zero than check the rest of the digit
                // to see if it is a valid hex string or invalid format.
                if ( underflow == true ) {
                    for ; start < length; start++ {
                        if ( fnDigitMap[str[start]] > 15 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.HexToInt64", str)
                }

                // If first digit is 8 and the rest of the is 0 then return the smallest int64 value.
                return -9223372036854775808, nil
            }

            // parsing for when length is 16 and the first digit is a 7 or below.
            for ; start < length; start++ {
                c = int64(fnDigitMap[str[start]])
                if ( c > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                }
                result = (result << 4) | c
            }

            return -result, nil
        }

        // parsing for when the length is 15 or below.
        for ; start < length; start++ {
            c = int64(fnDigitMap[str[start]])
            if ( c > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
            }
            result = (result << 4) | c
        }
        return -result, nil
    }
    /* HexToInt64 Negative Parsing End */

    /* HexToInt64 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest hex number that can fit into an int64 is 7FFFFFFFFFFFFFFF.
    // That is length 16.
    if ( runlen > 15 ) {
        // If the input number have a length of larger than 16 than it is a overflow
        // error. we just checking all the digits to see if it is a valid string
        // of hex number.
        if ( runlen > 16 ) {
            for ; start < length; start++ {
                if ( fnDigitMap[str[start]] > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.HexToInt64", str)
        }
        
        result = int64(fnDigitMap[str[start]])
        if ( result > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
        }
        start++
        // if the first digit is larger than a 7 than it is an overflow error.
        // just checking the string to see if it is a valid string of hex number.
        if ( result > 7 ) {
            for ; start < length; start++ {
                if ( fnDigitMap[str[start]] > 15 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.HexToInt64", str)
        }

        // parsing for when the string is 16 in length and the first digit is a 7 or below.
        for ; start < length; start++ {
            c = int64(fnDigitMap[str[start]])
            if ( c > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
            }
            result = (result << 4) | c
        }

        return result, nil
    }

    // parsing for when the string's length is 15 or below.
    for ; start < length; start++ {
        c = int64(fnDigitMap[str[start]])
        if ( c > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToInt64", str)
        }
        result = (result << 4) | c
    }

    return result, nil
    /* HexToInt64 Positive Parsing End */
}

// HexToUInt64 parse the input string as an unsigned hexadecimal integer string to an uint64 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that HexToUInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.HexToUInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid hexadecimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func HexToUInt64 (str string) (uint64, error) {
    var result uint64 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.HexToUInt64", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;
    // The largest hex number that can fit into a uint64 is FFFFFFFFFFFFFFFF.
    // The length is 16. If length larger than 16 then it is an overflow.
    if ( runlen > 16 ) {
        for ; start < length; start++ {
            if ( fnDigitMap[str[start]] > 15 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.HexToUInt64", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.HexToUInt64", str)
    }

    // parsing for when the length is 16 or less.
    for ; start < length; start++ {
        c = uint64(fnDigitMap[str[start]])
        if ( c > 15 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.HexToUInt64", str)
        }
        result = (result << 4) | c
    }

    return result, nil
}
