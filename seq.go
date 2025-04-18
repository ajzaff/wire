package wire

import (
	"iter"
)

func AppendSeq[T0 comparable](b []byte, seq iter.Seq[T0]) ([]byte, error) {
	for e := range seq {
		var err error
		if b, err = Append(b, e); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func AppendSeq2[T0, T1 comparable](b []byte, seq iter.Seq2[T0, T1]) ([]byte, error) {
	for e := range seq {
		var err error
		if b, err = Append(b, e); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func SizeSlice[E comparable, T ~[]E](x T) int {
	n := len(x)
	if n == 0 {
		return 0
	}
	if IsFixed[E](x[0]) {
		return n * Size(x[0])
	}
	var sz int
	for _, e := range x {
		sz += Size(e)
	}
	return sz
}

func AppendSlice[E comparable, T ~[]E](b []byte, x T) ([]byte, error) {
	for _, e := range x {
		var err error
		if b, err = Append(b, e); err != nil {
			return nil, err
		}
	}
	return b, nil
}
