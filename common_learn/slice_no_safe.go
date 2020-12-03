package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		slice = make([]int, 0, 10000)
		n     = 10000
		wg    sync.WaitGroup
		// lock  sync.Mutex
	)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			/*lock.Lock()
			defer lock.Unlock()*/
			slice = append(slice, a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("len:", len(slice))
	// slice并不是线程安全，结果小于10000
	// 原因：append的时候，可能底层的数组会被换掉。换底层数组的时候，
	// 切片同时被多个goroutine拿到，并执行append操作。存在一些goroutine的append结果会被覆盖
	// 使用锁对append加锁或者channel解决线程安全的问题
}
