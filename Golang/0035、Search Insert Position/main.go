func searchInsert(nums []int, target int) int {
    //这个也可以用二分，哈哈哈，直接找到target或者 满足以下条件
    //nums[mid] > target && nums[mid - 1] < target
    //nums[mid] < target && nums[mid + 1] > target
    
    lens := len(nums)
    if lens == 0 || nums[0] > target{
        return 0
    }
    if nums[lens-1] < target {
        return lens
    }
    
    left:= 0
    right := lens-1
    
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            if mid - 1 >= 0 && nums[mid - 1] < target {
                return mid
            }
            right = mid - 1
        } else {
            if mid + 1 < lens && nums[mid + 1] > target {
                return mid + 1
            }
            left = mid + 1
        }
    }
    
    return 0
}
