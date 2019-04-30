func threeSumClosest(nums []int, target int) (res int) {
    //最接近的三个数之和，这TM不还是和上一题一样吗，先确定一个，再找剩下的两数之和等于前一个数的相反数
    //只不过这里还不用去重，直接记录最接近的就行，简单
    lens := len(nums)
    if lens < 3 {
        return 0
    }
    sort.Ints(nums)
    res = nums[0]+nums[1]+nums[2]
    if lens == 3{
        return res
    }
    min := math.Abs(float64(res - target))
   
    for i:=0;i < lens-2;i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }  
        l := i+1
        r := lens-1
        for l < r {
            tmp := nums[l] + nums[r] + nums[i]
            if tmp == target {
                return target
            } else if tmp > target {
                if sum:=math.Abs(float64(tmp-target));sum< min {
                    min = sum
                   
                    res = tmp
                }
                for r > l && nums[r] == nums[r-1] {
                    r--
                }
                r--
            } else {
              
                 if sum:=math.Abs(float64(tmp-target));sum< min {
                    min = sum
                    res = tmp
                }
                for r > l && nums[l] == nums[l+1] {
                    l++
                }
                l++
            }
        }
    }
    
    return res
}