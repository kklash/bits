package bits

// Bit represents a single binary digit - a One or a Zero.
type Bit bool

const (
	ONE  Bit = true
	ZERO Bit = false
)

// Flip inverts the bit b and returns the result.
func (b Bit) Flip() Bit {
	return !b
}

// Or returns ONE if either b or c are 1. Otherwise it returns ZERO.
func (b Bit) Or(c Bit) Bit {
	return b || c
}

// Xor returns ONE if b != c. Otherwise it returns ZERO.
func (b Bit) Xor(c Bit) Bit {
	return b != c
}

// And returns ONE if both b and c are ONE. Otherwise it returns ZERO.
func (b Bit) And(c Bit) Bit {
	return b && c
}

// Int returns the bit b represented as an int.
func (b Bit) Int() int {
	return int(b.Uint())
}

// Uint returns the bit b represented as a uint.
func (b Bit) Uint() uint {
	if b {
		return 1
	}
	return 0
}

// String returns "1" if b is ONE, and "0" if b is ZERO.
func (b Bit) String() string {
	if b {
		return "1"
	}
	return "0"
}
