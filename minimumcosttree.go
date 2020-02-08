package minimumcosttree

import (
	"fmt"
	"math"
)

func toKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func minimumMult(arr []int, start, end int, memo map[string][]int) (int, int) {

	if end-start == 2 {
		return arr[start] * arr[end-1], int(math.Max(float64(arr[start]), float64(arr[end-1])))
	}

	if end-start == 1 {
		return 0, arr[start]
	}

	key := toKey(start, end)

	if values, saw := memo[key]; saw {
		return values[0], values[1]
	}

	minimum := math.MaxInt32
	var max int

	for i := start + 1; i < end; i++ {
		leftSum, leftMax := minimumMult(arr, start, i, memo)
		rightSum, rightMax := minimumMult(arr, i, end, memo)
		rootVal := leftMax * rightMax
		minimum = int(math.Min(float64(minimum), float64(rootVal+leftSum+rightSum)))
		max = int(math.Max(float64(max), float64(leftMax)))
		max = int(math.Max(float64(max), float64(rightMax)))

	}

	memo[key] = []int{minimum, max}

	return minimum, max

}

// MctFromLeafValues returns the minimum multiplation value of the tree
func MctFromLeafValues(arr []int) int {
	ret, _ := minimumMult(arr, 0, len(arr), make(map[string][]int, len(arr)))
	fmt.Printf("ret is %d\n", ret)
	return ret
}
