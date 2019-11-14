func maxProfit(prices []int) int {
    //状态机
    //dp[i][k][0] 为到第i天为止，进行k次交易，手上没有股票时获得的利润
    //dp[i][k][1] 为到第i天为止，进行k次交易，手上有股票时获得的利润
    
    //dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])
    //dp[i][k][1] = max(dp[i-1][k][0],dp[i-2][k][0] -prices[i])
    
    //这里k可以取无限次，因此k的影响可以直接忽略
    //以上式子简化为：
    //dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
    //dp[i][1] = max(dp[i-1][1],dp[i-1][1], dp[i-2][0]-prices[i])
    
    dp_i_0,dp_i_1 := 0,math.MinInt32
    d_pre_0 := 0 //代表dp[i-2][0] 为0  
    for i:=0;i<len(prices);i++{
        tmp := dp_i_0
        dp_i_0  = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1,d_pre_0-prices[i])
        d_pre_0 = tmp
    }
    return dp_i_0
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}
