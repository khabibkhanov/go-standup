package main

import "fmt"

func oddOrEven (num int) {
	if num%2 == 0 {
		fmt.Println(num, `is odd number`) 
	} else {
		fmt.Println(num, `is even number`) 
	}
}

func main() {
	oddOrEven(24)
}