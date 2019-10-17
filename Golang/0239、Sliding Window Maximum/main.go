// func maxSlidingWindow(nums []int, k int) []int {
//     //解法一：维护一个单调栈,从大到小排列，队首就是最大值
//     lens := len(nums)
//     if k * lens == 0 {
//         return []int{}
//     }
//     if k == 1 {
//         return nums
//     }
    
//     res := make([]int,0)
//     stack := make([]int,0)
    
//     for i:=0;i<lens;i++{
//         for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
//             stack = stack[:len(stack)-1]
//         }
//         stack = append(stack,i)
//         if stack[0] <= i-k {
//             stack = stack[1:]
//         }
//         if i >= k-1 && len(stack) > 0 {
//             res = append(res,nums[stack[0]])
//         }
//     }
    
//     return res
// }

//解法二，动态规划
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}

	size := len(nums)
	g := k - 1

	left := make([]int, size)

	for i := 0; i < size; i++ {
		if i%g == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	right[size-1] = nums[size-1]

	for i := size - 2; i >= 0; i-- {
		if (i+1) % g == 0 {
			right[i] = nums[i]
		} else {
			right[i] = max(nums[i], right[i+1])
		}
	}

	res := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		res[i] = max(right[i], left[i+k-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
