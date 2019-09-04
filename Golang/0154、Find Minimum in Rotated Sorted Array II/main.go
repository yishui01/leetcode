func findMin(nums []int) int {
    lens := len(nums)
    if lens == 0 {
        return 0
    }
    if lens == 1 {
        return nums[0]
    }
    if nums[0] < nums[lens-1] {
         return nums[0]
    }
    left := 0
    right := lens-1
    
    //大佬解法,直接比对的是当前右指针的节点
    for left < right {
        mid := left + (right - left) / 2;
        if nums[mid] > nums[right] {
            left = mid+1
        } else if nums[mid] < nums[right]{
            right=mid
        } else {
             right--
        }
    }
    return nums[right]
    
}



// func findMin(nums []int) int {
//    //这题有重复的，左右指针取出的中位数如果和最右一位数字相等右指针就往左移动一位，因为target在左半段，大于最后一位数字left=mid+1，小于最后一位数字right=mid-1
//     //每次拿到mid，都要比较mid和mid-1以及mid和mid+1，当然要去重之后再比较，比较后有任意一个满足条件，就ok
//     lens := len(nums)
//     if lens == 0 {
//         return 0
//     }
//     if lens == 1 {
//         return nums[0]
//     }
    
//     if nums[0] < nums[lens-1] {
//          return nums[0]
//     }
    
//     left := 0
//     right := lens-1
    
//     for left < right {
//         mid := (left + right) / 2
//         tmpMid := mid
//         for mid >1 && nums[mid-1] == nums[mid] {
//             mid--
//         }
//         if mid >= 1 && nums[mid-1] > nums[mid] {
//             return nums[mid]
//         }
        
//         for mid < lens-2 && nums[mid] == nums[mid+1] {
//             mid++
//         }
//         if mid <= lens-2 && nums[mid] > nums[mid+1] {
//             return nums[mid+1]
//         }
        
//         if nums[tmpMid] == nums[lens-1]{
//              right--
//         } else if nums[tmpMid] > nums[lens-1] {
//             left = tmpMid+1
//         } else {
//             right = tmpMid-1
//         }
        
//     }
    
//     return nums[0]
    
// }

