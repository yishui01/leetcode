
//参考：一个通用方法团灭 6 道股票问题   https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-wen
//三个状态：天数、最大交易次数、手中是否持有股票
//先列出手中是否持有股票的状态转移方程
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i]) //今天无股票=》昨天无股票今天无操作 / 昨天有今天卖出，卖出不用算k，买入时算k
// dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1]) //今天有股票=》昨天有今天无操作 / 昨天没有今天买入

func maxProfit(prices []int) int {
    //上面已经列出了天数的状态转移方程，那么现在就要穷举出所有状态，对天数i、最大交易次数k进行穷举
    
    lensI := len(prices) //所有天数
    
    if lensI <=0 {
        return 0
    }
    
    lensK := 2 //本题题目要求的最多两次交易
    
    dp := make([][][]int, lensI)  //分配第一维空间
    
    
    
    //开始穷举所有的天数和交易次数
    for i:=0; i < lensI; i++ {
        
        dp[i] = make([][]int, lensK) //分配第二维空间
        
        for k:=0; k < lensK; k++ {
            
            dp[i][k] = make([]int, 2) //分配第三维空间
            
            //这里要处理下初始化状态，一个 i-1 = -1 时，一个是 k-1=-1时，不然数组索引越界，且没有初始值无法递推
            if i-1 == -1 {
               // dp[0][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])   //化简过程可以参考函数后面的注释
               // dp[0][k][0] = max(0, -infinity + prices[i])
                  dp[0][k][0] = 0
                
               // dp[0][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])
               // dp[0][k][1] = max(0 - prices[i], -infinity)
                  dp[0][k][1] = -prices[i]
            } else if k-1 == -1 {
                dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
                dp[i][k][1] = max(-prices[i], dp[i-1][k][1])
            } else {
                dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
                dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])
            }
            
        }
    }
    
    return dp[lensI-1][lensK-1][0]
}
/*
dp[-1][k][0] = 0
解释：因为 i 是从 0 开始的，所以 i = -1 意味着还没有开始，这时候的利润当然是 0 。
dp[-1][k][1] = -infinity
解释：还没开始的时候，是不可能持有股票的，用负无穷表示这种不可能。
dp[i][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0 。
dp[i][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的，用负无穷表示这种不可能。
**/

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}









// 自己手写的递归，就是列举出所有的把数组分为两段的情况，分别计算出两段的最大值相加，找出最终最大值
// func maxProfit(prices []int) int {
//     //卖两次，那就直接硬写把，封装一个函数，找出当天卖出的最高利润
//     res := 0
//     lens := len(prices)
//     for i:=0; i < lens; i++ {
//         tmp := maxProfix(prices[:i]) + maxProfix(prices[i:])
//         if res < tmp{
//             res = tmp
//         }
//     }
//     return res
// }

// func maxProfix(prices []int) int {
//     max := 0
//     lens := len(prices)
//     if lens <2 {
//         return max
//     }
//     min := prices[0]
//     for i:=1; i < lens; i++{
//         if min > prices[i] {
//             min = prices[i]
//         }
//         if max < prices[i] - min {
//             max = prices[i] - min
//         }
//     }
    
//     return max
// }
