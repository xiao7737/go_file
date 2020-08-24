package main

/*import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string { //对error接口进行实现，方法和error的参数一致即可，鸭子类型
	return fmt.Sprintf("at:%v,%s", e.when, e.what)
}

func errorRun() error { //error是一个内置接口，只有一个函数，Error()
	return &MyError{
		time.Now(),
		"wrong",
	}
}

func main() {
	if err := errorRun(); err != nil { //err为nil表示成功
		fmt.Println(err)
		fmt.Println(errors.New("error"))
	}
}*/

import (
	"errors"
	"fmt"
)

func DoTheThing(reallyDoIt bool) (err error) {
	//var result string
	if reallyDoIt {
		result, err := tryTheThing()
		//if作用域影响，47行返回的err只是37行申明的err的零值nil，改成②处写法，会返回did not work
		//result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = errors.New("did not work") //② return err
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", errors.New("did not work")
}

func main() {
	fmt.Println(DoTheThing(true))  //nil   if的作用域
	fmt.Println(DoTheThing(false)) //nil
}

//简易版的作用域理解
//返回结果是a的零值0，而不是3
/*func main() {
	fmt.Println(test(true))
}

func test(b bool) (a int) {
	if b {
		a := 2
		a += 1
	}
	return a
}*/
