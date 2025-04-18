package pack

import (
	"iter"

	"github.com/ajzaff/huffman"
)

func copySeq[E any](dst []E, pullFn func() (x E, ok bool)) (n int) {
	for i := range len(dst) {
		e, ok := pullFn()
		if !ok {
			break
		}
		dst[i] = e
		n++
	}
	return
}

func AppendBoolSeq(b []byte, seq iter.Seq[bool]) []byte {
	buf := make([]bool, 8) // Max pack.
	pullFn, stop := iter.Pull(seq)
	defer stop()
	for {
		n := copySeq(buf[:], pullFn)
		if n == 0 {
			break
		}
		b = appendBoolSlice(b, buf[:n])
	}
	return b
}

func appendBoolField(x []bool) byte {
	var b byte
	for i, e := range x {
		if e {
			b |= 1 << i
		}
	}
	return b
}

func AppendBoolSlice(b []byte, x []bool) []byte {
	n := len(x)
	for i := 0; i < n; {
		i1 := i + 56
		if n < i1 {
			i1 = n
		}
		b = appendBoolSlice(b, x[i:i1])
		i = i1
	}
	return b
}

func appendBoolSlice(b []byte, x []bool) []byte {
	n := len(x)
	uv := huffman.Elem(56 - n)
	b = huffman.Append(b, uv, 56)
	for i := 0; i < n; {
		i1 := i + 8
		if n < i1 {
			i1 = n
		}
		b = append(b, appendBoolField(x[i:i1]))
		i = i1
	}
	return b
}
