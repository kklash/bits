package bits

func (a Bits) bitwiseMap(b Bits, op func(Bit, Bit) Bit) Bits {
	if len(a) != len(b) {
		panic("Cannot performing bitwise operation on Bit slices of different lengths")
	}

	c := make(Bits, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = op(a[i], b[i])
	}

	return c
}

// Or performs returns bitsA | bitsB.
// Both Bit slices must be of equal length, or the call will panic.
func (bitsA Bits) Or(bitsB Bits) Bits {
	return bitsA.bitwiseMap(bitsB, func(a, b Bit) Bit {
		return a.Or(b)
	})
}

// And performs returns bitsA & bitsB.
// Both Bit slices must be of equal length, or the call will panic.
func (bitsA Bits) And(bitsB Bits) Bits {
	return bitsA.bitwiseMap(bitsB, func(a, b Bit) Bit {
		return a.And(b)
	})
}

// Xor performs returns bitsA ^ bitsB.
// Both Bit slices must be of equal length, or the call will panic.
func (bitsA Bits) Xor(bitsB Bits) Bits {
	return bitsA.bitwiseMap(bitsB, func(a, b Bit) Bit {
		return a.Xor(b)
	})
}
