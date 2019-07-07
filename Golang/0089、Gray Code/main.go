func grayCode(n int) []int {
    if n <= 0 {
        return []int{0}
    }
    
    //dp 从低位开始，往高位递推
    dp := make([]int, 1<<uint(n))
    dp[0] = 0
    dp[1] = 1
    l := 2 //从第二位开始，2为10，1所在的位置是第二位，所以是2开始
    for i:=0; i < n-1; i++ {
        //每一次外层循环代表当前位的全部变化
        k := 0
        for j:=0; j < l; j++ {
            dp[l+j] = dp[l-1-k] + l
            k++
        }
        l *= 2
    }
    return dp
}


