func canPartition(nums []int) bool {
    //跪了，这题是01背包的变化版
    sum := 0
    for _,v := range nums {
        sum += v
    }

    if sum % 2 != 0 {
        return false
    }

    //设置dp[i][j] 为在[0,i]的范围内，是否满足一部分数的和 == j  ，满足为true，不满足为false
    //dp[i+1][j] == dp[i][j] || dp[i][j-nums[i]]

    target := sum / 2
    dp := make([][]bool,len(nums))
    for i:=0;i<len(nums);i++{
        dp[i] = make([]bool,target+1)
        dp[i][0] = true //把第一列的数字全部初始化为 true（dp[0][0],dp[1][0],dp[2][0]...），其他保持false
    }

    for i:=0;i<len(nums);i++{
        for j:=1;j<target+1;j++{
            if i == 0 {
                if j == nums[i] {
                    dp[i][j] = true
                }
                continue
            }

            if dp[i-1][j] {
                dp[i][j] = true
            } else if j>=nums[i] && dp[i-1][j-nums[i]] {
                dp[i][j] = true
            }
        }
    }
    return dp[len(nums)-1][target]

}

/*
还可以先确定一个target，然后倒序之后dfs，巧妙的是判断头部元素和target的大小，如果target还小于头部元素，这个数组怎么分都不能分成两半 && 每一半的和 == target了，因为头部元素太大了
func canPartition(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	n := sum/2
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) //倒序

	return dfsForPartition(nums, 0, n)
}

func dfsForPartition(nums []int, index, target int) bool {
	if index >= len(nums) || nums[index] > target {  //就这一个元素就比一半还大的话，那剩下的加起来都不如这一个，那没办法，直接false
		return false
	}

	for i := index; i < len(nums); i++ {
		if target - nums[i] == 0 {
			return true
		}
		if  dfsForPartition(nums, i+1, target-nums[i]) {
			return true
		}
	}
	return false
}*/
