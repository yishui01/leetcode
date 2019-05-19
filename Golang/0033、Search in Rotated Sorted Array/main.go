func search(nums []int, target int) int {
    //这种O(logN)的直接二分，不过这里的二分要考虑是对左半段二分还是右半段二分，
    //要点在于先找出中间的点，判断这个点和最右边的点谁大，如mid <right 代表以中点为起点的右边半段是升序
    //如果mid>right 代表以中点为尾部的左半段为升序，接下来可以用两段的首尾判断target所在区域，进行二分搜索
    lens := len(nums)
    res := -1
    if lens == 0 {
        return -1
    }
    left := 0
    right := lens-1
    
    //排除不可能的情况，头一个数肯定是和尾部是连着的，中间不会有间隙，如果target在这个间隙中，那就废了
    if target > nums[lens-1] && target < nums[0] {
        return -1
    }
    
    
    for left <= right {
        mid := (left+right)/2
        if target == nums[mid] {
            return mid
        }
        if target == nums[left] {
            return left
        }
        if target == nums[right] {
            return right
        }
        if nums[mid] < nums[right] {
            //右半段为升序
            if target >= nums[mid] && target <= nums[right] {
                //target在右半段
                left = mid+1
            } else {
                right = mid-1
            }
        } else {
            //左半段为升序
            if target >= nums[left] && target <= nums[mid] {
                //在左半段
                right = mid-1
            } else {
                left = mid+1
            }
        }
        
    }
    
    return res
    
}
