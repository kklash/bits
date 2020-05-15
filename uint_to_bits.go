package bits

import "reflect"

// UintToBits converts a unsigned integer to bits (big-endian). The size of the resulting Bit
// slice will depend on the bit size of n, as reported by the reflect package. Panics if n is
// not a uint, uint8, uint16, uint32, or uint64.
func UintToBits(n interface{}) Bits {
	nValue := reflect.ValueOf(n)
	size := nValue.Type().Bits()
	bits := make(Bits, size)

	v := nValue.Uint()
	for j := size - 1; j >= 0; j-- {
		bits[j] = v&1 == 1
		v >>= 1
	}

	return bits
}
