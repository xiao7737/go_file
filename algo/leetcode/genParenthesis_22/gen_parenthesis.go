package genParenthesis_22

// 给定n，返回所有合法的扩号组合

//思路：利用递归，对子问题进行减枝，在每次加扩号的时候，进行判断
func generateParenthesis(n int) []string {
	res := new([]string)
	helper(n, n, "", res)
	return *res
}

// left 表示左括号数量
// 时间O(2的N次方)，空间O(N)
func helper(left, right int, sub string, res *[]string) {
	// 此处判断可以只判断 right == 0
	// 因为一旦右括号用完了，则左括号也用完了
	if left == 0 && right == 0 { //跳出回溯条件，扩号个数被用完了
		*res = append(*res, sub)
		return
	}
	if left > 0 {
		helper(left-1, right, sub+"(", res)
	}
	if right > left {
		helper(left, right-1, sub+")", res)
	}
}
