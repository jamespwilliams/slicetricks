// Package slicetricks provides generic functions for performing (most of) the operations described in
// https://github.com/golang/go/wiki/SliceTricks.
//
// Ideas for future enhancement:
// * add safe versions of methods that can fail (like Cut, Delete, Pop, etc), which return errors in failure cases.
// * add a version of SortAndDeduplicate for non-comparable types.
package slicetricks

import (
	"sort"
)

func Copy[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

// Cut removes elements starting at start and ending at end (non-inclusive) from a.
func Cut[T any](a *[]T, start, end int) {
	copy((*a)[start:], (*a)[end:])
	for k, n := len(*a)-end+start, len(*a); k < n; k++ {
		var zero T
		(*a)[k] = zero
	}
	*a = (*a)[:len(*a)-end+start]
}

// Delete removes the i'th element from a.
func Delete[T any](a *[]T, i int) {
	copy((*a)[i:], (*a)[i+1:])
	var t T
	(*a)[len(*a)-1] = t
	*a = (*a)[:len(*a)-1]
}

// DeleteUnordered is a faster alternative to Delete if you don't care about changing the order
// of items in the slice.
func DeleteUnordered[T any](a *[]T, i int) {
	(*a)[i] = (*a)[len(*a)-1]
	var zero T
	(*a)[len(*a)-1] = zero
	*a = (*a)[:len(*a)-1]
}

// Expand inserts n elements of the zero value of T after the i'th element of a.
func Expand[T any](a *[]T, i, n int) {
	*a = append((*a)[:i], append(make([]T, n), (*a)[i:]...)...)
}

// Extend adds n elements of the zero value of T at the end of a.
func Extend[T any](a *[]T, n int) {
	*a = append(*a, make([]T, n)...)
}

// Insert inserts elem into a at index i.
func Insert[T any](a *[]T, i int, elem T) {
	var zero T
	*a = append(*a, zero /* use the zero value of the element type */)
	copy((*a)[i+1:], (*a)[i:])
	(*a)[i] = elem
}

// InsertMany inserts elems into a at index i.
//
// NOTE: the implementation of this method is different from that in SliceTricks itself. This implementation
// is optimised for doing the insertion in-place.
func InsertMany[T any](a *[]T, i int, elems ...T) {
	if n := len(*a) + len(elems); n <= cap(*a) {
		*a = (*a)[:n]
		copy((*a)[i+len(elems):], (*a)[i:i+len(elems)+1])
		copy((*a)[i:], elems)
		return
	}

	s2 := make([]T, len(*a)+len(elems))
	copy(s2, (*a)[:i])
	copy(s2[i:], elems)
	copy(s2[i+len(elems):], (*a)[i:])
	*a = s2
}

// Push adds elem to the end of a.
func Push[T any](a *[]T, elem T) {
	*a = append(*a, elem)
}

// PushFront inserts elem at the start of a.
func PushFront[T any](a *[]T, elem T) {
	*a = append([]T{elem}, *a...)
}

// Pop removes and returns the element at the end of a.
func Pop[T any](a *[]T) T {
	var x T
	x, *a = (*a)[len(*a)-1], (*a)[:len(*a)-1]
	return x
}

// PopFront removes and returns the element at the start of a.
func PopFront[T any](a *[]T) T {
	var x T
	x, *a = (*a)[0], (*a)[1:]
	return x
}

/* "Additional Tricks" */

// Batches returns batches of a with maximum size batchSize while performing minimal allocations. All elements in a will
// be returned in a batch - the last batch may be smaller than batchSize.
func Batches[T any](a []T, batchSize int) [][]T {
	if len(a) == 0 {
		return [][]T{}
	}

	batches := make([][]T, 0, (len(a)+batchSize-1)/batchSize)

	for batchSize < len(a) {
		a, batches = a[batchSize:], append(batches, a[0:batchSize:batchSize])
	}
	batches = append(batches, a)

	return batches
}

// Filter removes any elements from a for which keep returns false.
func Filter[T any](a *[]T, keep func(t T) bool) {
	n := 0
	for _, x := range *a {
		if keep(x) {
			(*a)[n] = x
			n++
		}
	}

	for i := n; i < len(*a); i++ {
		var zero T
		(*a)[i] = zero
	}

	*a = (*a)[:n]
}

// FilterNoGC removes any elements from a for which keep returns false.  It may not be possible to garbage collect
// elements after removal, so calling this method can lead to memory leaks.
func FilterNoGC[T any](a *[]T, keep func(t T) bool) {
	n := 0
	for _, x := range *a {
		if keep(x) {
			(*a)[n] = x
			n++
		}
	}

	*a = (*a)[:n]
}

func Reverse[T any](a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// SlidingWindow returns subarrays of a of size size, starting at increasing indices of a. For example,
// SlidingWindow([0 1 2 3 4 5], 3) = [[0 1 2] [1 2 3] [2 3 4] [3 4 5]].
func SlidingWindow[T any](a []T, size int) [][]T {
	if len(a) == 0 {
		return [][]T{}
	}

	if len(a) <= size {
		return [][]T{a}
	}

	// allocate slice at the precise size we need
	r := make([][]T, 0, len(a)-size+1)

	for i, j := 0, size; j <= len(a); i, j = i+1, j+1 {
		r = append(r, a[i:j])
	}

	return r
}

// SortAndDeduplicate sorts the given slice and removes duplicate elements.
func SortAndDeduplicate[T comparable](a *[]T, less func(i, j int) bool) {
	// TODO: maybe another verson of this function for non-comparable types (e.g: passing an equals() function or using
	// an interface) would be useful.
	sort.SliceStable(*a, less)

	j := 0
	for i := 1; i < len(*a); i++ {
		if (*a)[j] == (*a)[i] {
			continue
		}
		j++
		(*a)[j] = (*a)[i]
	}
	*a = (*a)[:j+1]
}

/* A couple more methods that aren't in SliceTricks but I couldn't help adding */

// All returns true iff all elements in elem evaluate to true when passed to f.
func All[T any](elems []T, f func(a T) bool) bool {
	for _, elem := range elems {
		if !f(elem) {
			return false
		}
	}
	return true
}

// Any returns true iff any element in elems evaluates to true when passed to f.
func Any[T any](elems []T, f func(a T) bool) bool {
	for _, elem := range elems {
		if f(elem) {
			return true
		}
	}
	return false
}

// Any returns true iff no element in elems evaluates to true when passed to f.
func None[T any](elems []T, f func(a T) bool) bool {
	return !Any(elems, f)
}

// ContainsComparable returns true iff the given haystack contains the given needle.
func ContainsComparable[T comparable](haystack []T, needle T) bool {
	return Any(haystack, func(elem T) bool {
		return elem == needle
	})
}
