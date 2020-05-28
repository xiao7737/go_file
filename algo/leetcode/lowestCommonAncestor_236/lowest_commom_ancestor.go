package lowestCommonAncestor_235

//题目： 给定任意二叉树，返回指定两个节点p,q的最近共同祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 时间O(N)	所有节点最多访问一次
// 空间O(N)	递归取决树的高度，最坏情况为一条链
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == root || q == root || root == nil {
		return root
	}
	//对左右子树分别查找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { // p q 分别在左右子树
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
