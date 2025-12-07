package main

import "fmt"

// 32bit -> [0, 0, 0, 0]
type Struct struct {
	Int16 int
	int16 int
}

func main() {
	// 1. membuat asembly yang tujuannya print nama kalian
	// 2. membuat golang code untuk print nama kalian juga
	//    menggunakan syscall (tidak boleh menggunakan fmt.Print)

	s := "Hello, World"g
	n := 123
	// bool -> 1 byte
	// float -> float64, float32
	// int, int8, int32, ...
	fmt.Println(s, n)
}
