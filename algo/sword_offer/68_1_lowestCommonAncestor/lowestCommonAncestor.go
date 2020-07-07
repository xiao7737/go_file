package _8_1_lowestCommonAncestor

//二叉搜索树的最近公共祖先

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

	if p.Val < root.Val && q.Val < root.Val { // p,q在root的左子树
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val { // p,q在root的右子树
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root //p,q一个在左子树，一个在右子树
	}
}

// 迭代法
// 时间O(N)	所有节点最多访问一次
// 空间O(N)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
