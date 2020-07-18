package _5_1_max_depth

//返回树的最大深度
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先的递归
// 时间空间都是O(N)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return max(leftMax, rightMax) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
