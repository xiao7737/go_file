package main

import (
	"fmt"
	"reflect"
)

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
func test1() {
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
	fmt.Println(reflect.TypeOf(s7)) //[]int
	fmt.Printf("%T\n", s7)          //[]int
	fmt.Println(s7)                 //0,1,2,3
}

//切片相加
/*func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...) //语法糖，循环，效果同上面的循环

	fmt.Println(s1)

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //00000123   初始化的slice的默认值占位
	w := make([]int, 0)
	w = append(w, 1, 2, 3)
	fmt.Println(w) //123
}*/

/*
slice也为值传递，虽然是引用类型，go语言只有值传递
slice的源码结构：
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

函数传递的是这个结构体，只是结构体里面有指针，实际还是传值
是否是引用传递：形参和实参的地址是否一致
*/

//切片扩容，不一定是双倍，因为有内存对齐的过程
//按照理解，cap应该为8，但是结果为6
/*func main() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}*/

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s) //myAppend改变了自己函数内的slice，但是main函数中的s并未改变，还是111
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s) //myAppend改变了自己函数内的slice，也改变了main函数中的s
	fmt.Println(s)
}

//当直接用切片作为函数参数时，可以改变切片的元素，不能改变切片本身

//https://www.cnblogs.com/qcrao-2018/p/10631989.html
