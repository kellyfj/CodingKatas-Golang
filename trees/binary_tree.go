package trees

type Node struct {
	Value  int
	Height int
	Left   *Node
	Right  *Node
}

//TOODO: Pass a function to apply to each node during traversal, instead of just visiting the node and appending its value to a result list. This would allow for more flexible operations during traversal, such as summing values, finding specific nodes, etc.
/**
 * Traversal algorithms for binary trees.
 * Time complexity: O(n) where n is the number of nodes in the tree.
 * Space complexity: O(n) in the worst case (for a completely unbalanced tree),
 *    O(log n) for a balanced tree.
 * @author kellyfj
 */
func breadthFirstTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}

func depthFirstTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Value)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

// findShortestPath finds the shortest path from the root to any leaf node in a binary tree.
// Algorithm: Breadth-First Search (BFS) is used to explore the tree level by level, ensuring that
// the first leaf node encountered is the closest one.
func findShortestPathBFS(root *Node) []int {
	if root == nil {
		return []int{}
	}

	queue := [][]*Node{{root}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1]

		if current.Left == nil && current.Right == nil {
			result := []int{}
			for _, node := range path {
				result = append(result, node.Value)
			}
			return result
		}

		if current.Left != nil {
			newPath := append([]*Node{}, path...)
			newPath = append(newPath, current.Left)
			queue = append(queue, newPath)
		}
		if current.Right != nil {
			newPath := append([]*Node{}, path...)
			newPath = append(newPath, current.Right)
			queue = append(queue, newPath)
		}
	}

	return []int{}
}

func findShortestPathRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Value}
	}

	leftPath := findShortestPathRecursive(root.Left)
	rightPath := findShortestPathRecursive(root.Right)

	if len(leftPath) == 0 {
		return append([]int{root.Value}, rightPath...)
	}
	if len(rightPath) == 0 {
		return append([]int{root.Value}, leftPath...)
	}

	if len(leftPath) < len(rightPath) {
		return append([]int{root.Value}, leftPath...)
	} else {
		return append([]int{root.Value}, rightPath...)
	}
}

func findLongestPath(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftPath := findLongestPath(root.Left)
	rightPath := findLongestPath(root.Right)

	if len(leftPath) > len(rightPath) {
		return append([]int{root.Value}, leftPath...)
	} else {
		return append([]int{root.Value}, rightPath...)
	}
}

func preorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Value}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

func inorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Value)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

// TBD
func inorderOrder1Space(root *Node) []int {
	return nil
}

func postorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Value)
	return result
}

func createListFromLeaves(root *Node) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Value}
	}

	result := []int{}
	result = append(result, createListFromLeaves(root.Left)...)
	result = append(result, createListFromLeaves(root.Right)...)
	return result
}

func lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func checkIsSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isHeightBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil {
		return false
	}
	if root.Right == nil {
		return false
	}
	leftHeight := root.Left.Height
	rightHeight := root.Right.Height

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return isHeightBalanced(root.Left) && isHeightBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
