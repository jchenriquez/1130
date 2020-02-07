package minimumcosttree

func minimumMult(arr []int, multed int, index int) int {
	if index == len(arr)-1 {
		return multed * arr[index]
	}

	mult1 := minimumMult(arr, multed*arr[index], index+1)
	mult2 := minimumMult(arr, 1, index+1)

	if mult1 < mult2*arr[index] {
		return mult1
	}

	return mult2 * arr[index]
}

// MctFromLeafValues returns the minimum multiplation value of the tree
func MctFromLeafValues(arr []int) int {
	return minimumMult(arr, 1, 0)
}
