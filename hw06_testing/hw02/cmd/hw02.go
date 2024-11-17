package main

import (
	"fmt"

	"hw06_testing/hw02/printer"
	"hw06_testing/hw02/reader"
	"hw06_testing/hw02/types"
)

func main() {
	var path string
	var err error

	fmt.Printf("Enter data file path: ")
	_, err = fmt.Scanln(&path)

	if err != nil && err.Error() != "unexpected newline" {
		fmt.Printf("Error: %v", err)
		return
	}

	if len(path) == 0 {
		path = "data.json"
	}

	var staff []types.Employee

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	printer.PrintStaff(staff)
}
