package main

import (
	"fmt"

	"hw06_testing/hw03/internal/printer"
	"hw06_testing/hw03/internal/reader"
)

func main() {
	fmt.Println(printer.GenerateChessboard(reader.ConsoleReader()))
}
