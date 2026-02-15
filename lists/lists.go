package lists

type ListElement struct {
	Value int
	Next  *ListElement
}

/**
 * EPIJ 7.1 Merge two sorted lists
 *
 * Time Complexity: O(n+m)
 * Space Complexity: O(1) reusing nodes
 */
func mergeTwoSortedLists(l1, l2 *ListElement) *ListElement {
	dummyHead := &ListElement{}
	returnVal := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			// move
			returnVal.Next = l1
			l1 = l1.Next
		} else {
			returnVal.Next = l2
			l2 = l2.Next
		}
		returnVal = returnVal.Next
	}

	if l1 != nil {
		returnVal.Next = l1
	} else {
		returnVal.Next = l2
	}

	return dummyHead.Next
}

func reverseListConstantStorage(head *ListElement) *ListElement {
	var previous *ListElement
	current := head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

func reverseRecursively(previous, current *ListElement) *ListElement {
	if current == nil {
		return previous
	}
	next := current.Next
	current.Next = previous
	return reverseRecursively(current, next)
}

func reverseIterativelyOrderNSpace(head *ListElement) *ListElement {
	var nodes []*ListElement
	current := head

	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	var previous *ListElement
	for i := 0; i < len(nodes); i++ {
		nodes[i].Next = previous
		previous = nodes[i]
	}
	return previous
}

/**
 * EPIJ 7.2: Reverse a single sublist within a list
 * Key - perform the the update in a single pass combining the identification with the reversal
 *
 * Time Complexity: Dominated by search for fth node i.e. O(f)
 * Space Complexity: O(1)
 */
func reverseSublist(head *ListElement, start, finish int) *ListElement {
	dummyHead := &ListElement{Next: head}
	sublistHead := dummyHead

	for i := 1; i < start; i++ {
		sublistHead = sublistHead.Next
	}

	sublistIter := sublistHead.Next

	// A --> B --> C --> D --> E start = 2, finish = 4
	// First iteration: A --> C --> B --> D --> E
	// Second iteration: A --> D --> C --> B --> E
	// Third iteration: A --> D --> C --> B	--> E
	for i := start; i < finish; i++ {
		temp := sublistIter.Next
		sublistIter.Next = temp.Next
		temp.Next = sublistHead.Next
		sublistHead.Next = temp
	}
	return dummyHead.Next
}

/**
 * EPIJ 7.3 Test for Cyclicity
 * 	Brute force solution uses a HashMap/HashSet and O(n) additional space.
 *  Brute force solution without additional storage is to use two for loops but is O(n^2)
 *  Can make it work with O(n) time and O(1) space by using two pointers - slow and fast.
 *  The list has a cycle if the two iterators meet.
 *
 *  Time Complexity: ???
 *  Space Complexity: O(1)
 */
// public static ListElement hasCycle(ListElement head) {
func hasCycle(head *ListElement) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/**
 * EPIJ 7.4: Test for overlapping lists - lists are cycle free
 *
 * Given two singly linked lists there may be list nodes that are common to both
 * Write a program that takes two such lists and determines if there is a node
 * that is common to both. Brute force - use a HashSet - takes O(n) time and
 * O(n) space.
 *
 * Lists overlap if and only if both have the same tail node: once the lists
 * converge at a node they cannot diverge at a later node. So checking for
 * overlap amounts to finding the tail nodes for each list.
 *
 * Time Complexity: O(n) Space Complexity: O(1)
 */
func overlappingLists(l1, l2 *ListElement) bool {
	if l1 == nil || l2 == nil {
		return false
	}

	iter1, iter2 := l1, l2

	for iter1.Next != nil {
		iter1 = iter1.Next
	}

	for iter2.Next != nil {
		iter2 = iter2.Next
	}

	return iter1 == iter2
}

/**
 * EPIJ 7.7: Remove k'th last elememnt from a list
 *
 * TIME COMPLEXITY: O(n) length of the list
 * SPACE COMPLEXITY: O(1)
 */
func removeKthLastElement(head *ListElement, k int) *ListElement {
	dummyHead := &ListElement{Next: head}
	fast, slow := dummyHead, dummyHead

	// Algorithm: Move fast k steps ahead, then move both iterators at the same pace
	// until fast reaches the end of the list. At this point, slow will be at the (k+1)th last element.
	// Update slow's next pointer to skip the kth last element.
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			// k is greater than the length of the list, return the original list
			return head
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

/**
 * EPIJ 7.8: Remove duplicates from a sorted list
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
func removeDuplicatesFromSortedList(head *ListElement) *ListElement {
	current := head

	for current != nil && current.Next != nil {
		if current.Value == current.Next.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
