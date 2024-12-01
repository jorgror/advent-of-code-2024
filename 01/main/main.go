package main

import (
	"fmt"
	"os"

	"github.com/jorgror/advent-of-code-2024/01/distance"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right, err := distance.Parse(file)

	if err != nil {
		panic(err)
	}

	distanceRes := distance.CalculateDistance(left, right)

	fmt.Println("Distance", distanceRes)

	similarity := distance.CalcSimilarity(left, right)

	fmt.Println("Similarity", similarity)
}
