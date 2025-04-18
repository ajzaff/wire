package wire

import (
	"cmp"
	"slices"
)

// SortTable sorts the table based on the given cmp function which may implement multimap sorting.
func SortTable[K cmp.Ordered, V comparable](table []PartialOrderedPair[K, V]) {
	slices.SortFunc(table, func(a, b PartialOrderedPair[K, V]) int {
		if a.E0 < b.E0 {
			return -1
		}
		if b.E0 < a.E0 {
			return +1
		}
		return 0
	})
}

// SortTableFunc sorts the table based on the given cmp function which may implement multimap sorting.
func SortTableFunc[K cmp.Ordered, V comparable](table []PartialOrderedPair[K, V], cmp func(a, b K) int) {
	slices.SortFunc(table, func(a, b PartialOrderedPair[K, V]) int { return cmp(a.E0, b.E0) })
}

// Sort the multi mapped table based on the given cmp function which may implement multimap sorting.
func SortMultiTableFunc[K, V cmp.Ordered](table []OrderedPair[K, V], cmp func(a, b OrderedPair[K, V]) int) {
	slices.SortFunc(table, cmp)
}

// Sort the multi mapped table based on the given cmp function which may implement multimap sorting.
func SortMultiTable[K, V cmp.Ordered](table []OrderedPair[K, V]) {
	slices.SortFunc(table, func(a, b OrderedPair[K, V]) int {
		if a.E0 < b.E0 {
			return -1
		}
		if b.E0 < a.E0 {
			return +1
		}
		if a.E1 < b.E1 {
			return -1
		}
		if b.E1 < a.E1 {
			return +1
		}
		return 0
	})
}

// CompactTable sorts the table entries and compacts equal keys.
//
// CompactTable removes multiple values for the same key, only keeping the first.
func CompactTable[K cmp.Ordered, V comparable](table []PartialOrderedPair[K, V]) []PartialOrderedPair[K, V] {
	SortTable(table)
	return slices.CompactFunc(table, func(a, b PartialOrderedPair[K, V]) bool { return a.E0 == b.E0 })
}

// CompactMultiTable sorts the table entries and compacts equal keys and values.
//
// CompactMultiTable removes duplicate (K, V) entries.
func CompactMultiTable[K, V cmp.Ordered](table []OrderedPair[K, V]) []OrderedPair[K, V] {
	SortMultiTable(table)
	return slices.CompactFunc(table, func(a, b OrderedPair[K, V]) bool { return a == b })
}

// CompactTableFunc sorts the table entries and compacts equal keys.
//
// CompactTableFunc removes multiple values for the same key, only keeping the first.
func CompactTableFunc[K cmp.Ordered, V comparable](table []PartialOrderedPair[K, V], cmp func(a, b K) int) []PartialOrderedPair[K, V] {
	SortTableFunc(table, cmp)
	return slices.CompactFunc(table, func(a, b PartialOrderedPair[K, V]) bool { return cmp(a.E0, b.E0) == 0 })
}

// CompactMultiTableFunc sorts the table entries and compacts equal keys and values.
//
// CompactMultiTableFunc removes duplicate (K, V) entries.
func CompactMultiTableFunc[K, V cmp.Ordered](t []OrderedPair[K, V], cmp func(a, b OrderedPair[K, V]) int) []OrderedPair[K, V] {
	SortMultiTableFunc(t, cmp)
	return slices.CompactFunc(t, func(a, b OrderedPair[K, V]) bool { return cmp(a, b) == 0 })
}
