package slicetricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	x := []int{1, 2, 3}
	y := Copy(x)
	assert.Equal(t, y, x)
}

func TestCopyEmpty(t *testing.T) {
	x := []int{}
	y := Copy(x)
	assert.Empty(t, y)
}

func TestCutStart(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Cut(&x, 0, 2)
	assert.Equal(t, []int{2, 3}, x)
}

func TestCutMiddle(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Cut(&x, 1, 3)
	assert.Equal(t, []int{0, 3}, x)
}

func TestCutEnd(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Cut(&x, 2, 4)
	assert.Equal(t, []int{0, 1}, x)
}

func TestDeleteStart(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Delete(&x, 0)
	assert.Equal(t, []int{1, 2, 3}, x)
}

func TestDeleteMiddle(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Delete(&x, 2)
	assert.Equal(t, []int{0, 1, 3}, x)
}

func TestDeleteEnd(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Delete(&x, 3)
	assert.Equal(t, []int{0, 1, 2}, x)
}

func TestDeleteUnorderedStart(t *testing.T) {
	x := []int{0, 1, 2, 3}
	DeleteUnordered(&x, 0)
	assert.ElementsMatch(t, x, []int{1, 2, 3})
}

func TestDeleteUnorderedMiddle(t *testing.T) {
	x := []int{0, 1, 2, 3}
	DeleteUnordered(&x, 2)
	assert.ElementsMatch(t, x, []int{0, 1, 3})
}

func TestDeleteUnorderedEnd(t *testing.T) {
	x := []int{0, 1, 2, 3}
	DeleteUnordered(&x, 3)
	assert.ElementsMatch(t, x, []int{0, 1, 2})
}

func TestExpandStart(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Expand(&x, 0, 3)
	assert.Equal(t, []int{0, 0, 0, 0, 1, 2, 3}, x)
}

func TestExpandMiddle(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Expand(&x, 2, 3)
	assert.Equal(t, []int{0, 1, 0, 0, 0, 2, 3}, x)
}

func TestExpandEnd(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Expand(&x, 4, 3)
	assert.Equal(t, []int{0, 1, 2, 3, 0, 0, 0}, x)
}

func TestExpandZero(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Expand(&x, 4, 0)
	assert.Equal(t, []int{0, 1, 2, 3}, x)
}

func TestExtend(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Extend(&x, 3)
	assert.Equal(t, []int{0, 1, 2, 3, 0, 0, 0}, x)
}

func TestExtendZero(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Extend(&x, 0)
	assert.Equal(t, []int{0, 1, 2, 3}, x)
}

func TestFilter(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Filter(&x, func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, []int{0, 2}, x)
}

func TestFilterZeroAlloc(t *testing.T) {
	x := []int{0, 1, 2, 3}
	FilterZeroAlloc(&x, func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, []int{0, 2}, x)
}

func TestFilterZeroAllocNoGC(t *testing.T) {
	x := []int{0, 1, 2, 3}
	FilterZeroAllocNoGC(&x, func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, []int{0, 2}, x)
}

func TestInsertStart(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Insert(&x, 0, 9)
	assert.Equal(t, []int{9, 0, 1, 2, 3}, x)
}

func TestInsertMiddle(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Insert(&x, 2, 9)
	assert.Equal(t, []int{0, 1, 9, 2, 3}, x)
}

func TestInsertEnd(t *testing.T) {
	x := []int{0, 1, 2, 3}
	Insert(&x, 4, 9)
	assert.Equal(t, []int{0, 1, 2, 3, 9}, x)
}

func TestInsertManyStartWithoutCapacity(t *testing.T) {
	var x = make([]int, 4, 4)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 0, 7, 8, 9)
	assert.Equal(t, []int{7, 8, 9, 0, 1, 2, 3}, x)
}

func TestInsertManyMiddleWithoutCapacity(t *testing.T) {
	var x = make([]int, 4, 4)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 2, 7, 8, 9)
	assert.Equal(t, []int{0, 1, 7, 8, 9, 2, 3}, x)
}

func TestInsertManyEndWithoutCapacity(t *testing.T) {
	var x = make([]int, 4, 4)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 4, 7, 8, 9)
	assert.Equal(t, []int{0, 1, 2, 3, 7, 8, 9}, x)
}

func TestInsertManyStartWithCapacity(t *testing.T) {
	var x = make([]int, 4, 1000)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 0, 7, 8, 9)
	assert.Equal(t, []int{7, 8, 9, 0, 1, 2, 3}, x)
}

func TestInsertManyMiddleWithCapacity(t *testing.T) {
	var x = make([]int, 4, 1000)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 2, 7, 8, 9)
	assert.Equal(t, []int{0, 1, 7, 8, 9, 2, 3}, x)
}

func TestInsertManyEndWithCapacity(t *testing.T) {
	var x = make([]int, 4, 1000)
	for i := range x {
		x[i] = i
	}
	InsertMany(&x, 4, 7, 8, 9)
	assert.Equal(t, []int{0, 1, 2, 3, 7, 8, 9}, x)
}

func TestPush(t *testing.T) {
	x := []int{0, 1, 2}
	Push(&x, 3)
	assert.Equal(t, []int{0, 1, 2, 3}, x)
}

func TestPushEmpty(t *testing.T) {
	x := []int{}
	Push(&x, 0)
	assert.Equal(t, []int{0}, x)
}

func TestPushFront(t *testing.T) {
	x := []int{0, 1, 2}
	PushFront(&x, 3)
	assert.Equal(t, []int{3, 0, 1, 2}, x)
}

func TestPushFrontEmpty(t *testing.T) {
	x := []int{}
	PushFront(&x, 0)
	assert.Equal(t, []int{0}, x)
}

func TestPop(t *testing.T) {
	x := []int{0, 1, 2}
	y := Pop(&x)
	assert.Equal(t, []int{0, 1}, x)
	assert.Equal(t, 2, y)
}

func TestPopFront(t *testing.T) {
	x := []int{0, 1, 2}
	y := PopFront(&x)
	assert.Equal(t, []int{1, 2}, x)
	assert.Equal(t, 0, y)
}

func TestBatchesEven(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5}
	batches := Batches(x, 3)
	assert.Equal(t, [][]int{{0, 1, 2}, {3, 4, 5}}, batches)
}

func TestBatchesUneven(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7}
	batches := Batches(x, 3)
	assert.Equal(t, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}}, batches)
}

func TestBatchesEmpty(t *testing.T) {
	x := []int{}
	batches := Batches(x, 3)
	assert.Equal(t, [][]int{}, batches)
}

func TestReverse(t *testing.T) {
	x := []int{0, 1, 2, 3, 4}
	Reverse(x)
	assert.Equal(t, []int{4, 3, 2, 1, 0}, x)
}

func TestReverseEmpty(t *testing.T) {
	x := []int{}
	Reverse(x)
	assert.Equal(t, []int{}, x)
}

func TestSlidingWindow(t *testing.T) {
	x := []int{0, 1, 2, 3, 4}
	windows := SlidingWindow(x, 3)
	assert.Equal(t, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}}, windows)
}

func TestSlidingWindowEmpty(t *testing.T) {
	x := []int{}
	windows := SlidingWindow(x, 3)
	assert.Equal(t, [][]int{}, windows)
}

func TestSlidingWindowBiggerThanSlice(t *testing.T) {
	x := []int{0, 1}
	windows := SlidingWindow(x, 5)
	assert.Equal(t, [][]int{{0, 1}}, windows)
}

func TestSortAndDeduplicate(t *testing.T) {
	x := []int{9, 3, 3, 4, 6, 3, 6, 9, 3, 5}
	SortAndDeduplicate(&x, func(i, j int) bool {
		return x[i] < x[j]
	})
	assert.Equal(t, []int{3, 4, 5, 6, 9}, x)
}

func TestAny(t *testing.T) {
	x := []int{2, 3, 4, 5}
	assert.Equal(t, true, Any(x, func(elem int) bool {
		return elem%2 == 0
	}))

	assert.Equal(t, false, Any(x, func(elem int) bool {
		return elem%7 == 0
	}))
}

func TestAll(t *testing.T) {
	x := []int{2, 3, 4, 5}
	assert.Equal(t, true, All(x, func(elem int) bool {
		return elem > 0
	}))

	assert.Equal(t, false, All(x, func(elem int) bool {
		return elem > 2
	}))
}

func TestNone(t *testing.T) {
	x := []int{2, 3, 4, 5}
	assert.Equal(t, true, None(x, func(elem int) bool {
		return elem > 10
	}))

	assert.Equal(t, false, None(x, func(elem int) bool {
		return elem > 2
	}))
}

func TestContainsComparable(t *testing.T) {
	x := []int{2, 3, 4, 5}
	assert.Equal(t, true, ContainsComparable(x, 2))
	assert.Equal(t, true, ContainsComparable(x, 5))
	assert.Equal(t, false, ContainsComparable(x, 7))
}
