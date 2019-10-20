package main

import "fmt"

func main() {
	m1 := map[string]string{"one": "one"}
	m2 := map[string]string{"one": "one"}
	m1["one"] = "one"
	m2["one"] = "one"
	if m1["one"] == m2["one"] {
		fmt.Println("m1 is equals to m2")
	}
}
