func lengthOfLIS(nums []int) int {
    //dp
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int,len(nums))
    for i:=0;i<len(dp);i++{
        dp[i] = 1
    }
    res := 0
    for i:=0;i<len(nums);i++ {
        for j:=0;j<i;j++{
            if nums[j] < nums[i] {
                dp[i] = max(dp[i],dp[j]+1)
            }
        }
        res = max(res,dp[i])
    }
    
    return res
}

func max(a,b int) int {
    if a >= b {
        return a
    }
    
    return b
}


