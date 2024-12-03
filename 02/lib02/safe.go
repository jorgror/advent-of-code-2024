package lib02

func CountSafe(reports [][]int, dampen bool) int {
	count := 0
	for _, report := range reports {
		if dampen {
			for i := range report {
				dampenedReport := make([]int, 0, len(report)-1)
				for j, val := range report {
					if i != j {
						dampenedReport = append(dampenedReport, val)
					}
				}
				if isSafe(dampenedReport) {
					count++
					break
				}
			}
		} else {

			if isSafe(report) {
				count++
			}
		}

	}
	return count
}

func isSafe(r []int) bool {
	if len(r) < 2 {
		return true // no need to check
	}

	dir := r[1] - r[0]
	if dir == 0 {
		return false
	} else if dir > 0 {
		dir = 1
	} else {
		dir = -1
	}

	for i := 0; i < len(r)-1; i++ {
		diff := r[i+1] - r[i]
		diffDir := diff * dir
		if diffDir <= 0 {
			return false
		}
		if diffDir > 3 {
			return false
		}

	}
	return true

}
