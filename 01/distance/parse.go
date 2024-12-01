package distance

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) (left []int, right []int, err error) {

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			return nil, nil, errors.New("invalid input format")
		}
		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return left, right, nil
}
