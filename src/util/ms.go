package util

func Ms(in_list []int) []int {
	in_chan := make(chan []int)
	go rec_ms(in_list,in_chan)
	return <- in_chan
}

func rec_ms(in_list []int, out_chan chan []int) {
	half := len(in_list)/2
	if len(in_list) > 2 {
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
	} else if len(in_list) == 2 && (in_list[0] > in_list[1]) {
		out_chan <- []int{in_list[1],in_list[0]}
	} else {
		out_chan <- in_list
	}
}

func simple_ms(in_list []int) []int {
	half := len(in_list)/2
	if len(in_list) > 2 {
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
	if len(in_list) == 2 && (in_list[0] > in_list[1]) {
		return []int{in_list[1],in_list[0]}
	} else {
		return in_list
	}
}
