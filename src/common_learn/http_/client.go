package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	//get content from resp
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
