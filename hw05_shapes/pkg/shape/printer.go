package shape

import "fmt"

func PrintShape(s any) {
	area, err := CalculateArea(s)
	if err != nil {
		fmt.Printf("Ошибка: %v.\n", err)
		return
	}
	fmt.Print(s)
	fmt.Printf("Площадь: %v\n\n", area)
}
