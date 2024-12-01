package distance

import "testing"

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		left     []int
		right    []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 9},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 0},
		{[]int{1, 2}, []int{1, 2, 3}, -1},
		{[]int{5, 10, 15}, []int{10, 5, 20}, 5},
	}

	for _, test := range tests {
		result := CalculateDistance(test.left, test.right)
		if result != test.expected {
			t.Errorf("CalculateDistance(%v, %v) = %d; expected %d", test.left, test.right, result, test.expected)
		}
	}
}
