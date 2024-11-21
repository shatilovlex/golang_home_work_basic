package main

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw03/internal/printer"
	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw03/internal/reader"
)

func main() {
	fmt.Println(printer.GenerateChessboard(reader.ConsoleReader()))
}
