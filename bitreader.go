package bitreader

import (
	"fmt"
)

const (
	BYTE int = 8
)

type BitReaderType struct {
	data []byte

	base   int
	offset int
}

func BitReader(data []byte) *BitReaderType {
	return &BitReaderType{
		data:   data,
		base:   0,
		offset: 0,
	}
}

func (reader *BitReaderType) SkipBytes(bytes int) error {
	return reader.SkipBits(bytes * BYTE)
}

func (reader *BitReaderType) SkipBits(bits int) error {
	err := reader.checkBits(bits)
	if err != nil {
		return err
	}
	bits += reader.offset
	reader.base += bits / BYTE
	reader.offset += bits % BYTE
	return nil
}

func (reader *BitReaderType) ReadBits64(bits int) (int64, error) {
	err := reader.checkBits(bits)
	if err != nil {
		return -1, err
	}
	rv := int64(0)
	for bits > 0 {
		value, readBits := reader.readInByte(bits)
		rv <<= int64(readBits)
		rv += value
		bits -= readBits
	}
	return rv, nil
}

func (reader *BitReaderType) ReadBits(bits int) (int, error) {
	val, err := reader.ReadBits64(bits)
	return int(val), err
}

func (reader *BitReaderType) ReadBit() (bool, error) {
	val, err := reader.ReadBits(1)
	return val != 0, err
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (reader *BitReaderType) checkBits(bits int) error {
	remain := len(reader.data) - reader.base
	remain *= BYTE
	remain -= reader.offset

	if remain < bits || bits < 0 {
		return fmt.Errorf("[Out of bound](size %d, base %d, offset %d) %d bits required, but only %d remained.", len(reader.data), reader.base, reader.offset, bits, remain)
	}
	return nil
}

func (reader *BitReaderType) readInByte(bits int) (int64, int) {
	readBits := min(bits, BYTE-reader.offset)
	unReadbits := BYTE - reader.offset - readBits
	mask := (0xff >> reader.offset) ^ (0xff >> (reader.offset + readBits))

	rv := (int64(reader.data[reader.base]) & int64(mask)) >> int64(unReadbits)

	if readBits+reader.offset == BYTE {
		reader.base++
		reader.offset = 0
	} else {
		reader.offset += readBits
	}

	return rv, readBits
}
