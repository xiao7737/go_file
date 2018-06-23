package main

import (
	 "fmt"
)

func main() {
	fmt.Println("The first is mine！")
}





//byte = uint8  0-255     int8 :-128-127
//rune = int32
//float32/float64   长度4字节，小数位7位/8字节，小数位15位
//零值不等于空值，是申明后的默认值。默认值为0，bool为false，string为空字符串
//int 和 float可以相互转换，bool和int不能
//string 转换 int   需要用到strconv包，strconv.Atoi
//int 转换 string   strconv.Itoa
//常量表达式中只能出现常量，常量组中，如果常量没有赋值，则使用上一行的值
 