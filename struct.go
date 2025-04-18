package wire

import "reflect"

func AppendStruct[T comparable](b []byte, x T) ([]byte, error) {
	rv := reflect.ValueOf(x)
	for ; rv.Kind() == reflect.Pointer; rv = rv.Elem() {
	}
	return appendStruct(b, rv)
}

func appendStruct(b []byte, rv reflect.Value) ([]byte, error) {
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)
		wire := ft.Tag.Get("wire")
		_ = wire // TODO: Implement field number logic.
		_ = fv   // TODO: Write struct field.
	}
	panic("not implemented")
}
