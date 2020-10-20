package main

import (
	"fmt"
	"strconv"
)

//Exmaple of Multiple return for a function
func returnmultiplevalue(num int) (square int, cube int) {

	square = num * num
	cube = num * num * num
	return
}

//One way of passing value (call by value ) and also the default type(question 3)
func callbyvalue(val string) {
	fmt.Println("The passed value is = ", val)
}

// Question2: call by reference
func callbyreference(val *string) {
	fmt.Println("The passed value is = ", *val)
}

//Pointer on pointer
func pointeronpointer(i *int) {
	var j **int
	j = &i
	fmt.Println("Pointer:= ", *i)
	fmt.Println("Pointer on Pointer := ", **j)
}

//Map creater
func createmap() {
	football := make(map[string]int)
	football["ronalodo"] = 7
	football["messi"] = 10
	fmt.Println("the given map ", football)
}

//type casting
func casting(a int, b float64, c string) { // These are formal parameter
	fmt.Println("Int to float64", float64(a)/5)
	fmt.Println("float64 to int", int(b))
	fmt.Println("Int to String", strconv.Itoa(a))
	conval, err := strconv.ParseInt(c, 10, 64)
	if err == nil {
		fmt.Println("string to int ", conval)
	} else {
		fmt.Println("while casting faced some issue ", err)
	}

}

func main() {

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
	casting(12, 10.91, "11") //the arguments passed are actual parameter
}
