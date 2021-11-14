package bitreader

import (
	"fmt"
	"log"
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

func (reader *BitReaderType) SkipBytes(bytes int) {
	reader.SkipBits(bytes * BYTE)
}

func (reader *BitReaderType) SkipBits(bits int) {
	reader.checkBits(bits)
	bits += reader.offset
	reader.base += bits / BYTE
	reader.offset += bits % BYTE
}

func (reader *BitReaderType) ReadBits64(bits int) int64 {
	reader.checkBits(bits)
	rv := int64(0)
	for bits > 0 {
		value, readBits := reader.readInByte(bits)
		rv <<= int64(readBits)
		rv += value
		bits -= readBits
	}
	return rv
}

func (reader *BitReaderType) ReadBits(bits int) int {
	return int(reader.ReadBits64(bits))
}

func (reader *BitReaderType) ReadBit() bool {
	return reader.ReadBits(1) != 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (reader *BitReaderType) checkBits(bits int) bool {
	remain := len(reader.data) - reader.base
	remain *= BYTE
	remain -= reader.offset

	if remain < bits || bits < 0 {
		log.Println(fmt.Sprintf("[Out of bound](size %d, base %d, offset %d) %d bits required, but only %d remained.", len(reader.data), reader.base, reader.offset, cnt, remain))
	}
	return true
}

func (reader *BitReaderType) readInByte(bits int) (rv int64, readBits int) {
	readBits = min(bits, BYTE-reader.offset)
	unReadbits := BYTE - reader.offset - readBits
	mask := (0xff >> reader.offset) ^ (0xff >> (reader.offset + readBits))

	rv = (int64(reader.data[reader.base]) & int64(mask)) >> int64(unReadbits)

	if readBits+reader.offset == BYTE {
		reader.base++
		reader.offset = 0
	} else {
		reader.offset += readBits
	}

	return
}
