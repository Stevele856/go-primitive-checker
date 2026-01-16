package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/prime-checker/services"
	"github.com/prime-checker/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Prime checker")
	fmt.Println("============")

	for {
		value, err := utils.InputValue(scanner, "Enter value: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if services.IsPrime(value) {
			fmt.Println(value, "is a prime number")
		} else {
			fmt.Println(value, "is not a prime number")
			continue
		}
		return
	}
}
