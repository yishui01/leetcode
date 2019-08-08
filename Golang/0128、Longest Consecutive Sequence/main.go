func longestConsecutive(nums []int) int {
    //这题想到一个解法，将数字作为map的key存进去，然后每次来新数字判断新数字的前后数字有没有在map中
    //我开始觉得这个解法很low，然后看了大佬的解法，发现也是这种。。。
    //然后这题改了5次才过，写的飞快，错的也飞快，要改掉这个习惯
    //先想清楚，再写，再肉眼debug，OK吗？写快了是个错的有毛用？？？
    
    lens := len(nums)
    if lens <= 1 {
        return lens
    }
    
    maps := make(map[int]int)
    
    for i := 0; i < lens ; i++ {
       //新数字添加进来后，要把两边的边界值全部带上来
        if _, ok := maps[nums[i]]; ok {
            continue
        }
        flag := false
        if low,okLow := maps[nums[i]-1]; okLow {
            flag = true
            //有比他小的左边界
            maps[nums[i]] = low + 1
            //再把最左边的边界值带上来
            maps[nums[i]-low] = low + 1
        }
        
        if high,okHigh := maps[nums[i]+1]; okHigh {
             flag = true
            //有比他大的右边界
            //此时分两种情况，一种是nums[i]存在，一种是不存在
            
            //先来看存在的
            if _, selfOk := maps[nums[i]]; selfOk{
                //存在代表之前也找到了前半段，现在要连成一个整体
                val := maps[nums[i]] + high
                //改变最左和最右的值
                maps[nums[i]-maps[nums[i]]+1] = val //最左
                maps[nums[i] + 1 + high - 1] = val //最右
                
            } else{
                //不存在前半段
                val := high + 1
                maps[nums[i]] = val
                maps[nums[i]+1 + high - 1] = val
            }
            
        }
        
        if !flag {
             maps[nums[i]] = 1
         }
        
    }
  
    max := 1
    for _,v := range maps{
        if v > max {
            max = v
        }
    }
    
    return max
    
}
