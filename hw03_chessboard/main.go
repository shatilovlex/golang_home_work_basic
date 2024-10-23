package main

import (
	"fmt"

	"github.com/shatilovlex/hw03_chessboard/internal/printer"
	"github.com/shatilovlex/hw03_chessboard/internal/reader"
)

func main() {
	var chessboardSize uint
	var err error
	chessboardSize, err = reader.ConsoleReader()

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	printer.PrintChessboard(chessboardSize)
}
