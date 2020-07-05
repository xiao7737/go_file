package _5_num_of_1

//统计一个数的二进制数中的1的个数
func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num = num & (num - 1)
	}
	return sum
}
