package reader

import (
	"fmt"
	"strconv"
)

const DEFAULT_SIZE = 8

func ConsoleReader() (uint, error) {
	var enteredSize string
	var err error

	fmt.Printf("Enter size: ")
	_, err = fmt.Scanln(&enteredSize)

	if err != nil && err.Error() != "unexpected newline" {
		return DEFAULT_SIZE, nil
	}
	if enteredSize == "" {
		return DEFAULT_SIZE, nil
	}

	return parseEnteredSize(enteredSize)
}

func parseEnteredSize(enteredSize string) (uint, error) {
	var size uint64
	var err error
	size, err = strconv.ParseUint(enteredSize, 10, 32)

	if err != nil {
		return 0, err
	}

	return uint(size), nil
}
