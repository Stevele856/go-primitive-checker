package utils

import (
	"fmt"
	"strconv"
)

func Validate(n string) (int, error){
	value, err := strconv.Atoi(n)
	if err != nil {
		return 0, fmt.Errorf("n must be a valid interger")
	}

	if value < 2 {
		return 0, fmt.Errorf("%d is not a prime number", value)
	}
	return value, nil
}