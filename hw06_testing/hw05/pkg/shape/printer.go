package shape

import "fmt"

func GetShapeString(s any) string {
	area, err := CalculateArea(s)
	if err != nil {
		return fmt.Sprintf("Ошибка: %v.", err)
	}
	return fmt.Sprintf("%vПлощадь: %v", s, area)
}
