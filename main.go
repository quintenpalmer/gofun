package main

import (
	"fmt"
	"strconv"
	"os"
	"util"
)

func die(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func help() {
	die("Usage : \nms <size> [rand]\nrlist <size>")
}

func main() {
	var size int
	if len(os.Args) < 2 {
		help()
	} else {
		if os.Args[1] == "ms" {
			if len(os.Args) < 3 {
				help()
			} else {
				var err error
				size, err = strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
			}
			var to_sort []int
			if len(os.Args) > 3 && os.Args[3] == "rand" {
				to_sort = util.Rlist(size)
			} else {
				to_sort = make([]int,size)
				for i:=0; i < size; i++ {
					to_sort[i] = size - i
				}
			}
			fmt.Println(to_sort)
			fmt.Println(util.Ms(to_sort))
		} else if os.Args[1] == "rlist" {
			if len(os.Args) < 3 {
				help()
			} else {
				var err error
				size, err = strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println(err)
				}
				to_sort := util.Rlist(size)
				fmt.Println(to_sort)
			}
		}
	}
}
