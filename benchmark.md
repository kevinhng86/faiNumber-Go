## Benchmark

How fast is faiNumber? faiNumber is really fast compared to the strconv
package.

Below are the benchmark results for faiNumber version 1.0.0 compared to
strconv.ParseInt, strconv.ParseUint, strconv.FormatInt, or strconv.FormatUint.

Each of faiNumber's string parsing functions and int to string functions
was tested against the appropriate function of strconv with the same feature
call. For example, faiNumber.DecToInt32 was tested against strconv.ParseInt(input, 10, 32),
faiNumber.DecToInt64 was tested against strconv.ParseInt(input, 10, 64) and
faiNumber.DecToUInt32 was tested against strconv.ParseUint(input, 10, 32).

The benchmark code was written using the "testing" package and included in
faiNumber test files.

For each benchmark, faiNumber's function and strconv's function parse
the same input.

###### goos: linux
###### goarch: amd64
###### cpu: AMD A10-7800 Radeon R7, 12 Compute Cores 4C+8G 

<br>

#### Decimal string to Int

|                      | Input: 1234567890  | Input: -1234567890 |
|----------------------|--------------------|--------------------|
| faiNumber.DecToInt32 | 37.76 ns/op        | 41.26 ns/op        |
| strconv.ParseInt     | 70.84 ns/op        | 72.35 ns/op        |

|                       | Input: 123456789  |
|-----------------------|-------------------|
| faiNumber.DecToUInt32 | 41.66 ns/op       |
| strconv.ParseUint     | 61.51 ns/op       |

|                      | Input: 3546874589456  | Input: -3546874589456 |
|----------------------|-----------------------|-----------------------|
| faiNumber.DecToInt64 | 48.77 ns/op           | 45.41 ns/op           |
| strconv.ParseInt     | 83.29 ns/op           | 79.19 ns/op           |

|                       | Input: 56478945612345  |
|-----------------------|------------------------|
| faiNumber.DecToUInt64 | 56.19 ns/op            |
| strconv.ParseUint     | 81.10 ns/op            |

#### Binary string to Int
             
|                         | Input: 10101110101111010101 | Input: -10101110101111010101 |
|-------------------------|-----------------------------|------------------------------|
| faiNumber.BinaryToInt32 | 60.65 ns/op                 | 58.82 ns/op                  |
| strconv.ParseInt        | 123.8 ns/op                 | 116.3 ns/op                  |

|                          | Input: 101010101010101001110 |
|--------------------------|------------------------------|
| faiNumber.BinaryToUInt32 | 65.04 ns/op                  |
| strconv.ParseUint        | 122.2 ns/op                  |

|                         | Input: 111111000000000000000111111110  | Input: -111111000000000000000111111110 |
|-------------------------|----------------------------------------|----------------------------------------|
| faiNumber.BinaryToInt64 | 86.07 ns/op                            | 90.40 ns/op                            |
| strconv.ParseInt        | 155.5 ns/op                            | 172.4 ns/op                            |

|                          | Input: 10101010101010101010101010 |
|--------------------------|-----------------------------------|
| faiNumber.BinaryToUInt64 | 79.41 ns/op                       |
| strconv.ParseUint        | 135.4 ns/op                       |

#### Octal string to Int
     
|                        | Input: 17777777777 | Input: -17777777777 |
|------------------------|--------------------|---------------------|
| faiNumber.OctalToInt32 | 39.85 ns/op        | 42.22 ns/op         |
| strconv.ParseInt       | 85.84 ns/op        | 92.53 ns/op         |

|                         | Input: 37777777777 |
|-------------------------|--------------------|
| faiNumber.OctalToUInt32 | 41.44 ns/op        |
| strconv.ParseUint       | 85.60 ns/op        |

|                        | Input: 777777777777777777777 | Input: -777777777777777777777 |
|------------------------|------------------------------|-------------------------------|
| faiNumber.OctalToInt64 | 57.03 ns/op                  | 57.71 ns/op                   |
| strconv.ParseInt       | 117.3 ns/op                  | 119.1 ns/op                   |

|                         | Input: 1777777777777777777777 |
|-------------------------|-------------------------------|
| faiNumber.OctalToUInt64 | 65.99 ns/op                   |
| strconv.ParseUint       | 125.7 ns/op                   |

#### Hex string to Int
                   
|                      | Input: ABCDEF | Input: -ABCDEF |
|----------------------|---------------|----------------|
| faiNumber.HexToInt32 | 36.48 ns/op   | 34.07 ns/op    |
| strconv.ParseInt     | 62.08 ns/op   | 62.95 ns/op    |

|                       | Input: 1A2B3C4D |
|-----------------------|-----------------|
| faiNumber.HexToUInt32 | 37.45 ns/op     |
| strconv.ParseUint     | 59.44 ns/op     |

|                      | Input: ABCDEFABCDEF  | Input: -ABCDEFABCDEF |
|----------------------|----------------------|----------------------|
| faiNumber.HexToInt64 | 44.19 ns/op          | 50.30 ns/op          |
| strconv.ParseInt     | 95.54 ns/op          | 85.81 ns/op          |

|                       | Input: 1A2B3C4D5E6F |
|-----------------------|---------------------|
| faiNumber.HexToUInt64 | 46.56 ns/op         |
| strconv.ParseUint     | 78.10 ns/op         |

#### Int to Decimal string

|                      | Input: 2147483647 | Input: -2147483647 |
|----------------------|-------------------|--------------------|
| faiNumber.Int32ToDec | 63.54 ns/op       | 64.61 ns/op        |
| strconv.FormatInt    | 68.32 ns/op       | 69.06 ns/op        |

|                       | Input: 4294967295 |
|-----------------------|--|
| faiNumber.UInt32ToDec | 58.19 ns/op |
| strconv.FormatUint    | 67.40 ns/op |

|                      | Input: 9223372036854775807 | Input: -9223372036854775807 |
|----------------------|----------------------------|-----------------------------|
| faiNumber.Int64ToDec | 84.04 ns/op                | 86.12 ns/op                 |
| strconv.FormatInt    | 91.64 ns/op                | 92.75 ns/op                 |

|                       | Input: 18446744073709551615 |
|-----------------------|-----------------------------|
| faiNumber.UInt64ToDec | 85.81 ns/op                 |
| strconv.FormatUint    | 92.71 ns/op                 |

#### Int to Binary string

|                         | Input: 2147483647 | Input: -2147483647 |
|-------------------------|-------------------|--------------------|
| faiNumber.Int32ToBinary | 96.32 ns/op       | 94.65 ns/op        |
| strconv.FormatInt       | 105.0 ns/op       | 105.4 ns/op        |

|                          | Input: 4294967295 |
|--------------------------|-------------------|
| faiNumber.UInt32ToBinary | 88.02 ns/op       |
| strconv.FormatUint       | 105.8 ns/op       |

|                         | Input: 9223372036854775807 | Input: -9223372036854775807 |
|-------------------------|----------------------------|-----------------------------|
| faiNumber.Int64ToBinary | 132.6 ns/op                | 133.2 ns/op                 |
| strconv.FormatInt       | 179.9 ns/op                | 180.5 ns/op                 |

|                          | Input: 18446744073709551615 |
|--------------------------|-----------------------------|
| faiNumber.UInt64ToBinary | 149.0 ns/op                 |
| strconv.FormatUint       | 183.1 ns/op                 |

#### Int to Octal string

|                        | Input: 2147483647 | Input: -2147483647 |
|------------------------|-------------------|--------------------|
| faiNumber.Int32ToOctal | 58.88 ns/op       | 60.53 ns/op        |
| strconv.FormatInt      | 64.98 ns/op       | 65.90 ns/op        |

|                         | Input: 4294967295  |
|-------------------------|--------------------|
| faiNumber.UInt32ToOctal | 61.87 ns/op        |
| strconv.FormatUint      | 78.22 ns/op        |

|                        | Input: 9223372036854775807 | Input: -9223372036854775807 |
|------------------------|----------------------------|-----------------------------|
| faiNumber.Int64ToOctal | 78.40 ns/op                | 76.80 ns/op                 |
| strconv.FormatInt      | 85.22 ns/op                | 86.52 ns/op                 |

|                         | Input: 18446744073709551615 |
|-------------------------|-----------------------------|
| faiNumber.UInt64ToOctal | 78.95 ns/op                 |
| strconv.FormatUint      | 90.42 ns/op                 |

#### Int to Hex string
                  
|                      | Input: 2147483647 | Input: -2147483647 |
|----------------------|-------------------|--------------------|
| faiNumber.Int32ToHex | 50.56 ns/op       | 52.45 ns/op        |
| strconv.FormatInt    | 63.68 ns/op       | 60.56 ns/op        |

|                       | Input: 4294967295 |
|-----------------------|-------------------|
| faiNumber.UInt32ToHex | 42.50 ns/op       |
| strconv.FormatUint    | 52.42 ns/op       |

|                      | Input: 9223372036854775807 | Input: -9223372036854775807 |
|----------------------|----------------------------|-----------------------------|
| faiNumber.Int64ToHex | 66.60 ns/op                | 68.92 ns/op                 |
| strconv.FormatInt    | 74.72 ns/op                | 78.18 ns/op                 |

|                       | Input: 18446744073709551615 |
|-----------------------|-----------------------------|
| faiNumber.UInt64ToHex | 65.94 ns/op                 |
| strconv.FormatUint    | 74.47 ns/op                 |

</div>






