package main

import (
	"fmt"
	"os"

	"github.com/jorgror/advent-of-code-2024/05/lib05"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules, updates := lib05.Parse(file)

	// fmt.Println("Rules:")
	// for _, rule := range rules {
	// 	fmt.Println(rule.Left, rule.Right)
	// }

	// fmt.Println("Updates:")

	// for _, update := range updates {
	// 	for _, val := range update {
	// 		fmt.Print(val, ",")
	// 	}
	// 	fmt.Println()
	// }

	res := lib05.SolvePrt1(rules, updates)

	fmt.Println("Result1:", res)

	res = lib05.SolvePrt2(rules, updates)
	fmt.Println("Result2:", res)

}
