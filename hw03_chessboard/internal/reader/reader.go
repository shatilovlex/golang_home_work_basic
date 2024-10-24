package reader

import (
	"fmt"
)

const DefaultSize int = 8

func ConsoleReader() int {
	var enteredSize int
	var err error

	fmt.Printf("Enter size: ")
	_, err = fmt.Scanf("%d", &enteredSize)

	if err != nil && err.Error() != "unexpected newline" || enteredSize <= 0 {
		return DefaultSize
	}

	return enteredSize
}
