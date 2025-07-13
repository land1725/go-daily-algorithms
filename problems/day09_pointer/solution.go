package day09_pointer

func pointAdd(p *int) int {
	if p == nil {
		return 0
	}
	return *p + 10
}

func pointMultiply(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	return arr
}
