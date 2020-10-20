package main

import (
	"fmt"
	"testing"
)

func Test_cases(t *testing.T) {
	var value string // variable declaration

	//Question1: Multiple return
	square, cube := returnmultiplevalue(2)
	fmt.Printf("Square=%v amd Cube=%v\n", square, cube)

	//Part of question 2 and question 3 : (call by value ) and also the default type
	callbyvalue(" Hello Tutree ")

	//Call by references part of question 2
	value = "Hello once again Tutree" //variable initialization Question
	callbyreference(&value)           // we need to pass the reference(address) so we areusing & operator
	//Closure function : which is a type of inline function which is used to access variable in different scope
	x := 1
	Sumvalue := func() int {
		return x + x
	}
	fmt.Println(Sumvalue())

	//Pointer on pointer
	pointeronpointer(&x)

	//Create a map
	createmap()
	//type casting
	casting(12, 10.91, "11")
}
