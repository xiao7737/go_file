package main

import (
	"fmt"
	"reflect"
)

// 利用反射调用函数
// 太秀了==！

func add(a, b int) int {
	return a + b
}
func main() {
	funcValue := reflect.ValueOf(add)                                      // 将函数包装为反射值对象
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)} // 构造函数参数, 传入两个整型值
	retList := funcValue.Call(paramList)                                   // 反射调用函数
	fmt.Println(retList[0].Int())                                          // 获取第一个返回值, 取整数值
}
