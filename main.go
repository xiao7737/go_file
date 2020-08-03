package main

import (
	"fmt"
)

func main() {
	str1 := "ee"
	strArr := []string{"aa", "bb", "cc", "dd"}
	exists := inArray(str1, strArr)
	fmt.Println(exists)
}

func inArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}
