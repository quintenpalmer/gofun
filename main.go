package main

import (
	"fmt"
	"strconv"
	"os"
	"util"
)

func help() {
	fmt.Println("Usage : \n  ms     <size> [rand]\n  rlist  <size>\n  fib    <size> [list]\n  pascal <size>")
	os.Exit(1)
}

func parseInt(args []string) int {
	if len(args) < 3 {
		fmt.Println("Must supply a size")
		help()
	}
	i, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
		help()
	}
	return i
}

func main() {
	var size int
	if len(os.Args) < 2 {
		help()
	} else {
		if os.Args[1] == "ms" {
			size = parseInt(os.Args)
			var to_sort []int
			to_sort = make([]int,size)
			for i:=0; i < size; i++ {
				to_sort[i] = size - i
			}
			if len(os.Args) > 3 && os.Args[3] == "rand" {
				to_sort = util.Shuffle(to_sort)
			}
			fmt.Println(to_sort)
			fmt.Println(util.Ms(to_sort))
		} else if os.Args[1] == "rlist" {
			size = parseInt(os.Args)
			to_sort := util.Rlist(size)
			fmt.Println(to_sort)
		} else if os.Args[1] == "fib" {
			size = parseInt(os.Args)
			if len(os.Args) > 3 && os.Args[3] == "list" {
				fmt.Println(util.FibList(size))
			} else {
				fmt.Println(util.FibIndex(size))
			}
		} else if os.Args[1] == "pascal" {
			size = parseInt(os.Args)
			triangle := util.Pascal(size)
			size = len(triangle)
			for i, line := range triangle {
				for j := 0; j < size - i; j++ {
					fmt.Printf("   ")
				}
				for _, value := range line {
					fmt.Printf(" %0-5d",value)
				}
				fmt.Printf("\n")
			}
		}
	}
}
