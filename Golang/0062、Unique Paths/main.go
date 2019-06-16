func uniquePaths(m int, n int) int {
    //dfs超时，只能dp
    //每一格只能向右或者向下
    //那么到达当前格子总共有两种方法，一种是从格子上面走下来，一种是从格子左边走到右边
    //dp[i][j] = dp[i-1][j] + dp[i][j-1]
    if m <= 1 || n <= 1 {
        return 1
    }
    dp := make([][]int, m+1)
    for i:=0; i < len(dp); i++ {
        dp[i] = make([]int, n+1)
    }
    for i:=0; i < len(dp); i++ {
        if i == 0 {
            for j:=0; j < n;j++ {
                dp[i][j] = 1
            }
        }
        dp[i][0] = 1
    }
    
    for i := 1; i < m; i++ {
        for j:=1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
    
}


