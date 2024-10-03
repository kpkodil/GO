package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is an even")
	} else {
		fmt.Println("7 is an odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divided by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "< 0")
	} else if num < 10 {
		fmt.Println(num, "consist of 1 number")
	} else {
		fmt.Println(num, "consist of many numbers")
	}
}