package _4_find_num_in_2D_arr

//判断一个数是否在一个二维数组中
//该二维数组每行递增，每列递增

//暴力法，直接遍历整个二维数组
//时间O(M*N)，空间O(1)
func FindNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//减枝，利用行和列都是递增的条件
//利用最右上角的数字进行对比
//如果比target大，则该列的数都比target大，则去除这一列
//如果比target小，则该行的数都比target小，则去除这一行，target肯定不在这一行
//整体思想类似二分查找和二叉搜索树
//时间O(M+N)，空间O(1)
func FindNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
