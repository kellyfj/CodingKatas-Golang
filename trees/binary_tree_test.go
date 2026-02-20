package trees

import "testing"

func TestBreadthFirstTraversal(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 2, 3, 4, 5}
	result := breadthFirstTraversal(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestDepthFirstTraversal(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 2, 4, 5, 3}
	result := depthFirstTraversal(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

// TestFindShortestPathBFS tests finding the shortest path to a leaf using BFS.
// Tree structure:
//
//	    1
//	   / \
//	  2   3 (leaf)
//	 / \
//	4   5
//
// Node 3 is a leaf with no children, making [1,3] the shortest path (length 2).
// BFS explores level-by-level and returns the first leaf node encountered.
func TestFindShortestPathBFS(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 3}
	result := findShortestPathBFS(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

// TestFindShortestPathRecursive tests finding the shortest path using recursion.
// Tree structure: Same as TestFindShortestPathBFS
//
//	    1
//	   / \
//	  2   3 (leaf)
//	 / \
//	4   5
//
// The recursive algorithm visits both subtrees and returns the path through the
// subtree with the shortest path. Node 3 is a leaf, so the right subtree has
// a path of length 1, making [1,3] the result.
func TestFindShortestPathRecursive(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 3}
	result := findShortestPathRecursive(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestFindShortestPathEmptyTree(t *testing.T) {
	expected := []int{}
	resultBFS := findShortestPathBFS(nil)
	resultRecursive := findShortestPathRecursive(nil)

	if len(resultBFS) != len(expected) {
		t.Errorf("Expected length %d, got %d for BFS", len(expected), len(resultBFS))
	}

	if len(resultRecursive) != len(expected) {
		t.Errorf("Expected length %d, got %d for Recursive", len(expected), len(resultRecursive))
	}
}

func TestFindShortestPathSingleNode(t *testing.T) {
	root := &Node{Value: 1}
	expected := []int{1}

	resultBFS := findShortestPathBFS(root)
	resultRecursive := findShortestPathRecursive(root)
	if len(resultBFS) != len(expected) {
		t.Errorf("Expected length %d, got %d for BFS", len(expected), len(resultBFS))
	}

	if len(resultRecursive) != len(expected) {
		t.Errorf("Expected length %d, got %d for Recursive", len(expected), len(resultRecursive))
	}

	for i := range expected {
		if resultBFS[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d for BFS", expected[i], i, resultBFS[i])
		}
		if resultRecursive[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d for Recursive", expected[i], i, resultRecursive[i])
		}
	}
}

// TestCompareShortestPath compares BFS and recursive approaches on the same tree.
// Tree structure:
//
//	     1
//	    / \
//	   2   3 (leaf)
//	  /
//	 4
//	/
//
// 5
// Node 3 is a leaf at depth 1, making it the shortest path regardless of algorithm.
func TestCompareShortestPath(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Left.Left = &Node{Value: 5}

	expected := []int{1, 3}
	resultBFS := findShortestPathBFS(root)
	resultRecursive := findShortestPathRecursive(root)
	if len(resultBFS) != len(expected) {
		t.Errorf("Expected length %d, got %d for BFS", len(expected), len(resultBFS))
	}

	if len(resultRecursive) != len(expected) {
		t.Errorf("Expected length %d, got %d for Recursive", len(expected), len(resultRecursive))
	}

	for i := range expected {
		if resultBFS[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d for BFS", expected[i], i, resultBFS[i])
		}
		if resultRecursive[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d for Recursive", expected[i], i, resultRecursive[i])
		}
	}
}

// TestFindShortestPathMultiplePaths tests behavior when multiple shortest paths exist.
// Tree structure: All leaves are at depth 3, so all shortest paths have equal length.
//
//	    1
//	   / \
//	  2   3
//	 / \ / \
//	4  5 6  7
//
// When multiple shortest paths exist:
// - BFS returns the first leaf encountered in breadth-first order: [1,2,4]
// - Recursive returns the path based on equal-length tie-breaking (picks right): [1,3,7]
// This demonstrates that the algorithms have different behaviors with ties.
func TestFindShortestPathMultiplePaths(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}

	expectedBFS := []int{1, 2, 4}
	expectedRecursive := []int{1, 3, 7}
	resultBFS := findShortestPathBFS(root)
	resultRecursive := findShortestPathRecursive(root)

	if len(resultBFS) != len(expectedBFS) {
		t.Errorf("Expected length %d, got %d for BFS", len(expectedBFS), len(resultBFS))
	}

	if len(resultRecursive) != len(expectedRecursive) {
		t.Errorf("Expected length %d, got %d for Recursive", len(expectedRecursive), len(resultRecursive))
	}

	for i := range expectedBFS {
		if resultBFS[i] != expectedBFS[i] {
			t.Errorf("Expected %d at index %d, got %d for BFS", expectedBFS[i], i, resultBFS[i])
		}
	}

	for i := range expectedRecursive {
		if resultRecursive[i] != expectedRecursive[i] {
			t.Errorf("Expected %d at index %d, got %d for Recursive", expectedRecursive[i], i, resultRecursive[i])
		}
	}
}

func TestFindLongestPath(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Left.Left = &Node{Value: 5}

	expected := []int{1, 2, 4, 5}

	result := findLongestPath(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestPreorderTraversal(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 2, 4, 5, 3}
	result := preorderTraversal(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{4, 2, 5, 1, 3}
	result := inorderTraversal(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestInOrderOrder1(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}

	expected := []int{2, 1, 3}
	result := inorderOrder1Space(root)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestPostorderTraversal(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{4, 5, 2, 3, 1}
	result := postorderTraversal(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestCreateListFromLeaves(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}

	expected := []int{4, 5, 6, 7}
	result := createListFromLeaves(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestLowestCommonAncestor(t *testing.T) {

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}

	expected := 1
	// Finding LCA of nodes 4 and 6, which should be 1 (the root)
	result := lowestCommonAncestor(root, root.Left.Left, root.Right.Left)

	if result == nil {
		t.Errorf("Expected non-nil result, got nil")
	} else if result.Value != expected {
		t.Errorf("Expected LCA value %d, got %d", expected, result.Value)
	}
}

func TestCheckIsSymmetric(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 2}
	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 4}
	root.Right.Left = &Node{Value: 4}
	root.Right.Right = &Node{Value: 3}

	expected := true
	result := checkIsSymmetric(root)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCheckIsSymmetricNonSymmetric(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 2}
	root.Left.Left = &Node{Value: 3}
	root.Left.Right = &Node{Value: 4}
	root.Right.Left = &Node{Value: 5} // Non-symmetric value
	root.Right.Right = &Node{Value: 3}

	expected := false
	result := checkIsSymmetric(root)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsHeightBalanced(t *testing.T) {
	root := &Node{Value: 1, Height: 0}
	root.Left = &Node{Value: 2, Height: 1}
	root.Right = &Node{Value: 3, Height: 1}
	root.Left.Left = &Node{Value: 4, Height: 2}
	root.Left.Right = &Node{Value: 5, Height: 2}

	expected := true
	result := isHeightBalanced(root)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsHeightBalancedUnbalanced(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Left.Left = &Node{Value: 3}
	root.Left.Left.Left = &Node{Value: 4} // Unbalanced node

	expected := false
	result := isHeightBalanced(root)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
