func increasingTriplet(nums []int) bool {
    
    //这个方法最后的结果是n1 > n2 > 某一个数时就会返回true
    //能到n2那个分支就代表有个数大于n1并且 小于n2 ，此时赋值给n2，组成n1 < n2
    //如果能到第三个分支，代表有个数能大于n1并且大于n2，由此构成n1 < n2 < x  达成要求，true
    n1,n2 := math.MaxInt32,math.MaxInt32
    for _,v := range nums {
        if n1 >= v {
            n1 = v
        } else if n2 >= v {
            n2 = v
        } else {
            return true
        }
    }

    return false
}

/* 解法二：DP，不符合题目要求
func increasingTriplet(nums []int) bool {
    dp := make([]int,len(nums))
    //dp[i]为到i为止小于当前数字的个数,只要有任意dp[i] >= 2 即为true
    for i:=0;i<len(nums);i++{
        for j:=0;j<i;j++{
            if nums[j] < nums[i] {
                dp[i] = max(dp[i],dp[j]+1)
                if dp[i] >= 2 {
                    return true
                }
            }
        }
    }

    return false
}

func max (i,j int)int{
    if i > j {
        return i
    }
    return j
}
*/
