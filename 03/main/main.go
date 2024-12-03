package main

import (
	"io"
	"os"

	"github.com/jorgror/advent-of-code-2024/03/lib03"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	input := string(data)

	vec := lib03.ParseMul(input)

	res := lib03.Calculator(vec)
	println("Result", res)

	preParsed := lib03.PreParse(input)
	vec = lib03.ParseMul(preParsed)
	res = lib03.Calculator(vec)
	println("Result after preparse", res)

}
