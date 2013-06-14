package main

import (
	"fmt"
	"os"
	"strconv"
)

func ms(in_list []int) []int {
	in_chan := make(chan []int)
	go rec_ms(in_list,in_chan)
	return <- in_chan
}

func rec_ms(in_list []int, out_chan chan []int) {
	half := len(in_list)/2
	if len(in_list) == 2 && (in_list[0] > in_list[1]) {
		out_chan <- []int{in_list[1],in_list[0]}
	} else if len(in_list) == 1 {
		out_chan <- in_list
	} else {
		in_chan := make(chan []int)
		go rec_ms(in_list[:half],in_chan)
		go rec_ms(in_list[half:],in_chan)
		left, right := <-in_chan, <-in_chan
		out_list := make([]int, 0)
		for i := 0; i < len(in_list); i++ {
			if (len(right) == 0) || (len(left) > 0 && len(right) > 0 && right[0] > left[0]) {
				out_list = append(out_list,left[0])
				left = append(left[:0], left[1:]...)
			} else {
				out_list = append(out_list,right[0])
				right = append(right[:0], right[1:]...)
			}
		}
		out_chan <- out_list
	}
}

func simple_ms(in_list []int) []int {
	half := len(in_list)/2
	if len(in_list) == 2 && (in_list[0] > in_list[1]) {
		return []int{in_list[1],in_list[0]}
	} else if len(in_list) == 1 {
		return in_list
	} else {
		left := simple_ms(in_list[:half])
		right := simple_ms(in_list[half:])
		out_list := make([]int, 0)
		for i := 0; i < len(in_list); i++ {
			if (len(right) == 0) || (len(left) > 0 && len(right) > 0 && right[0] > left[0]) {
				out_list = append(out_list,left[0])
				left = append(left[:0], left[1:]...)
			} else {
				out_list = append(out_list,right[0])
				right = append(right[:0], right[1:]...)
			}
		}
		return out_list
	}
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
	to_sort := make([]int,size)
	for i:=0; i < size; i++ {
		to_sort[i] = size - i
	}
	fmt.Println(to_sort)
	fmt.Println(ms(to_sort))
}
