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

import "fmt"

// FNError is use to record errors when they occur. A pointer to a new instance of FNError is returned whenever
// an error occur.
type FNError struct {
    FromFunction string // The function where the error is produced e.g DecToInt32, DecToUInt32, DecToInt64, DecToUInt64.
    Input string // The input string.
    Message string // The error message.
    Code int // The error code.
}

// An implement of the Error() method for *FNError. A string representation of the error is return.
func (e *FNError) Error() string {
	return fmt.Sprintf("Error from function %s. Error code: %d. Input value: \"%s\". Message: %s." , e.FromFunction, e.Code, e.Input, e.Message)
}

// This function make error Empty String whenever needed.
// Error Empty String is for when the input string is empty.
//
// faiNumber use error code 1 for all Empty String errors.
//
// This function return a pointer to a new instance of [FNError]
// with the FNError.Message = "Empty string error" and the FNError.Code = 1.
// The FunctionName is use for the name of the function where the error occur.
// The input is the input that produced the error.
func MakeErrorEmptyString(FunctionName string, input string) error {
	err := FNError{
		FromFunction: FunctionName,
		Input:  input,
        Message: "Empty string error",
        Code: 1,
	}
	return &err
}

// This function make error Invalid Format whenever needed.
// Error Invalid Format indicates that the parsing value is not the right
// format for the data type. e.g the input string is not a valid decimal number
// string when calling the function [DecToInt32]. All of faiNumber string parsing
// functions always check the input string to see whether if it is a valid numerical string
// type for the calling function even if an underflow or overflow error occur first.
// If the input string is not a valid string of the numerical type that the function required then
// the function will return error Invalid Format even if error overflow or error underflow
// occur first.
//
// faiNumber use error code 2 for all Invalid Format errors.
//
// This function return a pointer to a new instance of [FNError]
// with the FNError.Message = "Invalid format error" and the FNError.Code = 2.
// The FunctionName is use for the name of the function where the error occur.
// The input is the input that produced the error.
func MakeErrorInvalidFormat(FunctionName string, input string) error {
	err := FNError{
		FromFunction: FunctionName,
		Input:  input,
        Message: "Invalid format error",
        Code: 2,
	}
	return &err
}

// This function make error Underflow whenever needed.
// Error Underflow is produced when the input value is smaller
// than the smallest negative value that can be parse to a signed data type.
//
// faiNumber use error code 3 for all Underflow errors.
//
// This function return a pointer to a new instance of [FNError]
// with the FNError.Message = "Underflow error" and the FNError.Code = 3.
// The FunctionName is use for the name of the function where the error occur.
// The input is the input that produced the error.
func MakeErrorUnderflow(FunctionName string, input string) error {
	err := FNError{
		FromFunction: FunctionName,
		Input:  input,
        Message: "Underflow error",
        Code: 3,
	}
	return &err
}

// This function make error Overflow whenever needed.
// Error Overflow is produced when the input value is larger
// than the largest value that can be parsed to a data type. For signed type
// the value is larger than the largest positive value for the data type.
//
// faiNumber use error code 4 for all Overflow errors.
//
// This function return a pointer to a new instance of [FNError]
// with the FNError.Message = "Overflow error" and the FNError.Code = 4.
// The FunctionName is use for the name of the function where the error occur.
// The input is the input that produced the error.
func MakeErrorOverflow(FunctionName string, input string) error {
	err := FNError{
		FromFunction: FunctionName,
		Input:  input,
        Message: "Overflow error",
        Code: 4,
	}
	return &err
}
