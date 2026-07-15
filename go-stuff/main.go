package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := [][]int{
		{7, 2, 9, 1, 5},
		{12, 4, 8, 15, 3},
		{20, 11, 18, 14, 16},
		{6, 25, 21, 24, 22},
		{19, 13, 17, 10, 23},
	}
	sort2DArr(arr)
	fmt.Println(arr)
}
func sort2DArr(arr [][]int) {
	res := make(chan []int)
	for _, v := range arr {
		go func(ar []int) {
			slices.Sort(ar)
			res <- ar
		}(v)
	}
	for range arr {
		fmt.Println(<-res)
	}
}
