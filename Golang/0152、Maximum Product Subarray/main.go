func maxProduct(nums []int) int {
    //这个是最大子串积，要考虑负号相乘的情况，有可能最小的负数乘一个负数之后突然就变成最大的了，所以要保留最小的，也要保留最大的，遇到负号就交换
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    max,imax,imin := nums[0],1,1
    
    for i:=0;i<lens;i++ {
        if nums[i] < 0 {
            imax,imin = imin,imax
        }
        imax = Max(nums[i]*imax, nums[i])
        imin = Min(nums[i]*imin, nums[i])
        max = Max(imax,max)
    }
    
    return max
}


func Max(a,b int) int {
    if a >= b {
        return a
    }
    return b
}

func Min(a,b int) int {
    if a >= b {
        return b
    }
    return a
}
