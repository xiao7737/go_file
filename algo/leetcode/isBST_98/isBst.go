package isBST_98

import "math"

// 题目：给定二叉树，验证是否为二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1  中序遍历，利用一个前继节点，只需要判断当前节点和前继节点的大小关系
// 时间O(N) 每个节点最多访问一次，空间是pre节点的空间
// 或者 将节点装入一个数组，判断数组是否升序，但是空间为O(N)
var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = &TreeNode{Val: math.MinInt64}
	return helper(root)
}
func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if helper(root.Left) == false {
		return false
	}
	if root.Val <= pre.Val {
		return false
	}
	pre = root
	return helper(root.Right)
}

// 方法2  递归，分别对左右子树递归，返回左子树的最大值l_max和右子树的最小值r_min
// 判断是否满足 l_max < root < r_min
// 时间O(N)，每个节点最多访问一次
// 空间O(N)，递归取决树的高度，最坏情况为一条链
func isValidBST1(root *TreeNode) bool {
	return helper1(root, math.MinInt64, math.MaxInt64) //初始化传入最大和最小值
}

func helper1(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	//左子树：此时最大值为root.val ,右子树：此时最小值为root.val
	//递归，如果左右子树都满足，返回true
	return helper1(root.Left, min, root.Val) && helper1(root.Right, root.Val, max)
}
