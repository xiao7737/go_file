package design_pattern

import "sync"

//go实现单例模式
//利用sync.once的do方法，do方法只会执行一次
type Singleton struct {
}

var once sync.Once
var singleInstance *Singleton

func getSingleton() *Singleton {
	once.Do(func() {
		singleInstance = new(Singleton)
	})
	return singleInstance
}
