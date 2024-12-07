package main

import (
	"os"

	"github.com/jorgror/advent-of-code-2024/04/main/lib04"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	puzzle := lib04.Parse(file)

	res := lib04.SolvePrt1(puzzle)
	println("Result1:", res)

	res2 := lib04.SolvePrt2(puzzle)
	println("Result2:", res2)

}
