func minSubArrayLen(s int, nums []int) int {
    //用滑动窗口吧，左右指针,右指针不断拉大，拉到足够了就记录res长度，接下来就缩小左指针，缩小一次判断一次看是否符合条件，是就更新，不符合条件就继续拉大右指针
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    left , right := 0, 0
    res := 0
    sum := 0
    for i:=0;i<lens;i++{
        sum+=nums[right]
        if sum  >= s {
                tmpRes := right-left+1 //当前合理的长度，但是还未缩小
                for left < right { //尝试缩小范围
                        tmp := nums[left]
                        if sum - tmp >= s {
                            left++
                            sum -= tmp
                            tmpRes--
                        } else {
                            break
                    }
                }
            if res == 0 || tmpRes < res {
                res = tmpRes
            }
            
        }
        right++
    }
    
    return res
    
}
