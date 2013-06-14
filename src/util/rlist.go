package util

import (
	"math/rand"
)

func shuffle(in_list []int) []int {
	size := len(in_list)
	out_list := make([]int,0)
	for i := size; i > 0; i-- {
		r := rand.Intn(i)
		out_list = append(out_list,in_list[r])
		in_list = append(in_list[:r],in_list[r+1:]...)
	}
	return out_list
}

func Rlist(size int) []int {
	list := make([]int,size)
	for i:=0; i < size; i++ {
		list[i] = size - i
	}
	return shuffle(list)
}
