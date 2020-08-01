package _4_get_Kth_num_in_tree

//返回二叉搜索树的第k大元素
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据中序遍历的原理，右中左遍历得到从大到小的倒序结果
//时间空间都是O(N)
func kthLargest(root *TreeNode, k int) int {
	//用来装倒序节点的map
	res := make([]int, 0)
	helper(root, &res)
	return res[k-1]
}

func helper(root *TreeNode, res *[]int) {
	if root.Right != nil {
		helper(root.Right, res)
	}
	if root != nil {
		*res = append(*res, root.Val)
	}
	if root.Left != nil {
		helper(root.Left, res)
	}
}
