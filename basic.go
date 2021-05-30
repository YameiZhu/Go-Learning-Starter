package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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

func euler () {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))
	// fmt.Println(cmplx.Exp(1i * math.Pi) + 1) -> (0+1.2246467991473515e-16i)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1) // (0.000+0.000i)

}

func triangle () {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts () {
	const (
		filename = "abc.text"
		// const can be used as all types if we don't specify type
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums () {
	// normal enum
	const (
		cpp = iota
		_
		python 
		golang 
	)
	fmt.Println(cpp, python, golang)
	// self-inc enum
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	// 1 1024 1048576 1073741824 1099511627776 1125899906842624
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main () {
	// basic
	fmt.Println("Hello world")
	variablesZeroValue()
	variablesInitialValue()
	typeDeduction()
	variablesShorter()
	fmt.Println(aa, bb, ss)
	
	// inner variable type
	// 1. bool, string
	// 2. (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
		// uint: integer without sign, int: integer with sign
	// 3. byte, rune (4 bytes)
	// 4. float32, float64, complex64(实部虚部都是float32), complex128(实部虚部各64)
	euler()
	triangle()

	// const and enum
	consts()
	enums()
}



