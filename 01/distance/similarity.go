package distance

func countOccurances(list []int) map[int]int {
	occurances := make(map[int]int)
	for _, val := range list {
		occurances[val]++
	}
	return occurances
}

func CalcSimilarity(left, right []int) int {
	rightOccurances := countOccurances(right)

	similarity := 0

	for _, val := range left {
		res, ok := rightOccurances[val]
		if ok {
			similarity += res * val
		}
	}

	return similarity
}
