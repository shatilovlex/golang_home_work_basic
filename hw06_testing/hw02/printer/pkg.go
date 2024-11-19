package printer

import (
	"fmt"

	"github.com/shatilovlex/golang_home_work_basic/hw06_testing/hw02/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
