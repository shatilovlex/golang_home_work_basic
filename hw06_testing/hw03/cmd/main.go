package main

import (
	"fmt"

	printer "hw06_testing/hw03/internal/printer"
	reader "hw06_testing/hw03/internal/reader"
)

func main() {
	fmt.Println(printer.GenerateChessboard(reader.ConsoleReader()))
}
