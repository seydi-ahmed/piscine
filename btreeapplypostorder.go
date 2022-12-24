package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	} else {
		BTreeApplyPostorder(root.Left, f)
		BTreeApplyPostorder(root.Right, f)
		f(root.Data)
	}
}
