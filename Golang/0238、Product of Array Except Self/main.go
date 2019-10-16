func productExceptSelf(nums []int) []int {
    lens := len(nums)
    res := make([]int,lens)
    left,right := 1,1
    for i:=0;i<lens;i++{
        res[i] = 1
    }
    for i:=0;i<lens;i++{
        res[i] *= left
        left*=nums[i]
        
        res[lens-1-i] *= right
        right *= nums[lens-1-i]
    }
    
    return res
}
