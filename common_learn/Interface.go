package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}
type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

func main() {
	var writer DataWriter
	f := new(file)
	writer = f
	_ = writer.WriteData("input data source")
}

// output:    WriteData:  input data source

//line19:  将 *file 类型的 f 赋值给 DataWriter 接口的 writer，虽然两个变量类型不一致
//但是 writer 是一个接口，且 f 已经完全实现了 DataWriter() 的所有方法

/*
实现类似于继承的功能
接口是非入侵式，不依赖接口的定义，采用 duck type  鸭子，像鸭子的样子就叫鸭子了 🦆
接口设计优点：让接口和实现者真正解耦，降低项目的耦合度
*/
