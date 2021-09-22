package main

import "fmt"

func fizzbuzz(num int) {
	if num%5 == 0 && num%3 == 0  {
		fmt.Println("FizzBuzz")
	} else if num%5 == 0 {
		fmt.Println("Fizz")
	} else if num%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}

func main() {
	num := 13
	fizzbuzz((num))
}

