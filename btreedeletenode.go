package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		var tmp *TreeNode
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			root, root.Right = root.Right, root
		} else if root.Right == nil {
			root, root.Left = root.Left, root
		} else {
			tmp = BTreeMin(root.Right)
			root.Data = tmp.Data
			root.Right = BTreeDeleteNode(root.Right, node)
		}
	}
	return root
}
