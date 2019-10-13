package main

import "fmt"

type Phone interface {
	call()
}
type IPhone struct {
}

func (Phone *IPhone) call() {
	fmt.Println("is iphone")
}
func main() {
	var phone Phone
	phone = new(IPhone)
	phone.call()
}

/*定义了一个接口Phone，接口里面有一个方法call()
实现类似于继承的功能
在main函数里面定义了一个Phone接口类型变量，并分别为之赋值为IPhone。然后调用call()方法

接口是非入侵式，不依赖接口的定义，采用 duck type  鸭子，像鸭子的样子就叫鸭子了 🦆
*/
