package main

import "fmt"

//func main() {
/*s1 := []int{1, 2, 3, 4, 5, 6}
s2 := []int{7, 8}
copy(s1, s2) //将切片2拷贝到切片1,切片大小以目标切片为准

fmt.Println(s1) //783456，如将s1拷贝到s2.输出12*/

/*	p := []int{1, 2, 3, 4, 5}
	fmt.Println("p==", p)
	pLen := len(p)
	for i := 0; i < pLen; i++ {
		fmt.Printf("p[%d]==%d\n", i, p[i])
	}
}*/

//左闭右开  s[a:b]表示包含从a到b-1的数，slice从0开始计数，b-a就是slice的长度
func main() {
	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	s5 := s3[:3]
	s6 := s3[3:]
	fmt.Printf("The length of s4: %d\n", len(s4))   //3
	fmt.Printf("The capacity of s4: %d\n", cap(s4)) //6 = 9-3，底层数组长度减去截取前面的长度
	fmt.Printf("The value of s4: %d\n", s4)         //3,4,5
	fmt.Printf("The value of s5: %d\n", s5)         //0,1,2
	fmt.Printf("The capacity of s5: %d\n", cap(s5)) //9
	fmt.Printf("The value of s6: %d\n", s6)         //3,4,5,6,7,8
	fmt.Printf("The capacity of s6: %d\n", cap(s6)) //6    会截掉前面的长度

	s7 := append(s5, 3)
	fmt.Println(s7) //0,1,2,3
}
