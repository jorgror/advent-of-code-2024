package main

import (
	"os"

	"github.com/jorgror/advent-of-code-2024/06/lib06"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	board, guard := lib06.Parse(file)

	res := lib06.SolvePrt1(board, guard)
	board.Print()

	println("Result1:", res)

	res = lib06.SolvePrt2(board, guard)
	println("Result2:", res)
}
