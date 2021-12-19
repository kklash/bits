package bits

// ByteLength returns the number of bytes required to store bitSize Bits.
// If len(bits) % 8 != 0, it allows one extra byte to hold the overflow.
func ByteLength(bitSize uint32) uint32 {
	n := bitSize / 8
	if bitSize%8 != 0 {
		n += 1
	}
	return n
}

// ByteLength returns the number of bytes required to store len(bits) Bits.
// If len(bits) % 8 != 0, it allows one extra byte to hold the overflow.
func (bits Bits) ByteLength() uint32 {
	return ByteLength(uint32(len(bits)))
}
