package _7_mirror_tree

//求树的镜像树，即左右节点完全交换
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归，dfs，直接将左右节点进行交换即可
//时间，空间都是O(N)
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

//利用栈，入栈是先左后右，出栈就先填充左再后，就实现了左右交换
//时间，空间都是O(N)
