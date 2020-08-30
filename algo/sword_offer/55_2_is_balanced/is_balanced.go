package _5_2_is_balanced

import "math"

//判断树是否为平衡二叉树，即任意节点的左右子树的深度相差不超过1
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//采用dfs求出左右子树的最大深度，再递归判断
//从顶至底
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}
	max_left := dfs(root.Left)
	max_right := dfs(root.Right)
	if math.Abs(float64(max_left-max_right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	} else {
		return false
	}
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
