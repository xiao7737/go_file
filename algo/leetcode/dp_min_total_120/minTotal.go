package dp_min_total_120

//空间和时间都是O(M*N)  M和N表示层数和每层的个数
//从底向上进行动态规划查询，寻找最小的
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- { //  for j := 0; j <= i; j++  {    替换本行也可以
			//找到两个子节点中小的那一个，然后和当前节点相加，再逐级往上加
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	//直到加到第一个元素
	return triangle[0][0]
}

//一样的算法，采用状态压缩，用一位数组代替二维数组
//复用一个变量，空间O(N)
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	res := make([]int, len(triangle))
	//res := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			res[j] = triangle[i][j] + min(res[j], res[j+1])
		}
	}
	return res[0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
