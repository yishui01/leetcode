func findDuplicate(nums []int) int {
    fast,slow,res,start := 0,0,0,0
    //找到快慢指针相遇的点
    for {
        fast = nums[nums[fast]]
        slow = nums[slow]
        if fast == slow {
            break
        }
    }
    
    for {
        if start == slow {
            return start
        }
        start = nums[start]
        slow = nums[slow]
    }
    
    return res
}
