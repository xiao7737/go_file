package main

import "fmt"

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("is nokia")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
}

/*定义了一个接口Phone，接口里面有一个方法call()
实现类似于继承的功能

在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone。然后调用call()方法*/
