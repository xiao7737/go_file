package _2_huisu_dfs

//题目：判断一个矩阵中是否存在一个word
//思路：DFS的回溯处理

//时间O(MN)，M，N 为rows，cols
//最差时间O(3的k次方)，每个字符都有4种选择，但是该题会舍弃一种，为3种
//空间O(K) k=len(word)
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	//已经完全匹配到该word
	if k == len(word) {
		return true
	}

	//i,j边界
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] == word[k] {
		temp := board[i][j]
		//临时占位修改board[i][j]，避免下一次递归中使用到当前格子，因为题目要求每个格子只能使用一次
		//如果board[i][j]的上下左右都找不到符合条件的值，进行还原
		board[i][j] = '*'
		//对board[i][j]上下左右的元素进行匹配
		if dfs(board, i-1, j, k+1, word) ||
			dfs(board, i+1, j, k+1, word) ||
			dfs(board, i, j-1, k+1, word) ||
			dfs(board, i, j+1, k+1, word) {
			return true
		} else {
			//board[i][j]上下左右的元素都没有匹配，还原
			board[i][j] = temp
		}
	}
	return false
}
