package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign

	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Carrot"}

	fmt.Println(fruitSlice[0:2])

	fmt.Println((len(fruitSlice)))
	fmt.Println(fruitSlice[1:3])

}
