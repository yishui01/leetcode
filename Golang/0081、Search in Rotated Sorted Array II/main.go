func search(nums []int, target int) bool {
    //没写复杂度要求，应该是要logN,那就只能二分了
    //和之前那道题是一样的，直接二分，只不过这里可以出现重复字符，那么有可能mid == right
    //出现这种情况直接改一下左边界或者右边界就行了，因为左边界和右边界都是判断过了的，不是target，
    //所以将左边界往右或者右边界往左移动一位，目的是改变下mid，大不了一位一位判断咯
    
    lens := len(nums)
    if lens == 0 {
        return false
    }
    
    left := 0
    right := lens - 1
    
    for left <= right {
        if target == nums[left] || target == nums[right] {
            return true
        }
        mid := (left + right) / 2
        if target == nums[mid] {
            return true
        }
        //先判断有序列的那一段,如果target不在那一段，就对另一段二分
        if nums[mid] == nums[right] {
            //右半段都是重复的
            left++
        } else if nums[mid] > nums[right] {
            //左半段为升序
            if target < nums[mid] && target >= nums[left] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            //右半段为升序
            if target > nums[mid] && target <= nums[right] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    
    return false
    
}
