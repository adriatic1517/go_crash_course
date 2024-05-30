package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)
	fmt.Printf("Hex Pointer:0x%x\n", &a)
	fmt.Printf("Int Pointer: %d\n", unsafe.Pointer(&a))
	fmt.Printf("Binary Pointer: %b\n", unsafe.Pointer(&a))

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

}
