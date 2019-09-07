func maximumGap(nums []int) int {
	//现在要求的是线性的复杂度，桶排序或者基数排序
	//直接桶排序，上一个桶子的最大值和这一个桶子的最小值就是差值,找出最大的差值即可

	lens := len(nums)

	if lens < 2 {
		return 0
	}

	max := Max(nums)
	min := Min(nums)
	diff := max - min
	if diff == 0 {
		return 0
	}

	bs := 10 //打算分配10个桶子怎么样，这都是我随便分的哈哈哈哈
	buckets := make([][]int, bs)
	//那每个桶子要装的范围就是 (max - min + 1)/5
	ranges := (diff + 1) / bs
	if (diff+1)%bs != 0 {
		ranges++ //这个最后直接+个1是为了向上取整用的
	}
	for i := 0; i < lens; i++ {
		index := (nums[i] - min) / ranges
		flag := 0
		if len(buckets[index]) == 0 {
			buckets[index] = []int{nums[i]}
		} else {
			for j := 0; j < len(buckets[index]); j++ {
				if buckets[index][j] > nums[i] {
					if j == 0 {
						buckets[index] = append([]int{nums[i]}, buckets[index]...)
					} else {
						end := append([]int{}, buckets[index][j:]...)
						buckets[index] = append(append(buckets[index][:j], nums[i]), end...)
					}
					flag = 1
					break
				}
			}
			if flag == 0 {
				buckets[index] = append(buckets[index], nums[i])
			}
		}
	}
	prev := -1
	res := 0
	for i := 0; i < bs; i++ {
		if len(buckets[i]) > 0 {
			for j := 0; j < len(buckets[i]); j++ {
				if prev != -1 {
					if res < buckets[i][j]-prev {
						res = buckets[i][j] - prev
					}
				}
				prev = buckets[i][j]
			}
		}
	}

	return res

}

func Max(nums []int) int {
	lens := len(nums)
	max := nums[0]
	for i := 1; i < lens; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}

func Min(nums []int) int {
	lens := len(nums)
	min := nums[0]
	for i := 1; i < lens; i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	return min
}

