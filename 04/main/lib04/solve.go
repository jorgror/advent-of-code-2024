package lib04

import (
	"regexp"
)

func SolvePrt1(puzzle []string) int {

	width := len(puzzle[0])
	height := len(puzzle)

	res := 0

	for n := 0; n < 4; n++ {
		// Left to right
		for _, line := range puzzle {
			res += SolutionsInLine(line)
		}

		// Diagonal top left to bottom right
		for i := -width; i < height; i++ {
			line := ""
			for j := 0; j < width; j++ {
				if i+j < 0 || i+j >= height {
					continue
				}
				line += string(puzzle[i+j][j])
			}
			res += SolutionsInLine(line)
		}
		puzzle = Rotate(puzzle)
	}

	return res
}

func SolutionsInLine(puzzle string) int {
	re := regexp.MustCompile("XMAS")
	return len(re.FindAllStringIndex(puzzle, -1))
}

func Rotate(puzzle []string) []string {
	width := len(puzzle[0])
	height := len(puzzle)

	res := make([]string, width)
	for i := 0; i < width; i++ {
		line := ""
		for j := 0; j < height; j++ {
			line += string(puzzle[height-1-j][i])
		}
		res[i] = line
	}

	// for _, line := range res {
	// 	fmt.Println(line)
	// }
	// fmt.Println()
	return res
}

func SolvePrt2(puzzle []string) int {
	width := len(puzzle[0])
	height := len(puzzle)

	res := 0

	for i := 0; i < height-2; i++ {
		for j := 0; j < width-2; j++ {
			cube := make([]string, 3)
			for k := 0; k < 3; k++ {
				line := ""
				for l := 0; l < 3; l++ {
					line += string(puzzle[i+k][j+l])
				}
				cube[k] = line
			}
			if SolutionInSquare(cube) {
				res++
			}

		}
	}

	return res
}

func SolutionInSquare(puzzle []string) bool {

	found := 0
	for n := 0; n < 4; n++ {
		// Left top to right bottom
		line := ""
		for i := 0; i < 3; i++ {
			line += string(puzzle[i][i])
		}

		if line == "MAS" {
			found++
		}
		puzzle = Rotate(puzzle)
	}

	return (found == 2)
}
