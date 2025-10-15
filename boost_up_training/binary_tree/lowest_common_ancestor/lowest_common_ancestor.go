package lowest_common_ancestor

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// if both side subtree returned nodes, then current node is LCA
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}
