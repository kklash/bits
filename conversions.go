package bits

import "math/big"

// String calls Bit.String for every Bit in bits and concatenates the results.
// Note that leading zeros are not removed.
func (bits Bits) String() string {
	s := ""
	for _, bit := range bits {
		s += bit.String()
	}

	return s
}

// BigInt returns a pointer to a big.Int whose value is set
// to the value of bits as a big endian integer.
func (bits Bits) BigInt() *big.Int {
	i := new(big.Int)
	if len(bits) == 0 {
		return i
	}

	i.SetString(bits.String(), 2)
	return i
}

// Byte converts bits to a single byte. Panics if len(bits) != 8
func (bits Bits) Byte() (b byte) {
	if len(bits) != 8 {
		panic("Failed to convert Bits to byte; len(bits) must equal 8")
	}

	for j := 0; j < 8; j++ {
		if bits[j] {
			b += 1 << (7 - j)
		}
	}

	return
}

// Bytes converts bits to a slice of bytes. Panics if len(bits) % 8 != 0.
func (bits Bits) Bytes() []byte {
	if len(bits)%8 != 0 {
		panic("Failed to convert Bits to []byte; len(bits) must be divisible by 8")
	}

	bitGroups := bits.Split(8)
	byteSize := len(bitGroups)
	bytes := make([]byte, byteSize)

	for i := 0; i < byteSize; i++ {
		bytes[i] = bitGroups[i].Byte()
	}

	return bytes
}
