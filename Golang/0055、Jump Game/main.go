func canJump(nums []int) bool {
    //好像和上一题是一样的,找到最远的距离，如果最远的距离都不行，那就G了
    current := 0
    lens := len(nums)
    if lens == 0 {
        return true
    }
    for i:=0; i < lens;i++ {
        if i<=current { //这意思是说当前的i是不是在可达到的范围内，
           if nums[i] + i > current {
                current = nums[i] + i
                    }
                if current >= lens-1 {
                return true
             }
        } else {
            return false
        }
    }
    return false
}
