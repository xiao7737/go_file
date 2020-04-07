package design_pattern

import "sync"

type Singleton struct {
}

//方法1：利用sync.once的Do()
//Do方法只会执行1次，Do本质也是双重检测-源码
/*var once sync.Once
var singleInstance *Singleton

func getSingleton() *Singleton {
	once.Do(func() {
		singleInstance = new(Singleton)
	})
	return singleInstance
}*/

//方法2：饿汉式-线程安全，直接创建对象返回
//也可以利用init函数创建对象
/*var singleInstance *Singleton
func init() {
	singleInstance = new(Singleton)
}*/
/*var singleInstance *Singleton = new(Singleton)

func getSingleton() *Singleton {
	return singleInstance
}*/

//方法3：懒汉式-线程安全版
//每次调用都要加锁解锁，性能不高
/*var lock sync.Mutex
var singleInstance *Singleton

func getSingleton() *Singleton {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		singleInstance = new(Singleton)
	}
	return singleInstance
}*/

//方法3：双重检测
//懒汉式（线程安全）的优化版本，减少加锁操作，保证性能且线程安全
var lock sync.Mutex
var singleInstance *Singleton

func getSingleton() *Singleton {
	if singleInstance == nil {
		lock.Lock()
		if singleInstance == nil {
			singleInstance = new(Singleton)
			//singleInstance = &Singleton{}   //效果同上
		}
		lock.Unlock()
	}
	return singleInstance
}
