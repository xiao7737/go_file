package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	println("OK")
}

/*
1.13不会输出OK，因为sleep后，让出了唯一的G，G去运行空的for循环，出不来，一旦for里面有东西，就会输出OK
原因：1.13不会使一个没有主动放弃执行权、且不参与任何函数调用的goroutine被抢占
go里面运行的for循环，没有任何函数调用，所以不会被抢占，一直出不来，阻塞程序，影响gc

go1.14改进：基于信号的真抢占式调度，Go1.14 程序启动时， 会在函数 runtime.sighandler 中注册了 SIGURG 信号的处理函数
runtime.doSigPreempt，在触发垃圾回收的栈扫描时，调用函数挂起 goroutine，并向 M 发送信号，M 收到信号后，
会让当前 goroutine 陷入休眠继续执行其他的 goroutine*/
