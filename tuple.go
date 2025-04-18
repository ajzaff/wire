package wire

import (
	"cmp"
	"fmt"
	"iter"
)

type PartialOrderedPair[T0 cmp.Ordered, T1 comparable] struct {
	E0 T0
	E1 T1
}

type OrderedPair[T0, T1 cmp.Ordered] struct {
	E0 T0
	E1 T1
}

type Pair[T0, T1 comparable] struct {
	E0 T0
	E1 T1
}

func CollectPairs[T0, T1 comparable](seq iter.Seq2[T0, T1]) []Pair[T0, T1] {
	return AppendPairs(nil, seq)
}

func AppendPairs[T0, T1 comparable](pairs []Pair[T0, T1], seq iter.Seq2[T0, T1]) []Pair[T0, T1] {
	for e0, e1 := range seq {
		pairs = append(pairs, Pair[T0, T1]{e0, e1})
	}
	return pairs
}

func AppendPair[T0, T1 comparable](b []byte, x Pair[T0, T1]) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

type Triple[T0, T1, T2 comparable] struct {
	E0 T0
	E1 T1
	E2 T2
}

func AppendTriple[T0, T1, T2 comparable](b []byte, x Triple[T0, T1, T2]) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
