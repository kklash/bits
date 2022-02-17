package bits

import (
	"fmt"
	"testing"
)

func TestBits_Or(t *testing.T) {
	test := func(a, b, expected string) {
		bitsA, _ := StringToBits(a)
		bitsB, _ := StringToBits(b)

		if actual := bitsA.Or(bitsB).String(); actual != expected {
			t.Errorf("Failed to perform bitwise OR\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test("00010", "00000", "00010")
	test("1111", "1000", "1111")
	test("00000000", "00000000", "00000000")
}

func TestBits_And(t *testing.T) {
	test := func(a, b, expected string) {
		bitsA, _ := StringToBits(a)
		bitsB, _ := StringToBits(b)

		if actual := bitsA.And(bitsB).String(); actual != expected {
			t.Errorf("Failed to perform bitwise OR\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test("001", "110", "000")
	test("111111", "111000", "111000")
}

func TestBits_Xor(t *testing.T) {
	test := func(a, b, expected string) {
		bitsA, _ := StringToBits(a)
		bitsB, _ := StringToBits(b)

		if actual := bitsA.Xor(bitsB).String(); actual != expected {
			t.Errorf("Failed to perform bitwise OR\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test("1111", "1111", "0000")
	test("0000", "1111", "1111")
	test("110011", "000000", "110011")
}

func _TestBitwiseOrBenchmark(t *testing.T) {
	res1 := testing.Benchmark(BenchmarkBitwiseOr)
	fmt.Println("Bitwise OR benchmark:")
	fmt.Printf("- %d ns/op\n", res1.NsPerOp())
	fmt.Printf("- %s\n", res1.T)
}

func BenchmarkBitwiseOr(b *testing.B) {
	bits := Zeros(256)
	for i := 0; i < b.N; i++ {
		bits = bits.Or(bits)
	}
}
