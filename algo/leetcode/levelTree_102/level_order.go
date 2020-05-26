package levelTree_102

//题目： 将二叉树，按照层序遍历返回，返回每一层的元素，装在二维数组中

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先BFS，遍历每一层
// 每个节点都访问了一次，时间 O(N)，空间O(N)
// 如果不是树，而是图，要用一个visited的map去判重
func levelTree(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// 初始化第一层，只有root节点，用来存当前层的节点
	levelNodeArr := []*TreeNode{root}
	for i := 0; len(levelNodeArr) > 0; i++ {

		// 给二维数组，初始化一个一维int类型的空数组，后面填充数据
		res = append(res, []int{})

		//用来装下一层的元素
		nextLevelArr := []*TreeNode{}

		//遍历当前层的节点
		for j := 0; j < len(levelNodeArr); j++ {
			node := levelNodeArr[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				nextLevelArr = append(nextLevelArr, node.Left)
			}
			if node.Right != nil {
				nextLevelArr = append(nextLevelArr, node.Right)
			}
		}
		levelNodeArr = nextLevelArr
	}
	return res
}
