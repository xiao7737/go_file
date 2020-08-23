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
		//虽然函数返回的nil!=nil，但是此时err被:=重新赋值后为nil，
		//if内的err会覆盖函数的申明时的err变量，被重新赋值为nil
		//result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = errors.New("did not work")
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", errors.New("did not work")
}

func main() {
	fmt.Println(DoTheThing(true))  //nil   if的作用域影响和:=重赋值
	fmt.Println(DoTheThing(false)) //nil
}
