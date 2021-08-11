# Radix (Base) Conversion

Sequence of integers (`byte`, `int`) can be represent big number in positional numeral system with specified radix.

This project's functions convert input sequence from one radix to output sequence with other radix. Also encode/decode these sequences with specified alphabet.

> Aim of this project to implement universal radix converter. But implementation of `[]byte` input conversion using `big.Int` is more efficient than universtal implementation of `[]int` input conversion. Check benchmarks into `convert_test.go`.

Base58 encoding (used by Bitcoin and others) builds on this conversion of big number from base 256 to base 58.

### Examples

```Go
```

#### Base58

```Go
```
