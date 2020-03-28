package main

import "fmt"

//使用map和空接口设计一个可以保存任意值的字典
//key和value都是interface{}
type Dictionary struct {
	data map[interface{}]interface{}
}

//创建一个字典
func Create() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

//设置键值
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

//根据键值获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

//清空数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

//遍历字典返回kv--利用回调
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func main() {
	// 创建字典实例
	dict := Create()
	// 添加游戏数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)
	// 获取值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)
	// 遍历字典元素
	dict.Visit(func(key, value interface{}) bool {
		// 将值转为int类型，并判断是否大于40
		if value.(int) > 40 {
			// 输出很贵
			fmt.Println(key, "is expensive")
			return true //继续遍历
		}
		// 默认都是输出很便宜
		fmt.Println(key, "is cheap")
		return true //继续遍历
	})
}
