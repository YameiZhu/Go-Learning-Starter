package main

import (
	"fmt"
	"io/ioutil"
)
func grade (score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"	
	}
	return g
} 

// func main () {
// 	const filename = "abc.txt"
// 	// contents, err := ioutil.ReadFile(filename)
// 	// if err != nil {
// 	// 	fmt.Println(err)	
// 	// } else {
// 	// 	fmt.Printf("%s\n", contents)
// 	// }

// 	// if 条件里可以赋值
// 	// if 条件里赋值的变量作用域只在if里
// 	if contents, err := ioutil.ReadFile(filename); err != nil {
// 		fmt.Println(err)	
// 	} else {
// 		fmt.Printf("%s\n", contents)
// 	}

// 	fmt.Println(grade(0), 
// 				grade(59), 
// 				grade(61), 
// 				grade(79), 
// 				grade(89), 
// 				grade(99), 
// 				grade(101))
// }