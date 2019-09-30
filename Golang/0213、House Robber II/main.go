func rob(nums []int) int {
    //中间那部分其实和之前那道题一样的，那就把一头一尾分别去掉，两次遍历，取较大的那个
    //dp[i] = max(dp[i-1]+nums[i],dp[i-2])
    
    lens := len(nums)
    if lens == 1 {
        return nums[0]
    }
    
    //1、去掉头部
    dp1 := make([]int,lens-1+2)
    for i:=1;i<lens;i++{
        dp1[i+1] = max(dp1[i-1]+nums[i],dp1[i]) //从第三个元素开始递推，dp数组前两个是预置的0
    }
    
    //2、去掉尾部
    dp2 := make([]int, lens-1+2)
    for i:=0;i<lens-1;i++{
        dp2[i+2] = max(dp2[i]+nums[i], dp2[i+1])
    }
    
    return max(dp1[lens], dp2[lens])
    
}

func max(a,b int)int{
    if a >= b {
        return a
    }
    return b
}
