package minTreeDepth_111

// 题目：返回树的最小深度
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMini := minDepth(root.Left)
	rightMini := minDepth(root.Right)
	if leftMini == 0 || rightMini == 0 {
		return leftMini + rightMini + 1
	}
	return min(leftMini, rightMini) + 1
}

// == Math库里面的Min方法，是float类型的，只好自己实现==！
// 讲道理int类型的不是更常用咩
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
