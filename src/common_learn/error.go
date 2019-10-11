package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"wrong",
	}
}

func main() {
	if err := run(); err != nil { //err为nil表示成功
		fmt.Println(err)
	}
}
