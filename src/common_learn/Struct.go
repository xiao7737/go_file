package main

import "fmt"

/*type student struct {
	Name string
	Age  int
}

func main() {
	a := student{
		Name: "tom",
		Age:  16,
	}
	fmt.Println(a)
	ha(a)
	fmt.Println(a)
}
func ha(stu student) {
	stu.Age = 15
	fmt.Println("A", stu)
}*/
/*{tom 16}A {tom 15}{tom 16}*/

//值传递，不会影响原来的值。要改变原有的值， 需要用指针：

/*func main() {            //方式1 推荐方式。将a变成指向结构的指针
	a := &student{         //此处取地址
		Name: "tom",
		Age:  16,
	}
	fmt.Println(a)
	ha(a)
	fmt.Println(a)
}
func ha(stu *student) {      //指针
	stu.Age = 15
	fmt.Println("A", stu)}*/

/*func main() {                         //方式2
	a := student{
		Name: "tom",
		Age:  16,
	}
	fmt.Println(a)
	ha(&a)               //此处取地址
	fmt.Println(a)
}
func ha(stu *student) {
stu.Age = 15
fmt.Println("A", stu)}*/

/*type student struct {          //嵌套结构体1
	Name string
	Age  int
	Info struct {
		Phone, City string
	}
}

func main() {
	a := student{
		Name: "hei",
		Age:  12,
	}
	a.Info.City = "成都"
	a.Info.Phone = "1212121"
	fmt.Println(a)
}*/

//如果结构体1想要使用结构体2的属性，将结构体2的名字放入结构体1。
//使用结构体2的属性时，可以直接调用，不用加结构1的名字
type human struct {
	Sex int
}
type teacher struct {
	Name string
	human
}
type student struct {
	Name string
	human
}

func main() {
	a := teacher{Name: "tom", human: human{Sex: 1}}
	b := student{Name: "tony", human: human{Sex: 0}}
	//修改属性值
	a.Name = "tom2"
	a.Sex = 99 //也可a.human.sex。
	fmt.Println(a, b)
}
