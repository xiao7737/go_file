package _7_build_tree

//提供树的前序遍历和中序遍历，输出该二叉树
//输入内容都不含重复数字
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据前序和中序可知，3为根节点，9为左子树，15，20，7为右子树，对左右子树递归
//时间空间都是O(N)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 || inorder == nil || len(inorder) == 0 {
		return nil
	}
	var left, right []int
	for k, v := range inorder {
		if v == preorder[0] { //确定根节点
			left, right = inorder[:k], inorder[k+1:] //左子树,右子树
			break
		}
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:len(left)+1], left), //注意此处从1开始，因为第一个是root，需要跳过
		Right: buildTree(preorder[len(left)+1:], right),
	}
	return root
}
