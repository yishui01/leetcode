func moveZeroes(nums []int)  {
    //快慢指针，快指针一直前进，慢指针指向非0元素的当前索引，快指针遇到非0时，将慢指针索引赋值为非0，慢指针++，快指针++。一直保持快慢指针之间全部是0
    quick,slow := 0,0
    
    for quick < len(nums){
        if nums[quick] != 0 {
            nums[quick],nums[slow] = nums[slow],nums[quick]
            slow++
        }
        quick++
    }
}
