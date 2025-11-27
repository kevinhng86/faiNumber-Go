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

// OctalToInt32 parse the input string as a signed octal integer string to an int32 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that OctalToInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.OctalToInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid octal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int32 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func OctalToInt32 (str string) (int32, error) {
    var result int32 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.OctalToInt32", str)
    }

    /* OctalToInt32 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The smallest octal number that can be parse for an int32 is -20000000000.
        // The length is 11 without the negative sign.  
        if ( runlen > 10 ) {
            if ( runlen > 11 ) {
                for ; start < length; start++{
                    if ( str[start] ^ 48 > 7 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.OctalToInt32", str)
            }
            result = int32(str[start] ^ 48)
            if ( result > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
            }
            start++
            // If the it is length 11 and the first digit is larger than two
            // then it is an underflow. If it is two than the rest of the string must be zero
            // to be the smallest number otherwise it is an underflow error. If the first digit is
            // a one than the rest of the string can be anything and it won't underflow because
            // the next smallest digit is -17777777777.           
            if ( result > 1 ) {
                if ( result > 2 ) {
                    for ; start < length; start++ {
                        if ( str[start] ^ 48 > 7 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.OctalToInt32", str)
                }

                var underflow bool
                // checking to see if the rest of the digit are 0 or invalid format.
                for ; start < length; start++ {
                    if ( str[start] ^ 48 > 7 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                    }
                    if ( str[start] != '0' ) {
                        underflow = true
                        start++
                        break
                    }
                }
                // if one of the digit is larger than a zero than check the rest of the digit
                // to see if it is a valid octal string or invalid format.
                if ( underflow == true ) {
                    for ; start < length; start++ {
                        if ( str[start] ^ 48 > 7 ) {
                            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                        }
                    }
                    return 0, MakeErrorUnderflow("faiNumber.OctalToInt32", str)
                }

                // If first digit is 2 and the rest of the is 0 then return the smallest int32 value.
                return -2147483648, nil
            }

            // parsing for when length is 11 and the first digit is a 1.
            for ; start < length; start++ {
                c = int32(str[start] ^ 48)
                if ( c > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                }
                result = (result << 3) | c
            }

            return -result, nil
        }

        // parsing for when the length is 11 or below.
        for ; start < length; start++ {
            c = int32(str[start] ^ 48)
            if ( c > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
            }
            result = (result << 3) | c
        }
        return -result, nil
    }
    /* OctalToInt32 Negative Parsing End */

    /* OctalToInt32 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest octal number that can fit into an int32 is 17777777777.
    // That is length 11.
    if ( runlen > 10 ) {
        // If the input number have a length of larger than 11 than it is a overflow
        // error. we just checking all the digits to see if it is a valid string
        // of octal number.
        if ( runlen > 11 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToInt32", str)
        }
        
        result = int32(str[start] ^ 48)
        if ( result > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
        }
        start++
        // if the first digit is larger than a one than it is an overflow error.
        // just checking the string to see if it is a valid string of octal number.
        if ( result > 1 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToInt32", str)
        }

        // parsing for when the string is 11 in length and the first digit is a one.
        for ; start < length; start++ {
            c = int32(str[start] ^ 48)
            if ( c > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
            }
            result = (result << 3) | c
        }

        return result, nil
    }

    // parsing for when the string's length is 10 or below.
    for ; start < length; start++ {
        c = int32(str[start] ^ 48)
        if ( c > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt32", str)
        }
        result = (result << 3) | c
    }

    return result, nil
    /* OctalToInt32 Positive Parsing End */
}

// OctalToUInt32 parse the input string as an unsigned octal integer string to an uint32 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that OctalToUInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.OctalToUInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid octal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func OctalToUInt32 (str string) (uint32, error) {
    var result uint32 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.OctalToUInt32", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;
    // The largest octal number that can fit into a uint32 is 37777777777.
    // The length is 11. If length larger than 11 then it is an overflow.
    // If length is 11 than the largest first digit must be a three.
    if ( runlen > 10 ) {
        if ( runlen > 11 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToUInt32", str)
        }
        
        result = uint32(str[start] ^ 48)
        if ( result > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt32", str)
        }
        start++
        if ( result > 3 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToUInt32", str)
        }
        
        // parsing for when the first digit is a 3 or less and length is 11.
        for ; start < length; start++ {
            c = uint32(str[start] ^ 48)
            if ( c > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt32", str)
            }
            result = (result << 3) | c
        }
        
        return result, nil
    }

    // parsing for when the length is 10 or less.
    for ; start < length; start++ {
        c = uint32(str[start] ^ 48)
        if ( c > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt32", str)
        }
        result = (result << 3) | c
    }

    return result, nil
}

// OctalToInt64 parse the input string as a signed octal integer string to an int64 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that OctalToInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.OctalToInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid octal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int64 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func OctalToInt64 (str string) (int64, error) {
    var result int64 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.OctalToInt64", str)
    }

    /* OctalToInt64 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The smallest octal number that can be parse for an int64 is -1000000000000000000000.
        // The length is 22 without the negative sign.  
        if ( runlen > 21 ) {
            if ( runlen > 22 ) {
                for ; start < length; start++{
                    if ( str[start] ^ 48 > 7 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.OctalToInt64", str)
            }
            result = int64(str[start] ^ 48)
            if ( result > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
            }
            start++
            // If the it is length 22 and the first digit is larger than 1
            // then it is an underflow. If it is 1 than the rest of the string must be zero
            // to be the smallest number otherwise it is an underflow error. At length 22
            // there is only one number possible and that is the smallest number.
            if ( result > 1 ) {
                for ; start < length; start++ {
                    if ( str[start] ^ 48 > 7 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.OctalToInt64", str)
            }

            var underflow bool
            // checking to see if the rest of the digit are 0 or invalid format.
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
                }
                if ( str[start] != '0' ) {
                    underflow = true
                    start++
                    break
                }
            }
            // if one of the digit is larger than a zero than check the rest of the digit
            // to see if it is a valid octal string or invalid format.
            if ( underflow == true ) {
                for ; start < length; start++ {
                    if ( str[start] ^ 48 > 7 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.OctalToInt64", str)
            }

            // If first digit is 1 and the rest of the is 0 then return the smallest int64 value.
            return -9223372036854775808, nil
        }

        // parsing for when the length is 21 or below.
        for ; start < length; start++ {
            c = int64(str[start] ^ 48)
            if ( c > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
            }
            result = (result << 3) | c
        }
        return -result, nil
    }
    /* OctalToInt64 Negative Parsing End */

    /* OctalToInt64 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest octal number that can fit into an int64 is 777777777777777777777.
    // That is length 21.
    // If the input number have a length of larger than 21 than it is a overflow
    // error. we just checking all the digits to see if it is a valid string
    // of octal number.
    if ( runlen > 21 ) {
        for ; start < length; start++ {
            if ( str[start] ^ 48 > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.OctalToInt64", str)
    }
        
    // parsing for when the string's length is 21 or below.
    for ; start < length; start++ {
        c = int64(str[start] ^ 48)
        if ( c > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToInt64", str)
        }
        result = (result << 3) | c
    }

    return result, nil
    /* OctalToInt64 Positive Parsing End */
}

// OctalToUInt64 parse the input string as an unsigned octal integer string to an uint64 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that OctalToUInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.OctalToUInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid octal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func OctalToUInt64 (str string) (uint64, error) {
    var result uint64 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.OctalToUInt64", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;
    // The largest octal number that can fit into a uint64 is 1777777777777777777777.
    // The length is 22. If length larger than 22 then it is an overflow.
    // If length is 22 than the largest first digit must be a 1.
    if ( runlen > 21 ) {
        if ( runlen > 22 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToUInt64", str)
        }
        
        result = uint64(str[start] ^ 48)
        if ( result > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt64", str)
        }
        start++
        if ( result > 1 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 7 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.OctalToUInt64", str)
        }
        
        // parsing for when the first digit is a 1 or less and length is 22.
        for ; start < length; start++ {
            c = uint64(str[start] ^ 48)
            if ( c > 7 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt64", str)
            }
            result = (result << 3) | c
        }
        
        return result, nil
    }

    // parsing for when the length is 21 or less.
    for ; start < length; start++ {
        c = uint64(str[start] ^ 48)
        if ( c > 7 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.OctalToUInt64", str)
        }
        result = (result << 3) | c
    }

    return result, nil
}
