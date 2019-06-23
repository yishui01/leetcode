func climbStairs(n int) int {
    //简单dp，爬台阶
    //到达当前台阶的方式有两种(n-1阶跳上来的，或者是n-2阶跳上来的)
    //dp[i] = dp[i-1] + dp[i-2]
    if n == 0 {
        return 0
    }
    
    dp := make([]int, n + 1)
   
    for i := 0; i < n+1; i++ {
        if i < 3 {
            dp[i] = i
        } else {
           dp[i] = dp[i-1] + dp[i-2] 
        }
        
    }
    
    return dp[n]
}
