package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) //毫秒
		fmt.Println(s)
	}
}
func main() {
	go say("world") //将会异步打印这两个字符串，顺序不定
	say("hello")
}
