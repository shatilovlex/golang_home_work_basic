package reader

import (
	"fmt"
)

const DefaultSize uint = 8

func ConsoleReader() (uint, error) {
	var enteredSize int
	var err error

	fmt.Printf("Enter size: ")
	_, err = fmt.Scanf("%d", &enteredSize)

	if err != nil && err.Error() != "unexpected newline" || enteredSize <= 0 {
		return DefaultSize, nil
	}

	return uint(enteredSize), nil
}
