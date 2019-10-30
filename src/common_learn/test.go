package main

import "sync"

func main() {
	var rwm sync.RWMutex
	rwm.RLocker()
}
