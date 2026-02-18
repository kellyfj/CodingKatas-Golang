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

/*FIXME
func TestFindShortestPathBFS(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 2, 4}
	result := findShortestPathBFS(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}*/

/*FIXME
func TestFindShortestPathRecursive(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	expected := []int{1, 2, 4}
	result := findShortestPathRecursive(root)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}*/

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

/*FIXME
func TestCompareShortestPath(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Left.Left = &Node{Value: 5}

	expected := []int{1, 2, 4, 5}
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

func TestFindShortestPathMultiplePaths(t *testing.T) {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}

	expected := []int{1, 2, 4}
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
}*/

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
