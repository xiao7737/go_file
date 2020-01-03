package main

import "fmt"

func main() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}
func doubleScore(source float32) (score float32) {
	/*②*/ defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	/*③*/ return /*①*/ source * 2
}

// 0，40，50
// return 不是原子操作，分为返回值复制和return两个步骤
// 函数顺序为 ①②③
