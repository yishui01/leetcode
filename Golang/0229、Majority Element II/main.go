func majorityElement(nums []int) []int {
    
    //摩尔投票
    lens := len(nums)
    res := make([]int,0)
    if lens == 0 {
        return res
    }
    
   num1,num2,count1,count2 := 0,0,0,0

    for i:=0;i<lens;i++{
        switch { 
            case nums[i] == num1: 
                count1++
            
            case nums[i] == num2: 
                count2++
            
            case count1 == 0: 
                num1 = nums[i]
                count1++
            
            case count2 == 0: 
                num2 = nums[i]
                count2++
            
            default:
                count1--
                count2--
        }
    }
    
    //这里选出了num1和num2两个候选人，但是这两个候选人可能并不是出现频率超过1/3的，有可能只是恰好出现在最后一个或者后面的部分
    //验证
    count1 = 0
    count2 = 0
    for i:=0;i<lens;i++{
        if nums[i] == num1{
            count1++
        } else if nums[i] == num2{
            count2++
        }
    }
    
    if count1 > lens/3{
        res = append(res, num1)
    }
     if count2 > lens/3{
        res = append(res, num2)
    }
    
    return res
    
}
