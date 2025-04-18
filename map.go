package wire

import (
	"fmt"
)

func SizeMap[K comparable, V comparable, T ~map[K]V](m T) int {
	n := len(m)
	if n == 0 {
		return 0
	}
	var (
		k K
		v V
	)
	var (
		kFixed = IsFixed[K](k)
		vFixed = IsFixed[K](v)
	)
	var sz int
	if kFixed && vFixed {
		sz += n * (Size(k) + Size(v))
	} else if kFixed {
		sz += n * Size(k)
		for _, v := range m {
			sz += Size(v)
		}
	} else if vFixed {
		for k := range m {
			sz += Size(k)
		}
		sz += n * Size(v)
	} else {
		for k, v := range m {
			sz += Size(k) + Size(v)
		}
	}
	return sz
}

func AppendMap[K comparable, V comparable, T ~map[K]V](b []byte, m T) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func AppendSortedMap[K comparable, V comparable, T ~map[K]V](b []byte, m T) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
