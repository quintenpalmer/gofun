package util

func FibIndex(index int) int {
	old, curr := 0, 1
	for i := 0; i < index-2; i++ {
		old, curr = curr, old+curr
	}
	return curr
}

func FibList(size int) []int {
	old, curr := 0, 1
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = old
		old, curr = curr, old+curr
	}
	return list
}
