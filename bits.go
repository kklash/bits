// Package bits provides type declarations for working with arrays of binary bits.
package bits

// Bits represents a slice of binary numbers (ones and zeros). It has
// methods to manipulate and transform itself and other Bit slices using
// common bitwise operations. All Bits methods are big-endian.
type Bits []Bit

func nBits(b Bit, n uint32) Bits {
	bits := make(Bits, int(n))
	for i := 0; i < len(bits); i++ {
		bits[i] = b
	}
	return bits
}

// Zeros returns a Bit slice of length n composed entirely of zeros.
func Zeros(n uint32) Bits {
	return nBits(ZERO, n)
}

// Ones returns a Bit slice of length n composed entirely of ones.
func Ones(n uint32) Bits {
	return nBits(ONE, n)
}

// BitSize returns len(bits) as a uint32.
func (bits Bits) BitSize() uint32 {
	return uint32(len(bits))
}
