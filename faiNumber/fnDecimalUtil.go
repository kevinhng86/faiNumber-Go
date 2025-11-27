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

// DecToInt32 parse the input string as a signed decimal integer string to an int32 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that DecToInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.DecToInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid decimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int32 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func DecToInt32 (str string) (int32, error) {
    var result int32 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.DecToInt32", str)
    }

    /* DecToInt32 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        if ( runlen > 9 ) {
            if ( runlen > 10 ) {
                for ; start < length; start++{
                    if ( str[start] ^ 48 > 9 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.DecToInt32", str)
            }
            result = int32(str[start] ^ 48)
            if ( result > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
            }
            start++           
            if ( result > 2 ) {
                for ; start < length; start++ {
                    if ( str[start] ^ 48 > 9 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.DecToInt32", str)
            }

            for ; start < length - 1; start++ {
                c = int32(str[start] ^ 48)
                if ( c > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
                }
                result = (result << 1) + (result << 3) + c
            }


            c = int32(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
            }

            if ( result > 214748363 ) {
                if ( result > 214748364 ){
                    return 0, MakeErrorUnderflow("faiNumber.DecToInt32", str)
                }
                if ( c > 7 ) {
                    if ( c > 8 ) {
                        return 0, MakeErrorUnderflow("faiNumber.DecToInt32", str)
                    }
                    return -2147483648, nil
                }
            }

            result = (result << 1) + (result << 3) + c
            return -result, nil
        }

        for ; start < length; start++ {
            c = int32(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
            }
            result = (result << 1) + (result << 3) + c
        }
        return -result, nil
    }
    /* DecToInt32 Negative Parsing End */

    /* DecToInt32 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    if ( runlen > 9 ) {
        if ( runlen > 10 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToInt32", str)
        }
        result = int32(str[start] ^ 48)
        if ( result > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
        }
        start++
        if ( result > 2 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToInt32", str)
        }

        for ; start < length - 1; start++ {
            c = int32(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
            }
            result = (result << 1) + (result << 3) + c
        }


        c = int32(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
        }

        if ( result > 214748363 ) {
            if ( result > 214748364 ){
                return 0, MakeErrorOverflow("faiNumber.DecToInt32", str)
            }
            if ( c > 7 ) {
                return 0, MakeErrorOverflow("faiNumber.DecToInt32", str)
            }
        }

        result = (result << 1) + (result << 3) + c
        return result, nil
    }

    for ; start < length; start++ {
        c = int32(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt32", str)
        }
        result = (result << 1) + (result << 3) + c
    }

    return result, nil
    /* DecToInt32 Positive Parsing End */
}

// DecToUInt32 parse the input string as an unsigned decimal integer string to an uint32 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that DecToUInt32 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.DecToUInt32", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid decimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint32 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func DecToUInt32 (str string) (uint32, error) {
    var result uint32 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint32
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.DecToUInt32", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    if ( runlen > 9 ) {
        if ( runlen > 10 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToUInt32", str)
        }
        result = uint32(str[start] ^ 48)
        if ( result > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
        }
        start++
        if ( result > 4 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToUInt32", str)
        }

        for ; start < length - 1; start++ {
            c = uint32(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
            }
            result = (result << 1) + (result << 3) + c
        }

        c = uint32(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
        }

        if ( result > 429496728 ) {
            if ( result > 429496729 ){
                return 0, MakeErrorOverflow("faiNumber.DecToUInt32", str)
            }
            if ( c > 5 ) {
                return 0, MakeErrorOverflow("faiNumber.DecToUInt32", str)
            }
        }

        result = (result << 1) + (result << 3) + c
        return result, nil
    }

    for ; start < length; start++ {
        c = uint32(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt32", str)
        }
        result = (result << 1) + (result << 3) + c
    }

    return result, nil
}

// DecToInt64 parse the input string as a signed decimal integer string to an int64 value.
// The input string can begin with a leading sign: "+" or "-" to determine whether the input string represent a positive
// or negative value.
//
// The errors that DecToInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.DecToInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid decimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is smaller than the minimum value an int64 can hold then Error Code 3 is return.
// faiNumber classified this and an Underflow Error.
//
// If the value corresponding to str is larger than the maximum value an int64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func DecToInt64 (str string) (int64, error) {
    var result int64 = 0
    var length int = len(str)
    var runlen int
    var start int
    var c int64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.DecToInt64", str)
    }

    /* DecToInt64 Negative Parsing Start */
    if ( str[0] == '-' ) {
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
        }
        start = 1

        for ; start < length && str[start] == 48 ; start++ {}

        runlen = length - start;

        // Minimum value -9223372036854775808. Length 19.
        if ( runlen > 18 ) {
            if ( runlen > 19 ) {
                for ; start < length; start++ {
                    if ( str[start] ^ 48 > 9 ) {
                        return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
                    }
                }
                return 0, MakeErrorUnderflow("faiNumber.DecToInt64", str)
            }
            result = int64(str[start] ^ 48)
            if ( result > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
            }
            start++

            for ; start < length - 1; start++ {
                c = int64(str[start] ^ 48)
                if ( c > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
                }
                result = (result << 1) + (result << 3) + c
            }


            c = int64(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
            }

            if ( result > 922337203685477579 ) {
                if ( result > 922337203685477580 ){
                    return 0, MakeErrorUnderflow("faiNumber.DecToInt64", str)
                }
                if ( c > 7 ) {
                    if ( c > 8 ) {
                        return 0, MakeErrorUnderflow("faiNumber.DecToInt64", str)
                    }
                    return -9223372036854775808, nil
                }
            }

            result = (result << 1) + (result << 3) + c
            return -result, nil
        }

        for ; start < length; start++ {
            c = int64(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
            }
            result = (result << 1) + (result << 3) + c
        }
        return -result, nil
    }
    /* DecToInt64 Negative Parsing End */

    /* DecToInt64 Positive Parsing Start */
    if ( str[0] == '+' ){
        if ( length == 1 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
        }
        start = 1
    } else {
        start = 0
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // Maximum value 9223372036854775807. Length 19.
    if ( runlen > 18 ) {
        if ( runlen > 19 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToInt64", str)
        }
        result = int64(str[start] ^ 48)
        if ( result > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
        }
        start++

        for ; start < length - 1; start++ {
            c = int64(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
            }
            result = (result << 1) + (result << 3) + c
        }


        c = int64(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
        }

        if ( result > 922337203685477579 ) {
            if ( result > 922337203685477580 ){
                return 0, MakeErrorOverflow("faiNumber.DecToInt64", str)
            }
            if ( c > 7 ) {
                return 0, MakeErrorOverflow("faiNumber.DecToInt64", str)
            }
        }

        result = (result << 1) + (result << 3) + c
        return result, nil
    }

    for ; start < length; start++ {
        c = int64(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToInt64", str)
        }
        result = (result << 1) + (result << 3) + c
    }

    return result, nil
    /* DecToInt64 Positive Parsing End */
}

// DecToUInt64 parse the input string as an unsigned decimal integer string to an uint64 value.
// Sign parsing is not permitted. If there is a sign prefix then the function will yield an Invalid Format Error.
//
// The errors that DecToUInt64 returns are type [*FNError] and include FNError.Input = str,
// FNError.FromFunction = "faiNumber.DecToUInt64", FNError.Message = Error message, and FNError.Code = The error code.
// The error code and error message is different depend on which type of error.
//
// If str is empty then Error Code 1 is return. faiNumber classified this as an Empty String Error.
//
// If str contains invalid decimal digits then Error Code 2 is return. faiNumber classfied this
// as an Invalid Format Error.
//
// If the value corresponding to str is larger than the maximum value an uint64 can hold then Error Code 4 is return.
// faiNumber classfied this as an Overflow Error.
func DecToUInt64 (str string) (uint64, error) {
    var result uint64 = 0
    var length int = len(str)
    var runlen int
    var start int = 0
    var c uint64
     
    if ( length == 0 ) {
        return 0, MakeErrorEmptyString("faiNumber.DecToUInt64", str)
    }

    for ; start < length && str[start] == 48 ; start++ {}

    runlen = length - start;

    // Maximum value 18446744073709551615. Length 20.
    if ( runlen > 19 ) {
        if ( runlen > 20 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToUInt64", str)
        }
        result = uint64(str[start] ^ 48)
        if ( result > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
        }
        start++
        if ( result > 1 ) {
            for ; start < length; start++ {
                if ( str[start] ^ 48 > 9 ) {
                    return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
                }
            }
            return 0, MakeErrorOverflow("faiNumber.DecToUInt64", str)
        }

        for ; start < length - 1; start++ {
            c = uint64(str[start] ^ 48)
            if ( c > 9 ) {
                return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
            }
            result = (result << 1) + (result << 3) + c
        }

        c = uint64(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
        }

        if ( result > 1844674407370955160 ) {
            if ( result > 1844674407370955161 ){
                return 0, MakeErrorOverflow("faiNumber.DecToUInt64", str)
            }
            if ( c > 5 ) {
                return 0, MakeErrorOverflow("faiNumber.DecToUInt64", str)
            }
        }

        result = (result << 1) + (result << 3) + c
        return result, nil
    }

    for ; start < length; start++ {
        c = uint64(str[start] ^ 48)
        if ( c > 9 ) {
            return 0, MakeErrorInvalidFormat("faiNumber.DecToUInt64", str)
        }
        result = (result << 1) + (result << 3) + c
    }

    return result, nil
}
