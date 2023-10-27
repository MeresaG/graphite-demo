package main

import "fmt"

func main() {
	
	msg, err := hello_graohite("my message")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(msg)
}

func hello_graohite(msg string) (string, error) {
	return msg, nil
}
