package bits

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func TestRead_File(t *testing.T) {
	file, err := os.Open("go.mod")
	if err != nil {
		t.Errorf("Failed to open go.mod: %s", err)
		return
	}
	defer file.Close()

	bits := make(Bits, 232)
	n, err := Read(file, bits)
	if err != nil {
		t.Errorf("Failed to read from file: %s", err)
		return
	}

	if n != 232 {
		t.Errorf("Failed to read expected number of bits\nWanted %d\nGot    %d", 232, n)
		return
	}

	// "module github.com/kklash/bits"
	expected := "011011010110111101100100011101010110110001100101001000000110011101101001" +
		"01110100011010000111010101100010001011100110001101101111011011010010111101101011011" +
		"01011011011000110000101110011011010000010111101100010011010010111010001110011"

	if actual := bits.String(); actual != expected {
		t.Errorf("Failed to get expected bits\nWanted %s\nGot    %s", expected, actual)
		return
	}
}

func TestRead_Buffer(t *testing.T) {
	rawBytes := []byte("I am a teapot")
	buf := bytes.NewBuffer(rawBytes)

	bits := make(Bits, len(rawBytes)*8-5)

	n, err := Read(buf, bits)
	if err != nil {
		t.Errorf("Failed to read from buffer: %s", err)
		return
	}

	if n != len(bits) {
		t.Errorf("Failed to read expected number of bits\nWanted %d\nGot    %d", len(bits), n)
		return
	}

	expected := "010010010010000001100001011011010010000001100001001000000111010001100101011000010111000001101111011"
	if bits.String() != expected {
		t.Errorf("Failed to get expected bits\nWanted %s\nGot    %s", expected, bits)
		return
	}
}

func TestRead_EmptyBuffer(t *testing.T) {
	buf := new(bytes.Buffer)
	bits := make(Bits, 8)
	n, err := Read(buf, bits)
	if n != 0 {
		t.Errorf("Expected zero bits to be read - got %d", n)
	} else if err == nil {
		t.Errorf("Expected error to be returned if bitsRead is zero")
	} else if !errors.Is(err, io.EOF) {
		t.Errorf("Failed to get expected EOF error\nWanted %s\nGot    %s", io.EOF, err)
	}
}

func TestRead_NilReader(t *testing.T) {
	bits := make(Bits, 8)
	_, err := Read(nil, bits)
	if !errors.Is(err, ErrReadingIntoBits) {
		t.Errorf("Failed to get expected error\nWanted %s\nGot    %s", ErrReadingIntoBits, err)
	}
}

func TestBits_ReadFrom(t *testing.T) {
	{
		bits := make(Bits, 0)
		buf := bytes.NewBufferString("a")
		n, err := bits.ReadFrom(buf)
		if err != nil {
			t.Errorf("Failed to call bits.ReadFrom(bytes.Buffer): %s", err)
		} else if n != 0 {
			t.Errorf("expected ReadFrom to read 0 bytes; got %d", n)
		}
	}

	{
		bits := make(Bits, 30)
		buf := bytes.NewBufferString("abcd")
		n, err := bits.ReadFrom(buf)
		if err != nil {
			t.Errorf("Failed to call bits.ReadFrom(bytes.Buffer): %s", err)
		} else if n != 4 {
			t.Errorf("expected ReadFrom to read 4 bytes; got %d", n)
		}

		expected := "011000010110001001100011011001"
		bitString := bits.String()
		if bitString != expected {
			t.Errorf("did not get expected bits with bits.ReadFrom\nWanted %s\nGot    %s", expected, bitString)
		}
	}

	{
		bits := make(Bits, 8)
		buf := bytes.NewBufferString("a")
		n, err := bits.ReadFrom(buf)
		if err != nil {
			t.Errorf("Failed to call bits.ReadFrom(bytes.Buffer): %s", err)
		} else if n != 1 {
			t.Errorf("expected ReadFrom to read 1 byte; got %d", n)
		}

		expected := "01100001"
		bitString := bits.String()
		if bitString != expected {
			t.Errorf("did not get expected bits with bits.ReadFrom\nWanted %s\nGot    %s", expected, bitString)
		}
	}

	{
		bits := make(Bits, 10)
		buf := bytes.NewBufferString("a")
		n, err := bits.ReadFrom(buf)
		if err != nil {
			t.Errorf("Failed to call bits.ReadFrom(bytes.Buffer): %s", err)
		} else if n != 1 {
			t.Errorf("expected ReadFrom to read 1 byte; got %d", n)
		}

		expected := "0110000100"
		bitString := bits.String()
		if bitString != expected {
			t.Errorf("did not get expected bits with bits.ReadFrom\nWanted %s\nGot    %s", expected, bitString)
		}
	}
}
