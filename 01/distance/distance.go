package distance

import "slices"

func CalculateDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		return -1
	}

	distance := 0

	for i := 0; i < len(left); i++ {
		newDistance := left[i] - right[i]
		if newDistance < 0 {
			newDistance = -newDistance
		}
		distance += newDistance
	}

	return distance
}
