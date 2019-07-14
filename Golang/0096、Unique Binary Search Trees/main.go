func numTrees(n int) int {
    //这种求总数的，一般都是dp，看完答案，发现是卡特兰数
    //然后又给跪了，跪在地上留下了没有技术的泪水
    //擦完眼泪继续dp，设dp[i]为长度为i的序列共有多少种满二叉树组合
    //固定一个根节点，以当前数为根节点的树的种类 = 左子树数目 * 右子树数目
    //dp[0] = 1 //空数一算一种
    //dp[1] = dp[0] * dp[0] = 1 * 1 = 1
    //dp[2] = dp[0] * dp[1] + dp[1] * dp[0] = 1*1+1*1=2
    //dp[3] = dp[0] * dp[2] + dp[1] * dp[1] + dp[2] * dp[0]
    //转移方程为 dp[n] = (dp[0+i] * dp[n-1-i]) i从0到n-1 所有项之和 
    
    if n <=1 {
        return 1
    }
    
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i:=2; i <= n; i++ {
        tmp := 0
        for j:=0; j < i; j++ {
            tmp += dp[j] * dp[i-1-j]
        }
        dp[i] = tmp
    }
    
    return dp[n]
    
}
