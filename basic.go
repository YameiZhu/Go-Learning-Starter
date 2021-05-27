package main

import "fmt"

// outside the fuctions, we must use keyword var
// cannot use :=
// the scope of these variables -> inside the package
// var aa = 3
// var ss = "kkk"
// var bb = true
var (
	aa = 3
	bb = true
	ss = "kkk "
)

func variablesZeroValue () {
	var a int
	var s string
	// can not print empty string
	fmt.Println(a, s) 
	// try printf
	// %q -> quotation marks
	fmt.Printf("%d, %q\n", a, s)
}

func variablesInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s)
}

func typeDeduction () {
	var a, b, c, d = 3, 3, true, "def"
	fmt.Println(a, b, c, d)
}

func variablesShorter () {
	// := -> initialize a variable
	// := can only be used inside a fuction
	a, b, c, d := 3, 3, true, "def"
	b = 7
	fmt.Println(a, b, c, d)
}

func main () {
	fmt.Println("Hello world")
	variablesZeroValue()
	variablesInitialValue()
	typeDeduction()
	variablesShorter()
	fmt.Println(aa, bb, ss)
}