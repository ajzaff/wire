package wire

import (
	"encoding/binary"
)

func AppendUvarint(b []byte, x uint64) []byte { return binary.AppendUvarint(b, x) }

func AppendVarint(b []byte, x int64) []byte {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return AppendUvarint(b, ux)
}

func SizeUvarint(x uint64) int {
	n := 1
	for i := 0; x >= 0x80; i++ {
		n++
		x >>= 7
	}
	return n
}

func SizeVarint(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return SizeUvarint(ux)
}
