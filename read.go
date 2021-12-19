package bits

import (
	"errors"
	"fmt"
	"io"
)

// ErrReadingIntoBits is returned as a wrapped error by Read when
// reading from the given reader fails.
var ErrReadingIntoBits = errors.New("Failed to generate random bits")

// Read fills bits with data generated by r until either r reaches EOF or bits is filled up.
// Useful for generating random data. Returns the number of bits copied and an error if one
// was encountered.
//
// Note than the returned error will wrap io.EOF if the reader returns EOF without having
// successfully read any bytes, UNLESS bits == nil || len(bits) == 0, in which case, the
// number of bits read will be zero, and error will be nil.
func Read(r io.Reader, bits Bits) (int, error) {
	var (
		err                 error
		bitsRead, bytesRead int
	)
	if r == nil {
		return bitsRead, fmt.Errorf("%w: reader is nil", ErrReadingIntoBits)
	}

	if bits == nil || len(bits) == 0 {
		return bitsRead, nil
	}

	byteBuf := make([]byte, bits.ByteLength())

	for err == nil && bytesRead < len(byteBuf) {
		var j int
		j, err = r.Read(byteBuf[bytesRead:])
		bytesRead += j
	}

	for i := 0; i < bytesRead; i++ {
		bitsRead += copy(bits[i*8:], ByteToBits(byteBuf[i]))
	}

	if err != nil {
		if errors.Is(err, io.EOF) {
			err = fmt.Errorf("%s: %w", ErrReadingIntoBits, err) // Wrap io.EOF instead as clear indicator that EOF was reached
			return bitsRead, err
		}

		return bitsRead, fmt.Errorf("%w: %s", ErrReadingIntoBits, err)
	}

	return bitsRead, nil
}

// ReadFrom implements the io.ReaderFrom interface, populating the Bit slice with data
// from r until bits is filled, or an error occurs. EOF is treated as a nil error.
func (bits Bits) ReadFrom(r io.Reader) (int64, error) {
	n, err := Read(r, bits)
	if errors.Is(err, io.EOF) {
		err = nil
	}
	x := ByteLength(uint32(n))
	return int64(x), err
}
