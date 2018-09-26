package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLeaves(t *TreeNode, leaves []int) []int {
	// Setup recursive algo to check if has left or right, if not, append to list
	// If so, check left, then right, in that order.
	if t.Left != nil {
		leaves = getLeaves(t.Left, leaves)
	}
	if t.Right != nil {
		leaves = getLeaves(t.Right, leaves)
	}
	if (t.Left == nil) && (t.Right == nil) {
		leaves = append(leaves, t.Val)
	}
	return leaves
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var root1Leaves, root2Leaves []int
	root1Leaves = getLeaves(root1, root1Leaves)
	root2Leaves = getLeaves(root2, root2Leaves)

	if len(root1Leaves) != len(root2Leaves) {
		return false
	}
	for i := range root1Leaves {
		if root1Leaves[i] != root2Leaves[i] {
			return false
		}
	}
	return true
}
