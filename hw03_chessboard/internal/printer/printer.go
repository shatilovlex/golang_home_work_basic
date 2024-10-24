package printer

import (
	"fmt"
)

func PrintChessboard(size int) {
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if (i+j)%2 == size%2 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}
