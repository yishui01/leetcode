func coinChange(coins []int, amount int) int {
    //典型DP，设dp[i]为组成金额为i所需要的最小硬币
    //以coins=[1,2,5],amount=11为例，初始化dp数组为dp[12] //从dp[1]开始
    //dp[1],dp[2],dp[5] = 1  其余都为-1
    //思考dp[3] = ?
    //dp[3]代表金额为3，那么金额为3可以回退成哪些组成？coins有1,2,5，那么金额3可以减去1也可以减去2，得到dp[2]和dp[1],那么金额3为dp[3] = min(dp[3-1]+1,dp[3-2]+1)
    //那假设coins = [2,4,6] amount=11呢  考虑dp[1] ? 无法组成，返回-1，dp[3]可以化解成dp[3-2]+1 但是由于dp[1]为-1，不合法，因此dp[3]也依然为-1
    
    if amount == 0 || len(coins) == 0 {
        return 0
    }
    dp := make([]int,amount+1)
    for i:=1;i<len(dp);i++{
        for j:=0;j<len(coins);j++{
            if coins[j] == i {
                dp[i] = 1
                break
            } else if coins[j] < i && dp[i-coins[j]] != 0 {
                 if  dp[i-coins[j]]+1 < dp[i] || dp[i] == 0 {
                     dp[i] = dp[i-coins[j]]+1
                 }
            }
        }
    }
    if dp[amount] == 0 {
        return -1
    }
    return dp[amount]
}

