package main

import (
	"log"
	time2 "time"
)

func main() {
	now := time2.Now()                             //获取当前时间
	log.Println(now)                               //2020-03-04 13:40:23.874308 +0800 CST m=+0.000133747
	log.Println(now.Format("2006/01/02 15:04:05")) //2020/03/04 13:40:23
	timestamp := now.Unix()                        //时间戳
	timeObj := time2.Unix(timestamp, 0)            //将时间戳转为时间格式
	log.Println(timeObj.Weekday().String())        //星期几

	//Parse 函数可以解析一个格式化的时间字符串并返回它代表的时间
	var layout = "2006-01-02 15:04:05"
	var timeStr = "2019-12-12 15:22:12"
	timeObj1, _ := time2.Parse(layout, timeStr)
	log.Println(timeObj1) //2019-12-12 15:22:12 +0000 UTC
}
