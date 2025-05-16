package main

import (
	"calculator"
	"fmt"
)

func main() {
	fmt.Println("Hello from Consumer")
	fmt.Println(calculator.Multiply(3, 5))
}
