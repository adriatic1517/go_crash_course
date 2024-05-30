package main

import "fmt"

func main() {
	name, email := "Emad", "abc@def.com"

	var age int64 = 25
	var isTrue = true
	isTrue = false

	fmt.Println(name, email, age, isTrue)
	fmt.Printf("%T\n", isTrue)
}
