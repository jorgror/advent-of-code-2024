package lib03

func Calculator(vec [][]int) int {
	res := 0
	for _, v := range vec {
		res += v[0] * v[1]
	}
	return res
}
