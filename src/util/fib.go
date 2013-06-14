package util

func FibIndex(index int) int {
	old, curr := 0, 1
	for i := 0; i < index; i++ {
		old, curr = curr, old+curr
	}
	return curr
}
