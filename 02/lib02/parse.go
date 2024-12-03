package lib02

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) [][]int {

	res := [][]int{}
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			if len(res) <= i {
				res = append(res, []int{})
			}
			res[i] = append(res[i], val)
		}
		i++
	}
	return res
}
