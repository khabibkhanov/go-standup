package main

import (
  "fmt"
)

func main() {
  var x = fib(10)
  fmt.Println(x)
}

func fib(num int) int {
if num == 0 || num == 1{
  return num
}

return fib(num - 2) + fib(num - 1)
}