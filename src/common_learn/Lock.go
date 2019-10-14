package main

import (
	"fmt"
	"sync"
)

//线程安全，确保并发
/*func main() {
	var mut sync.Mutex //声明一个互斥锁
	counter := 0
	for i := 0; i < 5000; i++ {

		go func() { //利用defer来释放锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}*/

//准确等待协程执行时间
func main() {
	var mut sync.Mutex    //声明一个互斥锁
	var wg sync.WaitGroup //申明一个协程等待
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)   //新增一个等待
		go func() { //利用defer来释放锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() //等待完成
		}()
	}
	wg.Wait() //等待
	fmt.Println(counter)
}

//线程不安全版本，结果不唯一
/*func main() {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}*/
