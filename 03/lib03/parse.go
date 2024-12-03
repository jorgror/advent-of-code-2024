package lib03

import (
	"regexp"
	"strconv"
)

func ParseMul(data string) [][]int {

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	match := re.FindAllStringSubmatch(string(data), -1)
	res := make([][]int, 0, len(match))
	for _, m := range match {
		left, _ := strconv.Atoi(m[1])
		right, _ := strconv.Atoi(m[2])
		res = append(res, []int{left, right})
	}

	return res

}

func PreParse(data string) string {

	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	doMatches := reDo.FindAllStringIndex(data, -1)
	dontMatches := reDont.FindAllStringIndex(data, -1)

	do := 0
	dont := 0

	add := true
	res := ""
	for i := range data {
		if i == doMatches[do][0] {
			add = true
			do++
			if do == len(doMatches) {
				do = 0
			}
		}
		if i == dontMatches[dont][0] {
			add = false
			dont++
			if dont == len(dontMatches) {
				dont = 0
			}
		}
		if add {
			res += string(data[i])
		}
	}
	return res
}
