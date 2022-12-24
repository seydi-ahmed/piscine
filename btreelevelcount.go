package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := BTreeLevelCount(root.Left)
	b := BTreeLevelCount(root.Right)
	if a > b {
		return a + 1
	} else {
		return b + 1
	}
}
