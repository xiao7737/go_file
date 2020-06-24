package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	<-done
}

func main() {
	go aGoroutine()
	done <- true
	println(msg)
}
