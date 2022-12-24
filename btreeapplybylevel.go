package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		PlusGrand(root, i, f)
	}
}

func PlusGrand(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		PlusGrand(root.Left, i-1, f)
		PlusGrand(root.Right, i-1, f)
	}
}
