package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		temp := n % 2
		result =  strconv.Itoa(temp) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// equivalent to while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever () {
	// infinite loop  
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 110
		convertToBin(13), // 1101
		convertToBin(745463),
		convertToBin(0),
	)

	printFile("abc.txt")
}