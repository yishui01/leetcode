func removeDuplicates(nums []int) int {
    lens := len(nums)
    if lens <= 1 {
        return lens
    }
    
    tmp := nums[0]+1 //这里是为了让tmp和第一个元素不一样，这样就不用在迭代循环中判断i=0的情况
    res := 0
    for i:=0; i<lens; i++ {
        if nums[i] != tmp {
            nums[res] = nums[i]
            tmp = nums[i]
            res++
        }
    }
    return res
}
