package _8_is_symmetric

//判断树是否对称
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
//时间O(N)，最多调用N/2次helper()递归
//时间O(N)，树最差退化为链表，递归使用N的栈空间
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	//两边都是nil
	if left == nil && right == nil {
		return true
	}
	//如果任意一边不为nil，或者两边不等
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	//对左右子树递归
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
