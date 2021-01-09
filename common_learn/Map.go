package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[string]int{
		"apple":  1,
		"orange": 3,
		"beer":   2,
		"banana": 4,
	}

	// put the key in a slice and sort it
	var keys []string
	for key := range map1 {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 此处采用字母排序

	// display the key according the sorted slice
	/*for _, key := range keys {
		fmt.Printf("%s:%v\n", key, map1[key])
	}*/
	for _, v := range map1 {
		fmt.Printf("%d\n", v)
	}
}
