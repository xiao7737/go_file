package main

import "fmt"

//生产者
/*func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //生产数据完毕后，主动关闭ch
		wg.Done()
	}()
}

//消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}*/

/*func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将sum送入管道c
}
func main() {
	a := []int{1, 2, 3, 4}
	c := make(chan int)
	go sum(a[:len(a)/2], c) //送入c的值为3=1+2
	go sum(a[len(a)/2:], c) //送入c的值为7=3+4
	x, y := <-c, <-c        //从管道中获取
	fmt.Println(x, y)       //7，3
}*/

//用channel select实现斐波纳契数据
/*func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(10 * time.Millisecond):
			fmt.Println("time out")
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}*/

//channel源码解析
//channel底层数据是hchan struct

//同步阻塞
func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	<-ch
	fmt.Println("all done")
}

/*
96行读取前，会阻塞90行里面的goroutine内容
wait goroutine
start goroutine
exit goroutine
all done
*/

//会输出543210后报panic，报死锁
//main和go里面 都在对ch操作，会死锁，因为同时在写和读
//通道写完后，必须关闭通道，否则range遍历会出现死锁   118行的处理
func main1() {
	ch := make(chan int)
	go func() {
		for i := 5; i >= 0; i-- {
			ch <- i
		}
		//close(ch)
	}()
	for data := range ch {
		fmt.Println(data)
		/*if data == 0 {
			break
		}*/
	}
}
