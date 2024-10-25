package main

import (
	"github.com/shatilovlex/golang_home_work_basic/hw03_chessboard/internal/printer"
	"github.com/shatilovlex/golang_home_work_basic/hw03_chessboard/internal/reader"
)

func main() {
	printer.PrintChessboard(reader.ConsoleReader())
}
