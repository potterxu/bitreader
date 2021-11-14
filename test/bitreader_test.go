package bitreader_test

import (
	"testing"

	"github.com/potterxu/bitreader"
)

var (
	// 00010111 10101001 01001011 11110101
	TEST_DATA = [...]byte{23, 169, 75, 245}
)

func TestReadBit(t *testing.T) {
	bitreader := bitreader.BitReader(TEST_DATA[:])

	expected := []bool{false, false, false, true}

	for i := 0; i < len(expected); i++ {
		actual, err := bitreader.ReadBit()
		if err != nil {
			t.Fatal(err)
		}
		if actual != expected[i] {
			t.Fatalf("Failed: Expected %t, Actual %t", expected, actual)
		}
	}
}

func TestReadBits(t *testing.T) {
	bitreader := bitreader.BitReader(TEST_DATA[:])
	expected := 6202671

	bitreader.SkipBits(3)
	actual, err := bitreader.ReadBits(23)

	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf("Failed: Expected %d, Actual %d", expected, actual)
	}
}
