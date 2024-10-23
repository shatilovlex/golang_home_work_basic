package printer

import (
	"fmt"
)

func PrintChessboard(size uint) {
	firstSquareColorIsWhite := firstSquareColorCalculator(size)
	var currentSquareColor string
	for i := uint(0); i < size; i++ {
		if firstSquareColorIsWhite {
			firstSquareColorIsWhite = false
			currentSquareColor = " "
		} else {
			firstSquareColorIsWhite = true
			currentSquareColor = "#"
		}
		for j := uint(0); j < size; j++ {
			fmt.Printf("%s", currentSquareColor)
			currentSquareColor = switchColor(currentSquareColor)
		}
		fmt.Println()
	}
}

func switchColor(color string) string {
	if color == "#" {
		color = " "
	} else {
		color = "#"
	}

	return color
}

// Левая нижняя клетка всегда черная.
// Соответственно для четного кол-ва полей первая клетка доски - белая(пробел), для нечетного - черная(#)

func firstSquareColorCalculator(chessboardSize uint) bool {
	return chessboardSize%2 == 0
}
