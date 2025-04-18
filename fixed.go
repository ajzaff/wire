package wire

import "math"

// copied from math pkg.
const intSize = 32 << (^uint(0) >> 63)

func SizeFixed[T comparable](x T) (int, bool) {
	switch any(x).(type) {
	case bool, uint8, int8:
		return 1, true
	case uint16, int16:
		return 2, true
	case uint32, int32, float32:
		return 4, true
	case uint64, int64, float64, complex64:
		return 8, true
	case uint, int:
		return intSize >> 3, true
	case complex128:
		return 16, true
	default:
		return 0, false
	}
}

func IsFixed[T any](x any) bool {
	switch any(x).(type) {
	case struct{},
		bool,
		uint8,
		uint16,
		uint32,
		uint64,
		int8,
		int16,
		int32,
		int64,
		uint,
		int,
		float32,
		float64,
		complex64,
		complex128:
		return true
	default:
		return false
	}
}

func AppendBool(b []byte, x bool) []byte {
	if x {
		return AppendUint8(b, 1)
	}
	return AppendUint8(b, 0)
}

func AppendUint8(b []byte, x uint8) []byte   { return append(b, x) }
func AppendUint16(b []byte, x uint16) []byte { return append(b, byte(x), byte(x>>8)) }
func AppendUint32(b []byte, x uint32) []byte {
	return append(b, byte(x), byte(x>>8), byte(x>>16), byte(x>>24))
}
func AppendUint64(b []byte, x uint64) []byte {
	return append(b, byte(x), byte(x>>8), byte(x>>16), byte(x>>24), byte(x>>32), byte(x>>40), byte(x>>48), byte(x>>56))
}
func AppendUint(b []byte, x uint) []byte {
	if intSize == 32 {
		return AppendUint32(b, uint32(x))
	}
	return AppendUint64(b, uint64(x))
}

func AppendInt8(b []byte, x int8) []byte   { return AppendUint8(b, uint8(x)) }
func AppendInt16(b []byte, x int16) []byte { return AppendUint16(b, uint16(x)) }
func AppendInt32(b []byte, x int32) []byte { return AppendUint32(b, uint32(x)) }
func AppendInt64(b []byte, x int64) []byte { return AppendUint64(b, uint64(x)) }
func AppendInt(b []byte, x int) []byte     { return AppendUint(b, uint(x)) }

func AppendFloat32(b []byte, x float32) []byte {
	ux := math.Float32bits(x)
	return AppendUint32(b, ux)
}

func AppendFloat64(b []byte, x float64) []byte {
	ux := math.Float64bits(x)
	return AppendUint64(b, ux)
}

func AppendComplex64(b []byte, x complex64) []byte {
	b = AppendFloat32(b, real(x))
	return AppendFloat32(b, imag(x))
}

func AppendComplex128(b []byte, x complex128) []byte {
	b = AppendFloat64(b, real(x))
	return AppendFloat64(b, imag(x))
}
