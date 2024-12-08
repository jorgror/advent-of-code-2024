package lib05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Rule struct {
	Left  int
	Right int
}

func Parse(input io.Reader) ([]Rule, [][]int) {
	scanner := bufio.NewScanner(input)
	rules := []Rule{}

	// Parse the rules
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		// Split the line into two parts
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) != 2 {
			panic("invalid input format")
		}

		// Parse the left part
		leftVal, _ := strconv.Atoi(parts[0])
		rightVal, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{Left: leftVal, Right: rightVal})
	}

	// Parse the updates
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		res := []int{}
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			res = append(res, val)
		}
		updates = append(updates, res)
	}

	return rules, updates
}
