func removeElement(nums []int, val int) int {
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    res := 0
    for i:=0; i < lens; i++ {
        if nums[i] != val {
            nums[res] = nums[i]
            res++
        }
    }
    
    return res
}
