func calculateMinimumHP(dungeon [][]int) int {
    /*这题我从左上往右下推的，挂在这个case
    [1, -3, 3]
    [0, -2, 0]
    [-3,-3,-3]
    */
    //结果看了网上大佬的答案是先确定右下角的最终hp为1，然后反向往左上推，我想问这到底怎么想出来的？？？
    //右边和下边的较小值-当前格子  只要保证骑士有往右走或者往下走的可能性即可，因为我们是总右下角推出来的，最终会推到左上角，不会让骑士没有路走
    M := len(dungeon)
    N := len(dungeon[0])
    // hp[i][j] represents the min hp needed at position (i, j)
    // Add dummy row and column at bottom and right side 这...还在右边和下边添加个假的行和列，好吧我的智商是硬伤
    dp := make([][]int,M+1)
    for i:=0;i<M+1;i++{
        dp[i] = make([]int,N+1)
        for j:=0;j < N+1;j++{
            dp[i][j] = math.MaxInt32
        }
    }
    dp[M][N - 1] = 1;
    dp[M - 1][N] = 1;
    
    for  i := M - 1; i >= 0; i--{
        for  j := N - 1; j >= 0; j-- {
             need := min(dp[i + 1][j], dp[i][j + 1]) - dungeon[i][j];
                if need <= 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = need
                }
            }
        }
   return dp[0][0];
}

func min(a,b int)int{
    if a<= b{
        return a
    }
    return b
}


