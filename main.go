package main

import (
	"errors"
	"fmt"
)

func DoTheThing(reallyDoIt bool) (err error) {
	//var result string
	if reallyDoIt {
		result, err := tryTheThing()
		//result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = errors.New("did not work")
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", errors.New("did not work")
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
