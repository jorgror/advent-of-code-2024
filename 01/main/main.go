package main

import (
	"fmt"
	"os"

	"github.com/jorgror/advent-of-code-2024/01/lib01"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right, err := lib01.Parse(file)

	if err != nil {
		panic(err)
	}

	distanceRes := lib01.CalculateDistance(left, right)

	fmt.Println("Distance", distanceRes)

	similarity := lib01.CalcSimilarity(left, right)

	fmt.Println("Similarity", similarity)
}
