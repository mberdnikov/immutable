package immutable

import (
	"encoding/binary"
	"math"
)

func varintSize(i uint64) int {
	switch {
	case i <= 0xFC:
		return 1
	case i <= math.MaxUint16:
		return 2 + 1
	case i <= math.MaxUint32:
		return 4 + 1
	default:
		return 8 + 1
	}
}

func putVarint(b []byte, i uint64) int {
	switch {
	case i <= 0xFC:
		b[0] = byte(i)
		return 1
	case i <= math.MaxUint16:
		binary.LittleEndian.PutUint16(b[1:], uint16(i))
		b[0] = 0xFD
		return 2 + 1
	case i <= math.MaxUint32:
		binary.LittleEndian.PutUint32(b[1:], uint32(i))
		b[0] = 0xFE
		return 4 + 1
	default:
		binary.LittleEndian.PutUint64(b[1:], i)
		b[0] = 0xFF
		return 8 + 1
	}
}

func getVarint(b []byte) (uint64, int) {
	switch b[0] {
	case 0xFF:
		return binary.LittleEndian.Uint64(b[1:]), 8 + 1
	case 0xFE:
		return uint64(binary.LittleEndian.Uint32(b[1:])), 4 + 1
	case 0xFD:
		return uint64(binary.LittleEndian.Uint16(b[1:])), 2 + 1
	default:
		return uint64(b[0]), 1
	}
}
