package main

import (
	"fmt"
)

func main() {
	permuteUnique([]int{1,2,2})
}


func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	lens := len(nums)
	recursion(nums, lens, 0, &res)
	return res
}

func recursion(nums []int, lens int, start int, res *[][]int) {
	if start == lens {
		dist := make([]int, lens)
		copy(dist, nums)
		*res = append(*res, dist)
		fmt.Println(dist)
		return
	}

	maps := make(map[int]bool)
	for i:=start; i < lens; i++ {
		if maps[nums[i]] { //这一手我给跪了，我是想不出来好吧，不过想想也简单，就是本轮交换过的数字就不用再交换进行下一轮了，否则传到下一轮的会是个重复的
			continue 
		}
		if i!=start {
			nums[i], nums[start] = nums[start], nums[i]
		}
		recursion(nums, lens, start+1, res)
		if i!=start {
			nums[i], nums[start] = nums[start], nums[i]
		}
		maps[nums[i]] = true

	}

}

