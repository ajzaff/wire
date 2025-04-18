package wire

import (
	"fmt"
	"reflect"
)

func Append[T comparable](b []byte, x T) ([]byte, error) {
	switch x := any(x).(type) {
	case struct{}:
		return b, nil
	case bool:
		if x {
			return AppendUint8(b, 1), nil
		}
		return AppendUint8(b, 0), nil
	case uint8:
		return AppendUint8(b, x), nil
	case uint16:
		return AppendUint16(b, x), nil
	case uint32:
		return AppendUint32(b, x), nil
	case uint64:
		return AppendUint64(b, x), nil
	case int8:
		return AppendUint8(b, uint8(x)), nil
	case int16:
		return AppendUint16(b, uint16(x)), nil
	case int32:
		return AppendUint32(b, uint32(x)), nil
	case int64:
		return AppendUint64(b, uint64(x)), nil
	case uint:
		return AppendUint(b, x), nil
	case int:
		return AppendUint(b, uint(x)), nil
	case float32:
		return AppendFloat32(b, x), nil
	case float64:
		return AppendFloat64(b, x), nil
	case complex64:
		return AppendComplex64(b, x), nil
	case complex128:
		return AppendComplex128(b, x), nil
	case string:
		return append(b, x...), nil
	case []byte:
		return append(b, x...), nil
	}
	rv := reflect.ValueOf(x)
	for ; rv.Kind() == reflect.Pointer; rv = rv.Elem() {
	}
	switch rt := rv.Type(); rt.Kind() {
	case reflect.Array:
		for e := range rv.Seq() {
			var err error
			if b, err = Append(b, e.Interface()); err != nil {
				return nil, err
			}
		}
		return b, nil
	case reflect.Struct:
		return appendStruct(b, rv)
	default:
		return nil, fmt.Errorf("Append: unsupported input type")
	}
}

func Size[T any](e T) int {
	panic("not implemented")
}
