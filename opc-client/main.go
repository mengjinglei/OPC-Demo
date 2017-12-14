package main

import (
	"C"
	"fmt"
)

func main() {
	client := NewClient("localhost", 16664)
	if client == nil {
		fmt.Printf("Create client fail")
	}
	fmt.Println("create client successfully")
	value := client.ReadSingleValue(1, "the.answer")
	fmt.Printf("the value is %v\n", value)
	client.BrowserNode(0)

	fmt.Println("ping server")
	err := client.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
}
