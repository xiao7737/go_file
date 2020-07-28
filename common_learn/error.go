package main

import (
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
}

/*package main

import (
"errors"
"fmt"
)

func DoTheThing(reallyDoIt bool) (err error) {
	//var result string
	if reallyDoIt {
		result, err := tryTheThing()
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
	fmt.Println(DoTheThing(true))    nil   因为50行返回的err和44行的err不是同一个err，if的作用域影响
	fmt.Println(DoTheThing(false))   nil
}*/
