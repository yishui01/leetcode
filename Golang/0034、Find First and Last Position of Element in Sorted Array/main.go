func searchRange(nums []int, target int) []int {
    
    //二分，分到target了就两端各自二分，
    //左边的目标为target的前一个数，等式为 nums[mid] < target && nums[mid+1] == target
    //右边的目标为target的后一个书，等式为 nums[mid] > target && nums[mid-1] == target
    
    lens := len(nums)
    res := []int{-1,-1}
    
    if lens == 0 {
        return res
    }
    
    left := 0
    right := lens-1
    isFind := false
    for left <= right {
        if isFind {
            break
        }
       
        mid := (left + right) / 2
        if (nums[mid] == target) {
            isFind = true
            //找到target了,开始左右二分
            //左边二分
            subLeft := 0
            subRight := mid
            for subLeft <= subRight{
                subMid := (subLeft + subRight) / 2
                if nums[subMid] < target {
                    if subMid + 1 < lens && nums[subMid+1] == target {
                        //这就是左边的答案
                        res[0] = subMid+1
                        break
                    } else {
                        //这里就是锁subMid对应的值还是比target小
                        subLeft = subMid + 1 
                    }
                } else {
                    if target == nums[subMid] {
                         res[0] = subMid
                    }
                    //大于等于 target ,要移动右边
                    subRight = subMid - 1 
                }
            }
            
            //右边二分
            subLeft2 := mid
            subRight2 := lens-1
            for subLeft2 <= subRight2 {
             
                subMid2 :=  (subLeft2 + subRight2) / 2
                if nums[subMid2] > target {
                    if subMid2 - 1 >= 0 && nums[subMid2-1] == target {
                        //这就是右边的答案
                        res[1] = subMid2-1
                        break
                    } else {

                        //这里就是锁subMid对应的值还是比target大
                        subRight2 = subMid2 - 1 
                    }
                } else {
                    if target == nums[subMid2] {
                         res[1] = subMid2
                    }
                    //小于等于 target ,要移动右边
                    subLeft2 = subMid2 + 1 
                }
            }
            
        } else {
            if nums[mid] > target {
                right = mid -1 
            } else {
                left = mid + 1
            }
        }
    }
    
    return res
    
}
