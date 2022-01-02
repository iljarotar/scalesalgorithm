package scalesalgorithm

func GetScales(n, k, ref int) (scales [][]int) {
	scales = [][]int{}
	if k == 1 {
		scale := []int{n}
		return append(scales, scale)
	}
	for i := 1; i < n; i++ {
		for _, element := range GetScales(n-i, k-1, ref) {
			scale := append([]int{i}, element...)
			if k != ref || !duplicate(scale, scales) {
				scales = append(scales, scale)
			}
		}
	}
	return scales
}

func duplicate(scale []int, scales [][]int) bool {
	for _, element := range scales {
		if equivalent(scale, element) {
			return true
		}
	}
	return false
}

func equivalent(scale1, scale2 []int) bool {
	if len(scale1) != len(scale2) {
		return false
	}
	n := len(scale1)
	for i := 0; i < n; i++ {
		if identical(scale1, rotate(scale2, i)) {
			return true
		}
	}
	return false
}

func rotate(scale []int, r int) []int {
	rotated := []int{}
	n := len(scale)
	for i := 0; i < n; i++ {
		rotated = append(rotated, scale[(i+r)%n])
	}
	return rotated
}

func identical(scale1, scale2 []int) bool {
	if len(scale1) != len(scale2) {
		return false
	}
	n := len(scale1)
	for i := 0; i < n; i++ {
		if scale1[i] != scale2[i] {
			return false
		}
	}
	return true
}
