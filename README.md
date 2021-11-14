# bitreader
A bit reader tool written in golang

Usage:

```
import "github.com/potterxu/bitreader"

// data is a []byte you want to read in bitmode
reader := bitreader.Bitreader(data) // call the constructor with raw data

// skip data
reader.SkipBits(bitsCnt)
reader.SkipBytes(bytesCnt)

// read flag
flag, err := reader.ReadBit()

// read int
intVal, err := reader.ReadBits(bitsCnt)

// read int64
int64Val, err := reader.ReadBits(bitsCnt)

```
