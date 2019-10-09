package main

import (
	"fmt"
)

func main() {
	/*a := 1                            //for 1
	  for i := 0; i < 3; i++ {
	  	a++
	  }
	  fmt.Println(a)*/

	/*a := 1                            //for 2   少申明一个变量
	  for a <= 3 {
	  	a++
	  }
	  fmt.Println(a)*/

	/*a := 1                            //switch  1
	  switch a {
	  case 1:
	  	fmt.Println("a value is 1")
	  case 0:
	  	fmt.Println("false")
	  }*/

	/*a := 1                              //switch 2
	  switch {
	  case a > 1:
	  	fmt.Println("a value is bigger than 1")
	  case a == 1:
	  	fmt.Println("a value is 1")
	  }*/

	/*switch a := 1; {                           //switch 3 初始化表达式，右侧需要分号
	  case a == 1:
	  	fmt.Println("a value is 1")
	  	fallthrough    							 //执行后面的case，否则将就在这里跳出循环，默认自带break
	  case a > 0:
	  	fmt.Println("a value is bigger than 0")
	  	fallthrough
	  default:
	  	fmt.Println("a has a default value")
	  }*/

	a := [...]int{1, 3, 4, 5, 7, 9} //从大到小冒泡排序
	fmt.Println(a)

	length := len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

/*var slice1 = []int{1, 2, 3, 4}

func main() {
	for i, v := range slice1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}*/

/*func main() {
x := []string{"a", "b", "c"}
/*for _, v := range x {
	fmt.Println(v)   //abc
}*/
/*for v := range x {
		fmt.Println(v) //012   会默认循环键，而不是值
	}
}*/

/*func main() {
	strings := []string{"one", "two", "three"}
	for _, s := range strings {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s", s)
		}()
	}
	time.Sleep(3 * time.Second)
	//输出three three three
	//go func中的s和for中的s一样，地址不变化，
	//遍历完成，最终s读到最后一个元素
	//应该改写for循环
	//for i=0;i<len(strings) ;i++  {}

}*/