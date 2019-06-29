func removeDuplicates(nums []int) int {
    //把后面不重复的移动到前面重复的元素的位置覆盖掉
    res := 0
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    
    tmp := nums[0]
    nowTimes := 1
    res++
    
    
    for i :=1; i < lens; i++ {
        if nums[i] == tmp && nowTimes < 2{
            //和之前的一样
                nums[res] = nums[i]
                res++
                nowTimes++
        } else if nums[i] != tmp {
            tmp = nums[i]
            nums[res] = nums[i]
            res++
            nowTimes = 1
        }
    }
    
    return res
}
