package merge_sort

// 归并排序  稳定排序
// 时间 O(NlogN)，空间O(N)
// 归并的数据拷贝和合并操作会严重影响排序速度
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) >> 1
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

// 合并两个子集，使其有序
func merge(a, b []int) (c []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	c = append(c, a[i:]...)
	c = append(c, b[j:]...)
	return c
}
