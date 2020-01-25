func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
    flag := false
	for _, n := range nums {
		if n == max1 || n == max2 { // 过滤掉前两大的重复数据
			continue
		}
		switch {
		case max1 < n:
			max3, max2, max1 = max2, max1, n
		case max2 < n:
			max3, max2 = max2, n
		case max3 <= n: //这里加个等于，确保IntMin能进来
			max3 = n
            flag = true
		}
	}


	// 说明没有第三大的数
	if max3 == math.MinInt64 && !flag { //Bug fix,数组中出现IntMin的时候不能直接通过max == math.MinInt64来判断，需要判断flag是否赋值， max3的值为min && flag也没动过，那才是没有第三大的数
		return max1
	}

	return max3
}



/*
func thirdMax(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    //遍历三次...
    max1 := nums[0]
    max2 := 0
    max3 := 0
    min := nums[0]
    //先找最大的
    for _,v := range nums {
        if v > max1 {
            max1 = v
        }
        if v < min {
            min = v
        }
    }

    max2,max3 = min,min //先赋值为最小的
    //找第二大的
    for _,v := range nums {
        if v > max2 && v != max1 {
            max2 = v
        }
    }

    if max2 == max1 {
        return max1
    }

    //找第三
    for _,v := range nums {
        if v > max3 && v != max1 && v != max2 {
            max3 = v
        }
    }

    if max3 != max2 {
        return max3
    }

    return max1
    
}
*/
