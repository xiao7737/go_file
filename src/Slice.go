package main

import "fmt"

func main() {
	/*s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8}
	copy(s1, s2) //将切片2拷贝到切片1,切片大小以目标切片为准

	fmt.Println(s1) //783456，如将s1拷贝到s2.输出12*/

	p := []int{1, 2, 3, 4, 5}
	fmt.Println("p==", p)
	pLen := len(p)
	for i := 0; i < pLen; i++ {
		fmt.Printf("p[%d]==%d\n", i, p[i])
	}

}
