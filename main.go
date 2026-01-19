package main

import (
	"bufio"
	"fmt"
	"os"

	prime "github.com/prime-checker/model"
	input "github.com/prime-checker/view"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Prime checker")
	fmt.Println("============")

	for {
		value, err := input.InputValue(scanner, "Enter value: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if prime.CheckPrime(value) {
			fmt.Println(value, "is a prime number")
		} else {
			fmt.Println(value, "is not a prime number")
			continue
		}
		return
	}
}
