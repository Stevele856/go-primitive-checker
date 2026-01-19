package input

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/prime-checker/controller"
)

func InputValue(scanner *bufio.Scanner, prompt string) (int, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return 0, fmt.Errorf("cannot read primitive value")
	}
	input := strings.TrimSpace(scanner.Text())

	return controller.Validate(input)
}
