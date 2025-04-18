package wire

import "iter"

type Span[E comparable] struct{ E E }

func AppendSpan[T comparable](b []byte, x T) ([]byte, error) {
	sz := Size(b)
	b = AppendUvarint(b, uint64(sz))
	return Append(b, x)
}

func SpanSeqSlice[E comparable, T ~[]E](x T) iter.Seq[Span[E]] {
	return func(yield func(Span[E]) bool) {
		for _, e := range x {
			if !yield(Span[E]{e}) {
				break
			}
		}
	}
}

func SpanSeq[E comparable](x iter.Seq[E]) iter.Seq[Span[E]] {
	return func(yield func(Span[E]) bool) {
		for e := range x {
			if !yield(Span[E]{e}) {
				break
			}
		}
	}
}
