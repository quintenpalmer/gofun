package main

import (
	"math/rand"
	"os"
	"fmt"
	"strconv"
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

func rlist(size int) []int {
	list := make([]int,size)
	for i:=0; i < size; i++ {
		list[i] = size - i
	}
	return shuffle(list)
}

func main() {
	size := 26
	if len(os.Args) < 2 {
		fmt.Println("using default size of %d",size)
	} else {
		var err error
		size, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
	}
	to_sort := rlist(size)
	fmt.Println(to_sort)
}
