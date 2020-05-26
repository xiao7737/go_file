package maxTreeDepth_104

// 题目：求树的最大深度
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先 DFS 的递归方法
// 时间O(N)，空间O(logN)-O(N)
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftMax := maxDepth(node.Left)
	rightMax := maxDepth(node.Right)
	return max(leftMax, rightMax) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
