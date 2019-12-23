func combinationSum4(nums []int, target int) int {
    //设dp[i]为target为i时的组合数
    //以 [1,2,3] 4  为例，target为4，第一轮有三种方法可选，可以选1（剩余target为4-1=3），可以选2（剩余target为4-2=2），可以选3（剩余target为4-3=1）
    //因此dp[4] = dp[3]+dp[2]+dp[1]
    //得出dp[i] = dp[i-nums[0]]+dp[i-nums[1]]+dp[i-nums[2]]+...(直到i >= nums[?])

    dp := make([]int,target+1)
    dp[0] = 1
    sort.Ints(nums)
    for i:=1;i<=target;i++{
        for j:=0;j<len(nums);j++{
            if i >= nums[j] {
                dp[i] += dp[i-nums[j]]
            } else {
                break
            }
        }
    }
    return dp[target]

}


/* 记忆dfs
func combinationSum4(nums []int, target int) int {
    caches := make(map[int]int)
    return  helper(nums,target,caches)
}

func helper(nums []int,target int,maps map[int]int)int {
    if (target < 0) { 
        return 0;
    }
    if (target == 0) {
        return 1;
    }
    if val,ok := maps[target];ok {
        return val
    }

    res := 0
    for i:=0;i<len(nums);i++{   
        //这里要改变target，因为我们是用target为key来缓存的，所以不能让target一直保持不变
        res += helper(nums,target-nums[i],maps)
    }

    maps[target] = res

    return res
}
*/
