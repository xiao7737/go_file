package main

//方法1：简单易懂，但是存在多次内存分配，效率低
func QuickSort1(sli []int) []int {
	sliLen := len(sli)
	if sliLen <= 1 {
		return sli
	}
	//选择第一个元素为基准元素
	baseNum := sli[0]
	//将基准元素两边的元素放入两个切片
	leftSli := []int{}  //小于基准元素
	rightSli := []int{} //大于基准元素
	for i := 1; i < sliLen; i++ {
		if sli[i] < baseNum {
			leftSli = append(leftSli, sli[i])
		} else {
			rightSli = append(rightSli, sli[i])
		}
	}
	//分别对两个切片进行递归处理
	leftSli = QuickSort1(leftSli)
	rightSli = QuickSort1(rightSli)

	//合并
	leftSli = append(leftSli, baseNum)
	return append(leftSli, rightSli...)
}

//方法2：高效版
//时间复杂度：nlogn  存在递归调用，所以空间复杂度logn
func QuickSort2(sli []int) []int {
	return partition2(sli, 0, len(sli)-1)
}
func partition2(sli []int, left, right int) []int {
	if left < right {
		//选择基准元素
		//选择基准元素，可以采用三数取中法
		baseNum := sli[(left+right)/2]
		i, j := left, right
		//i j相遇就停止
		for i < j {
			//i从前往后走  找到比baseNum大的
			for baseNum > sli[i] {
				i++
			}
			//j从后往前走 ，找到比baseNum小的
			for baseNum < sli[j] {
				j--
			}
			//找到符合条件的两个数，进行交换
			sli[i], sli[j] = sli[j], sli[i]
		}
		partition2(sli, left, i-1)
		partition2(sli, j+1, right)
	}
	return sli
}

//方法4，同方法2，最简化版本
func QuickSort4(sli []int) []int {
	left, right := 0, len(sli)-1
	if left < right {
		//选择基准元素
		//选择基准元素，可以采用三数取中法
		baseNum := sli[(left+right)/2]
		i, j := left, right
		//i j相遇就停止
		for i < j {
			//i从前往后走  找到比baseNum大的
			for baseNum > sli[i] {
				i++
			}
			//j从后往前走 ，找到比baseNum小的
			for baseNum < sli[j] {
				j--
			}
			//找到符合条件的两个数，进行交换
			sli[i], sli[j] = sli[j], sli[i]
		}
		QuickSort4(sli[left:i])
		QuickSort4(sli[j+1 : right+1])
	}
	return sli
}

//方法3   和方法2类似，但是比方法2性能差
func QuickSort3(sli []int) []int {
	return partition3(sli, 0, len(sli)-1)
}
func partition3(sli []int, left, right int) []int {
	baseNum := sli[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && sli[j] >= baseNum {
			j--
		}
		if j >= p {
			sli[p] = sli[j]
			p = j
		}
		for i <= p && sli[i] <= baseNum {
			i++
		}
		if i <= p {
			sli[p] = sli[i]
			p = i
		}
	}
	sli[p] = baseNum
	if p-left > 1 {
		partition3(sli, left, p-1)
	}
	if right-p > 1 {
		partition3(sli, p+1, right)
	}
	return sli
}
