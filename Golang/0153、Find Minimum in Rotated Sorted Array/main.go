func findMin(nums []int) int {
    lens := len(nums)
    if lens == 0 {
        return lens
    }
    
    //旋转数组用二分吧，中间的数大于第一个元素，最小值在右边，中间值小于第一个元素，最小值在左边，终止条件为nums[mid-1] > numd[mid], nums[mid-1]就是最小值
    left:=0
    right:=lens-1
    
    if nums[0] < nums[lens-1] {
        return nums[0]
    }
    
    for left < right {
        mid := (left + right)/2
        if mid-1 >= 0 && nums[mid] < nums[mid-1] {
            return nums[mid]
        }
        if mid+1 <= lens-1 && nums[mid]>nums[mid+1] {
            return nums[mid+1]
        }
        
        if nums[mid] > nums[0] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    
    return nums[0] //只有一个元素的情况会走到这里
    
    
}
