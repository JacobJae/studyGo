// When project starts, compiler try to find main package and main function
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// Variables
	// go is type language, variables needs to be predefined its type
	// eg.
	const name string = "Jacob"
	fmt.Println("const name is : " + name)
	// const can not be change
	// eg. name = "asd" << not working
	// if you use var, you can change
	// eg.
	var name2 string = "Before"
	fmt.Println("var name2 is : " + name2)
	name2 = "Changed"
	fmt.Println("change name2 is : " + name2)
	// or use short cut
	name3 := "Declare and assign"
	fmt.Println(name3)
	// Declaring with same name is forbidden
	// eg.
	// var name2 string					(forbidden)
	// name3 := "Declare and assign"	(forbidden)
	// When you use short hand assignment( := ), go will determine the type based on value
	// Also ( := ) can not be used outside of func. Only (var name2 string = "Before") will work

	// Functions
	fmt.Println(multiply(2, 3))
	// In go, you can return multiple values at the same time
	totalLength, upperName := lenAndUpper("Jacob")
	fmt.Println("Type of totalLength: ", reflect.TypeOf(totalLength), "Type of upperName: ", reflect.TypeOf(upperName))
	fmt.Println(totalLength, upperName)
	//fmt.Println("Length is: " + totalLength + " name: " + upperName)
	// If you want to ignore 1 value, you can use _
	totalLength2, _ := lenAndUpper("Jacob")
	fmt.Println(totalLength2)
	// Variadic functions (Using unlimited arguments)
	// repeatMe has 1 variable but with ... notation, it will repeat function with all paremeters
	repeatMe("Test1")
	repeatMe("Test2", "Test3")
	repeatMe("Test4", "Test5", "Test6", "Test7")
	// Naked return
	totalLength2, upperName2 := lenAndUpper2("Jacob Jae")
	// Defer, this function has defer element so when this function is done(value has been returned) then defer clause will be executed
	fmt.Println(totalLength2, upperName2)

	// for, range, args
	// In GoLang, loop can be done only with for, no while or others
	superAdd(3, 2, 1, 5, 5, 6)

	// if and else
	// you can declare variable inside of if
	// eg. if koreanAge := age + 2; koreanAge < 18 {}

	// Switch
	// just like if, you can declare variable after switch clause

	// Pointers
	pointer1()
	pointer2()
}

// a int, b int >> a, b int
func multiply(a, b int) int {
	return a * b
}

// Get string and return int, string
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// Naked fucntion, we declare return variable along with return type
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("lenAndUpper2 is done")
	// inside of this function, length and uppercase has been already declared so := won't work
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	fmt.Println(numbers)
}

func pointer1() {
	a := 2
	b := a // copies value of a to b
	a = 10
	fmt.Println("a <- 2, b <- a, A <- 10")
	fmt.Println("Print a b: ", a, b)
	fmt.Println("Adress of a and b", &a, &b)
}

func pointer2() {
	a := 2
	b := &a // copies value of a to b
	a = 10
	fmt.Println("a <- 2, b <- &a, A <- 10")
	fmt.Println("Print a *b: ", a, *b)
	fmt.Println("*b <- 30")
	*b = 30
	fmt.Println("Print a *b: ", a, *b)
	fmt.Println("Adress of a and b", &a, &b)
}
