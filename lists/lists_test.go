package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 3}}
	l2 := &ListElement{Value: 2}
	result := mergeTwoSortedLists(l1, l2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	assert.Equal(t, expected, result)
}

func TestMergeTwoSortedListsOneEmpty(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 3}}
	var l2 *ListElement
	result := mergeTwoSortedLists(l1, l2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 3}}
	assert.Equal(t, expected, result)
}

func TestMergeTwoSortedListsBothEmpty(t *testing.T) {
	var l1, l2 *ListElement
	result := mergeTwoSortedLists(l1, l2)
	assert.Nil(t, result)
}

func TestMergeTwoSortedListsDifferentLengths(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 3, Next: &ListElement{Value: 5}}}
	l2 := &ListElement{Value: 2}
	result := mergeTwoSortedLists(l1, l2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 5}}}}
	assert.Equal(t, expected, result)
}

func TestMergeTwoSortedListsWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 3}}
	l2 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := mergeTwoSortedLists(l1, l2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}}
	assert.Equal(t, expected, result)
}

func TestReverseListConstantStorage(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	result := reverseListConstantStorage(l1)
	expected := &ListElement{Value: 3, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseListConstantStorageSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := reverseListConstantStorage(l1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestReverseListConstantStorageEmpty(t *testing.T) {
	var l1 *ListElement
	result := reverseListConstantStorage(l1)
	assert.Nil(t, result)
}

func TestReverseListConstantStorageTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := reverseListConstantStorage(l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 1}}
	assert.Equal(t, expected, result)
}

func TestReverseListConstantStorageWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2}}}
	result := reverseListConstantStorage(l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseRecursively(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	result := reverseRecursively(nil, l1)
	expected := &ListElement{Value: 3, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseRecursivelySingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := reverseRecursively(nil, l1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestReverseRecursivelyEmpty(t *testing.T) {
	var l1 *ListElement
	result := reverseRecursively(nil, l1)
	assert.Nil(t, result)
}

func TestReverseRecursivelyTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := reverseRecursively(nil, l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 1}}
	assert.Equal(t, expected, result)
}

func TestReverseRecursivelyWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2}}}
	result := reverseRecursively(nil, l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseIterativelyOrderNSpace(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	result := reverseIterativelyOrderNSpace(l1)
	expected := &ListElement{Value: 3, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseIterativelyOrderNSpaceSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := reverseIterativelyOrderNSpace(l1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestReverseIterativelyOrderNSpaceEmpty(t *testing.T) {
	var l1 *ListElement
	result := reverseIterativelyOrderNSpace(l1)
	assert.Nil(t, result)
}

func TestReverseIterativelyOrderNSpaceTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := reverseIterativelyOrderNSpace(l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 1}}
	assert.Equal(t, expected, result)
}

func TestReverseIterativelyOrderNSpaceWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2}}}
	result := reverseIterativelyOrderNSpace(l1)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 1}}}
	assert.Equal(t, expected, result)
}

func TestReverseSublist(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 4, Next: &ListElement{Value: 5}}}}}
	result := reverseSublist(l1, 2, 4)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 4, Next: &ListElement{Value: 3, Next: &ListElement{Value: 2, Next: &ListElement{Value: 5}}}}}
	assert.Equal(t, expected, result)
}

func TestReverseSublistSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := reverseSublist(l1, 1, 1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestReverseSublistEmpty(t *testing.T) {
	var l1 *ListElement
	result := reverseSublist(l1, 1, 1)
	assert.Nil(t, result)
}

func TestReverseSublistTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := reverseSublist(l1, 1, 2)
	expected := &ListElement{Value: 2, Next: &ListElement{Value: 1}}
	assert.Equal(t, expected, result)
}

func TestReverseSublistWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}}
	result := reverseSublist(l1, 2, 3)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}}
	assert.Equal(t, expected, result)
}

func TestHasCycle(t *testing.T) {
	l1 := &ListElement{Value: 1}
	l2 := &ListElement{Value: 2}
	l3 := &ListElement{Value: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1 // creates a cycle

	assert.True(t, hasCycle(l1))
}

func TestHasCycleNoCycle(t *testing.T) {
	l1 := &ListElement{Value: 1}
	l2 := &ListElement{Value: 2}
	l3 := &ListElement{Value: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = nil // no cycle

	assert.False(t, hasCycle(l1))
}

func TestHasCycleSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	l1.Next = l1 // creates a cycle

	assert.True(t, hasCycle(l1))
}

func TestHasCycleEmpty(t *testing.T) {
	var l1 *ListElement
	assert.False(t, hasCycle(l1))
}

func TestHasCycleTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1}
	l2 := &ListElement{Value: 2}
	l1.Next = l2
	l2.Next = l1 // creates a cycle

	assert.True(t, hasCycle(l1))
}

func TestHasCycleWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1}
	l2 := &ListElement{Value: 2}
	l3 := &ListElement{Value: 2}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l1 // creates a cycle

	assert.True(t, hasCycle(l1))
}

func TestRemoveKthLastElement(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 4, Next: &ListElement{Value: 5}}}}}
	result := removeKthLastElement(l1, 2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 5}}}}
	assert.Equal(t, expected, result)
}

func TestRemoveKthLastElementSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := removeKthLastElement(l1, 1)
	assert.Nil(t, result)
}

func TestRemoveKthLastElementEmpty(t *testing.T) {
	var l1 *ListElement
	result := removeKthLastElement(l1, 1)
	assert.Nil(t, result)
}

func TestRemoveKthLastElementTwoElements(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	result := removeKthLastElement(l1, 1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestRemoveKthLastElementWithDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}}
	result := removeKthLastElement(l1, 2)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	assert.Equal(t, expected, result)
}

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 3}}}}}
	result := removeDuplicatesFromSortedList(l1)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	assert.Equal(t, expected, result)
}

func TestRemoveDuplicatesFromSortedListNoDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	result := removeDuplicatesFromSortedList(l1)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	assert.Equal(t, expected, result)
}

func TestRemoveDuplicatesFromSortedListSingleElement(t *testing.T) {
	l1 := &ListElement{Value: 1}
	result := removeDuplicatesFromSortedList(l1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestRemoveDuplicatesFromSortedListEmpty(t *testing.T) {
	var l1 *ListElement
	result := removeDuplicatesFromSortedList(l1)
	assert.Nil(t, result)
}

func TestRemoveDuplicatesFromSortedListAllDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 1, Next: &ListElement{Value: 1}}}
	result := removeDuplicatesFromSortedList(l1)
	expected := &ListElement{Value: 1}
	assert.Equal(t, expected, result)
}

func TestRemoveDuplicatesFromSortedListWithSomeDuplicates(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3, Next: &ListElement{Value: 3}}}}}
	result := removeDuplicatesFromSortedList(l1)
	expected := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: &ListElement{Value: 3}}}
	assert.Equal(t, expected, result)
}

func TestOverlappingLists(t *testing.T) {
	// Create two lists that overlap at a common tail
	common := &ListElement{Value: 3, Next: &ListElement{Value: 4}}
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2, Next: common}}
	l2 := &ListElement{Value: 5, Next: common}

	// Check if the two lists overlap
	assert.True(t, overlappingLists(l1, l2))
}

func TestOverlappingListsNoOverlap(t *testing.T) {
	l1 := &ListElement{Value: 1, Next: &ListElement{Value: 2}}
	l2 := &ListElement{Value: 3, Next: &ListElement{Value: 4}}

	// Check if the two lists overlap
	assert.False(t, overlappingLists(l1, l2))
}

func TestOverlappingListsOneEmpty(t *testing.T) {
	var l1 *ListElement
	l2 := &ListElement{Value: 3, Next: &ListElement{Value: 4}}

	// Check if the two lists overlap
	assert.False(t, overlappingLists(l1, l2))
}

func TestOverlappingListsBothEmpty(t *testing.T) {
	var l1, l2 *ListElement

	// Check if the two lists overlap
	assert.False(t, overlappingLists(l1, l2))
}
