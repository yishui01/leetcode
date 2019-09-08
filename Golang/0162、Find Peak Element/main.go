func findPeakElement(nums []int) int {
    //这是大佬答案的二分，主要返回的left指针而不是mid指针，这个...我再想想吧，返回的时候实际上就是left == right的时候，那么这个时候left一定是局部峰值？还有这里的right=mid而left=mid+1
    //这个倒是理解了，right不能==mid+1是因为此时mid可能是峰值，这个不会造成无限循环，因为两个指针相邻的时候，mid计算出来必为左指针，（i+i+1）/2 肯定 == i，无论是奇数在前还是偶数在前。左指针这里是mid+1的，会增加，不会原地踏步
    lens := len(nums)
    if lens <= 1 {
        return lens - 1
    }
    left := 0
    right := lens - 1
    
    var mid int
    
    for left < right {
          mid = (left+right)/2
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid+1
        }
    }
    
    return left
    
}


// func findPeakElement(nums []int) int {
//     //好吧、、、看漏了，要LogN的时间复杂度。。。
//     //这题还有个条件就是说 nums[i] != nums[i+1]
//     //那二分的时候找到mid时候将其与nums[i+1]比对即可，如果大于 nums[i+1] 则舍去右边，反之舍去左边，
//     lens := len(nums)
//     if lens <= 1 {
//         return lens - 1
//     }
//     left := 0
//     right := lens - 1
    
//     var mid int
    
//     for left <= right {
//         mid = (left+right)/2
//         if mid == 0 && nums[mid] > nums[mid+1] || mid == lens -1 && nums[mid] > nums[mid-1] || mid > 0 && mid < lens-1 && nums[mid] > nums[mid-1] && nums[mid] >nums[mid+1]{
//             return mid
//         }
//         if nums[mid] > nums[mid+1] {
//             right = mid
//         } else {
//             left = mid+1
//         }
//     }
    
//     return mid
    
// }


// func findPeakElement(nums []int) int {
//     //这...有什么难的吗，遍历每个元素，比对两边的就行了啊
//     lens := len(nums)
//     if lens <= 1 {
//         return lens - 1
//     }
    
//     for i:=1; i < lens-1; i++ {
//         if nums[i] > nums[i-1] && nums[i] > nums[i+1]{
//             return i
//         }
//     }
    
//     if nums[0] > nums[1] {
//         return 0
//     }
    
//     if nums[lens-1] > nums[lens-2] {
//         return lens-1
//     }
    
//     return -1
  
// }
