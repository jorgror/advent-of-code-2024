package main

import (
	"fmt"
	"os"

	"github.com/jorgror/advent-of-code-2024/02/lib02"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := lib02.Parse(file)

	count := lib02.CountSafe(reports, false)
	fmt.Println("Safe reports", count)

	count = lib02.CountSafe(reports, true)
	fmt.Println("Safe reports with dampening", count)
	// for _, report := range reports {
	// 	for j, val := range report {
	// 		if j > 0 {
	// 			print(" ")
	// 		}
	// 		print(val)
	// 	}
	// 	println()
	// }

}
