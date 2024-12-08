package lib05

func SolvePrt1(rules []Rule, updates [][]int) int {

	score := 0
	for _, update := range updates {
		updateValid := CheckUpdate(update, rules)
		if updateValid {
			score += update[len(update)/2]
		}
	}
	return score
}

func SolvePrt2(rules []Rule, updates [][]int) int {
	score := 0
	for _, update := range updates {
		updateValid := CheckUpdate(update, rules)
		if !updateValid {
			update = OrderCorrectly(update, rules)
			score += update[len(update)/2]
		}
	}
	return score
}

func CheckUpdate(update []int, rules []Rule) bool {
	for i := 0; i < len(update)-1; i++ {
		pageValid := CheckCombination(update[i], update[i+1], rules)
		if !pageValid {
			return false
		}
	}
	return true
}

func CheckCombination(left, right int, rules []Rule) bool {
	for _, rule := range rules {
		if right == rule.Left && left == rule.Right {
			return false
		}
	}

	return true
}

func OrderCorrectly(update []int, rules []Rule) []int {
	for CheckUpdate(update, rules) == false {
		for i := 0; i < len(update)-1; i++ {
			if !CheckCombination(update[i], update[i+1], rules) {
				update[i], update[i+1] = update[i+1], update[i]
			}
		}
	}
	return update
}
