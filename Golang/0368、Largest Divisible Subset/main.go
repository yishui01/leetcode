func largestDivisibleSubset(nums []int) []int {
    //1、先从小到大排序
    //2、设dp[i]为以i结尾的最大序列的长度，last[i]在最大序列中 nums[i]的上一个元素在nums出现的下标(last数组就是用来回溯路径构成res的，厉害了)
    //3、双重循环迭代dp[i]，nums[i]%nums[j] == 0 && dp[j] > dp[i] {dp[i] = dp[i+1]}
    //4、最后根据last数组，从最大序列的最后一个索引end开始，一直从last找上一个数字的索引，迭代end,真TM精髓

    sort.Ints(nums)
    dp := make([]int,len(nums))
    last := make([]int,len(nums))
    res := make([]int,0)
    max,end := 0,-1

    for i:=0;i<len(nums);i++{
        dp[i] = 1
        last[i] = -1
        for j:=0;j<i && 2*nums[j] <= nums[i];j++{ //2*nums[j] <= nums[i]也TM是精髓，nums[i]要想除尽nums[j]，至少2倍以上，否则免谈除尽
            if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
                
                dp[i] = dp[j]+1
                last[i]=j
            }            
        }
        if dp[i] > max {
            max = dp[i]
            end = i
        }
    }
    if end != -1 {
        for end != -1 {
            res = append([]int{nums[end]},res...)
            end = last[end]
        }
    }

    return res


}
