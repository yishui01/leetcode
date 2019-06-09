func maxSubArray(nums []int) int {
    //dp,以i结尾的最大子串和=max(i-1的最大子串和+nums[i], nums[i])
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    
    dp := make([]int, 1)
    dp[0] = nums[0]
    max := nums[0]
    pre := nums[0]
    for i:=1; i < lens; i++ {
        if pre > 0 {
            pre = nums[i] + pre
        } else {
            pre = nums[i]
        }
         if max < pre {
                max = pre
          }
    }
    
    return max
    
}
