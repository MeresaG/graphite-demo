package main

import (
	"errors"
	"fmt"
)

func main() {

	msg, err := hello_graohite("my message")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(msg)
}

func hello_graohite(msg string) (string, error) {

	if len(msg) < 10 {
		return msg, errors.New("few msg content")
	}

	return msg, nil
}
