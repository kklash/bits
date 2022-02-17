# bits

Golang types for manipulating small-to-medium size arrays of binary bits with fine detail, with helpers for splitting, concatenation, shifting, padding, and conversion between other types.

See [the GoDoc](https://pkg.go.dev/github.com/kklash/bits) for full API documentation.

## Approach

A `Bit` is a boolean value with some helper methods:

```go
package bits // import "local/bits"

type Bit bool
    Bit represents a single binary digit - a One or a Zero.

const (
	ONE  Bit = true
	ZERO Bit = false
)

func (b Bit) And(c Bit) Bit
func (b Bit) Flip() Bit
func (b Bit) Int() int
func (b Bit) Or(c Bit) Bit
func (b Bit) String() string
func (b Bit) Uint() uint
func (b Bit) Xor(c Bit) Bit
```

A slice of these booleans is called `Bits`:

```go
type Bits []Bit
    Bits represents a slice of binary numbers (ones and zeros). It has methods
    to manipulate and transform itself and other Bit slices using common bitwise
    operations. All Bits methods are big-endian.

func BoolsToBits(bools []bool) Bits
func ByteToBits(b byte) Bits
func BytesToBits(bytes []byte) Bits
func Join(bitGroups []Bits) Bits
func Ones(n uint32) Bits
func StringToBits(bitString string) (Bits, error)
func UintToBits(n interface{}) Bits
func Zeros(n uint32) Bits
func (bitsA Bits) And(bitsB Bits) Bits
func (bits Bits) BigInt() *big.Int
func (bits Bits) BitSize() uint32
func (bits Bits) Bools() []bool
func (bits Bits) Byte() (b byte)
func (bits Bits) ByteLength() uint32
func (bits Bits) Bytes() []byte
func (bits Bits) LeadingZeros() uint32
func (bitsA Bits) Or(bitsB Bits) Bits
func (bits Bits) PadLeft(n uint32) Bits
func (bits Bits) PadRight(n uint32) Bits
func (bits Bits) ReadFrom(r io.Reader) (int64, error)
func (bits Bits) ShiftLeft(n uint32) Bits
func (bits Bits) ShiftRight(n uint32) Bits
func (bits Bits) Split(n uint32) []Bits
func (bits Bits) SplitPadLeft(n uint32) []Bits
func (bits Bits) SplitPadRight(n uint32) []Bits
func (bits Bits) String() string
func (bits Bits) Trim() Bits
func (bitsA Bits) Xor(bitsB Bits) Bits
```

## Examples

- [Splitting a bit slice into groups of 5 bits](https://github.com/kklash/bits/blob/8820e70ebaed215b79ed8ab1ecaabe6edeb548a1/split_test.go#L63-L71)
- [Left-shifting](https://github.com/kklash/bits/blob/8820e70ebaed215b79ed8ab1ecaabe6edeb548a1/shift_test.go#L59-L68) and [right-shifting](https://github.com/kklash/bits/blob/8820e70ebaed215b79ed8ab1ecaabe6edeb548a1/shift_test.go#L28-L37) bit slices
- [Padding a bit slice](https://github.com/kklash/bits/blob/8820e70ebaed215b79ed8ab1ecaabe6edeb548a1/pad_test.go#L30-L38)
- [XORing two bit slices](https://github.com/kklash/bits/blob/8820e70ebaed215b79ed8ab1ecaabe6edeb548a1/bitwise_test.go#L52-L60)
