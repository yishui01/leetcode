func rob(nums []int) int {
    //经典DP，dp[i]为到当前房子为止能得到的最大财宝数目
    //dp[i] = max(dp[i-2]+nums[i],dp[i-1])
    lens := len(nums)
    dp := make([]int,lens)
    if lens ==0{
        return 0
    }
    if lens == 1 {
        return nums[0]
    }
    if lens == 2 {
        return max(nums[0],nums[1])
    }
    dp[0] = nums[0]
    dp[1] = max(nums[0],nums[1])
    
    for i:=2;i<lens;i++{
        dp[i] = max(dp[i-2]+nums[i],dp[i-1])
    }
    return dp[lens-1]
}

func max(a, b int)int{
    if a >=b{
        return a
    }
    return b
}
