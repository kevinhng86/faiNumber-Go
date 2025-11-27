# faiNumber-Go

Package faiNumber mainly deals with converting numerical strings to a value of
an int data type, converting a value of an int data type to a numerical string.
faiNumber is an extremely fast library for parsing string and converting int
values to strings. Be it whether if it is an overflow, underflow, empty strings
where it shouldn't be, or an invalid input, faiNumber functions were designed to
always return an error with a code when an error has occurred where it is possible
to have an error.

faiNumber is extremely fast. faiNumber may be the fastest golang library for numerical
string parsing and converting an int data type to a numerical string. faiNumber
functions were tested and is much faster than strconv package. But faiNumber only
support parsing decimal, binary, octal, hexadecimal string to int32, uint32, int64 or
uint64 data types. For converting an int value to a numerical string, faiNumber only support
convert int32, uint32, int64 or uint64 to a string of decimal, binary, octal or hexadecimal.

Each of faiNumber parsing function were written separately for each usage case and they were
optimize for speed. For converting to and from each data type of int32, uint32, int64, uint64,
a separate function was created to optimize for speed and ensure no type casting is required
for using any of faiNumber function.

## Benchmark

faiNumber's benchmark compare to strconv can be found here: <https://github.com/kevinhng86/faiNumber-Go/blob/main/benchmark.md>

## Parsing decimal strings to int data type

For parsing decimal strings to int faiNumber has the following functions:
DecToInt32, DecToUInt32, DecToInt64 and DecToUInt64.

## Parsing binary strings to int data type

For parsing decimal strings to int faiNumber has the following functions:
BinaryToInt32, BinaryToInt32AsUnsigned, BinaryToUInt32, BinaryToInt64,
BinaryToInt64AsUnsigned and BinaryToUInt64.

## Parsing octal strings to int data type

For parsing octal strings to int faiNumber has the following functions:
OctalToInt32, OctalToUInt32, OctalToInt64 and OctalToUInt64.

## Parsing hex strings to int data type

For parsing hex strings to int faiNumber has the following functions:
HexToInt32, HexToUInt32, HexToInt64, and HexToUInt64.

## Converting an int value to string

For converting an int value to string faiNumber has the following functions:
Int32ToDec, UInt32ToDec, Int32ToBinary, UInt32ToBinary, Int32ToBinaryAsUnsigned,
Int32ToOctal, UInt32ToOctal, Int32ToHex, UInt32ToHex, Int64ToDec, UInt64ToDec,
Int64ToBinary, UInt64ToBinary, Int64ToBinaryAsUnsigned, Int64ToOctal,
UInt64ToOctal, Int64ToHex, and UInt64ToHex.

## Documentation

faiNumber's documentation can be found here <https://lib.fai.host/go/github.com/kevinhng86/faiNumber-Go/faiNumber/>
