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

// BinaryToInt32 parse the input string as a signed binary integer string to an int32 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that BinaryToInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int32 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToInt32 (str string) (int32, error) {
    var result int32 = 0
    var length int = len(str)
    var runlen int
    var start int
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToInt32", str)
    }

    /* BinaryToInt32 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The longest and smallest negative number is -10000000000000000000000000000000. Which is length 32.
        // At length 32 there is only 1 valid number to parse.
        if ( runlen > 31 ) {
            if ( runlen > 32 ) {
                for ; start < length; start++{
                    if ( str[start] != '0' && str[start] != '1' ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.BinaryToInt32", str)
            }
            // Since we skip leading zero if the string is a valid binary string then
            // it can only be a one here. If it not a one then it not a valid format for binary string
            // an invalid format error is return.
            if ( str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
            }
            start++
            
            // Since there is only one number that can be parse at this length and it a leading
            // one with all zero we only have to check the entire string to see if it a zero.
            // If it contain at least one one then we have an underflow value. If it contain
            // anything besides a one and a zero then it a invalid format string.
            // If we hit a one then we break out of the loop and check for the rest of the digit
            // to see if it a valid format string of binary.
            var underflow bool
            for ; start < length; start++ {
                if ( str[start] != '0' ) {
                    if ( str[start] == '1' ) {
                        underflow = true
                        start++
                        break
                    }
                    return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
                }
            }

            if ( underflow == true ){
                for ; start < length; start++ {
                    if ( str[start] != '0' && str[start] != '1' ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.BinaryToInt32", str)
            }

            // If the string is a leading one with the rest of the digits are zero then return
            // the smallest value which is -2147483648.
            return -2147483648, nil
        }

        // Parsing for when the string's length is less than 32.
        for ; start < length; start++ {
            if ( str[start] == '1' ) {
                result = (result << 1) | 1
            } else if ( str[start] == '0' ) {
                result = result << 1
            } else {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
            }
        }
        
        return -result, nil
    }
    /* BinaryToInt32 Negative Parsing End */

    /* BinaryToInt32 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // For positive parsing, the longest possible binary string is 31 in length.
    // Anything over the length of 31 is overflow. If length is larger then 31 then
    // we just have to check all the digit to see if the input string is a valid
    // binary string or is it an invalid format.
    if ( runlen > 31 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToInt32", str)
    }

    // Parsing all string with a length of 31 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32", str)
        }
    }

    return result, nil
    /* BinaryToInt32 Positive Parsing End */
}

// BinaryToInt32AsUnsigned parse the input string as a binary integer string to an int32 value. This
// function treat the int32 data type as unsigned and parse the raw bit value of the input string without
// considering negative values. Sign parsing is not permitted. If there is a sign prefix then the
// function will yield an Invalid Format Error.
//
// The errors that BinaryToInt32AsUnsigned returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToInt32AsUnsigned", FNError.Message = Error message, and
// FNError.Code = The error code. The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the parsing length of str, excluding leading zeroes is larger than 32 then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToInt32AsUnsigned (str string) (int32, error) {
    var result int32 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToInt32AsUnsigned", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // If we parsing to an int32 as unsigned, we are parsing the raw bit value
    // of the binary string and it longest possible string is 32 in length.
    // Anything over 32 we just return an overflow error. Because we parsing as
    // unsigned we don't return underflow error. We do check the entire string to
    // see weather it is a valid binary string.
    if ( runlen > 32 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32AsUnsigned", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToInt32AsUnsigned", str)
    }

    // Parsing anything with a parse length of 32 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt32AsUnsigned", str)
        }
    }

    return result, nil
}

// BinaryToUInt32 parse the input string as an unsigned binary integer string to an uint32 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that BinaryToUInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToUInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToUInt32 (str string) (uint32, error) {
    var result uint32 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToUInt32", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest possible binary number that can fit in a uint32 is 32 in length.
    // Anything that longer than 32 digits is an overflow error. If it is overflow
    // then check the rest of the digit to see if it a valid binary string to
    // return an invalid format error if it is not a valid binary string. Or return
    // an overflow error if it is a valid binary string.
    if ( runlen > 32 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToUInt32", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToUInt32", str)
    }

    // Parsing anything with a parse length of 32 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToUInt32", str)
        }
    }

    return result, nil
}

// BinaryToInt64 parse the input string as a signed binary integer string to an int64 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that BinaryToInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int64 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToInt64 (str string) (int64, error) {
    var result int64 = 0
    var length int = len(str)
    var runlen int
    var start int
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToInt64", str)
    }

    /* BinaryToInt64 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // The longest and smallest negative number is
        // -1000000000000000000000000000000000000000000000000000000000000000.
        // Which is length 64. At length 64 there is only 1 valid number to parse.
        if ( runlen > 63 ) {
            if ( runlen > 64 ) {
                for ; start < length; start++{
                    if ( str[start] != '0' && str[start] != '1' ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.BinaryToInt64", str)
            }
            // Since we skip leading zero if the string is a valid binary string then
            // it can only be a one here. If it not a one then it not a valid format for binary string
            // an invalid format error is return.
            if ( str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
            }
            start++
            
            // Since there is only one number that can be parse at this length and it a leading
            // one with all zero we only have to check the entire string to see if it a zero.
            // If it contain at least one one then we have an underflow value. If it contain
            // anything besides a one and a zero then it a invalid format string.
            // If we hit a one then we break out of the loop and check for the rest of the digit
            // to see if it a valid format string of binary.
            var underflow bool
            for ; start < length; start++ {
                if ( str[start] != '0' ) {
                    if ( str[start] == '1' ) {
                        underflow = true
                        start++
                        break
                    }
                    return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
                }
            }

            if ( underflow == true ){
                for ; start < length; start++ {
                    if ( str[start] != '0' && str[start] != '1' ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.BinaryToInt64", str)
            }

            // If the string is a leading one with the rest of the digits are zero then return
            // the smallest value which is -9223372036854775808.
            return -9223372036854775808, nil
        }

        // Parsing for when the string's length is less than 64.
        for ; start < length; start++ {
            if ( str[start] == '1' ) {
                result = (result << 1) | 1
            } else if ( str[start] == '0' ) {
                result = result << 1
            } else {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
            }
        }
        
        return -result, nil
    }
    /* BinaryToInt64 Negative Parsing End */

    /* BinaryToInt64 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // For positive parsing, the longest possible binary string is 63 in length.
    // Anything over the length of 63 is overflow. If length is larger then 63
    // then we just have to check all the digit to see if the input string is
    // a valid binary string or is it an invalid format.
    if ( runlen > 63 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToInt64", str)
    }

    // Parsing all string with a length of 63 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64", str)
        }
    }

    return result, nil
    /* BinaryToInt64 Positive Parsing End */
}

// BinaryToInt64AsUnsigned parse the input string as a binary integer string to an int64 value. This
// function treat the int64 data type as unsigned and parse the raw bit value of the input string without
// considering negative values. Sign parsing is not permitted. If there is a sign prefix then the
// function will yield an Invalid Format Error.
//
// The errors that BinaryToInt64AsUnsigned returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToInt64AsUnsigned", FNError.Message = Error message, and
// FNError.Code = The error code. The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the parsing length of str, excluding leading zeroes is larger than 64 then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToInt64AsUnsigned (str string) (int64, error) {
    var result int64 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToInt64AsUnsigned", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // If we parsing to an int64 as unsigned, we are parsing the raw bit value
    // of the binary string and it longest possible string is 64 in length.
    // Anything over 64 we just return an overflow error. Because we parsing as
    // unsigned we don't return underflow error. We do check the entire string to
    // see weather it is a valid binary string.
    if ( runlen > 64 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64AsUnsigned", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToInt64AsUnsigned", str)
    }

    // Parsing anything with a parse length of 64 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToInt64AsUnsigned", str)
        }
    }

    return result, nil
}

// BinaryToUInt64 parse the input string as an unsigned binary integer string to an uint64 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that BinaryToUInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.BinaryToUInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid binary digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func BinaryToUInt64 (str string) (uint64, error) {
    var result uint64 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.BinaryToUInt64", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // The largest possible binary number that can fit in a uint64 is 64 in length.
    // Anything that longer than 64 digits is an overflow error. If it is overflow
    // then check the rest of the digit to see if it a valid binary string to
    // return an invalid format error if it is not a valid binary string. Or return
    // an overflow error if it is a valid binary string.
    if ( runlen > 64 ) {
        for ; start < length; start++{
            if ( str[start] != '0' && str[start] != '1' ) {
                return 0, MakeErrorInvalidFormat("faiNumber.BinaryToUInt64", str)
            }
        }
        return 0, MakeErrorOverflow("faiNumber.BinaryToUInt64", str)
    }

    // Parsing anything with a parse length of 64 or below.
    for ; start < length; start++ {
        if ( str[start] == '1' ) {
            result = (result << 1) | 1
        } else if ( str[start] == '0' ) {
            result = result << 1
        } else {
            return 0, MakeErrorInvalidFormat("faiNumber.BinaryToUInt64", str)
        }
    }

    return result, nil
}
