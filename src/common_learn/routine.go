package main

import (
	"fmt"
	"runtime"
)

//不要对go函数的执行时间做任何假设

/*①
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond) //毫秒
		fmt.Println(s)
	}
}
func main() {
	go say("world") //将会异步打印这两个字符串，顺序不定
	say("hello")
}*/

/*②
func main() {
	name := "tom1"
	go func() {
		fmt.Printf("hello,%s\n", name)
	}()
	name = "tom2"
	time.Sleep(time.Millisecond)
	//会输出hello，tom2，如果交换最后两行顺序，将会输出hello，tom1
	//主函数也是一个goroutine，主函数只对name重复赋值了两次，如果交换顺序，则赋值一次
}*/

/*③
func main() {
	names := []string{"tom1", "tom2", "tom3"}
	for _, name := range names {
		go func() {
			fmt.Printf("hello:%s\n", name)
		}()
	}
	time.Sleep(time.Millisecond)
}*/

//结果为三个hello:tom3,因为for循环完毕后才执行的打印函数
//第一种方法：将睡眠函数放在for循环里面
//第二种方法：将循环的值当作go函数的参数传入，如下所示
/*func main() {
	names := []string{"tom1", "tom2", "tom3"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("hello:%s\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
}*/

func main() {
	//show number of goroutine
	fmt.Println(runtime.NumGoroutine())
}
