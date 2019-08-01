func minimumTotal(triangle [][]int) int {
    //到达每一个格子只有两种可能，那么dp数组就可以建立了
    //dp[i] = 上一层的一个数或者两个数里较小的那个累加当前格子的数
    //直接用triangle作为dp数组
    lens := len(triangle)
    if lens == 0 {
        return 0
    }
    
    for i:=1; i < lens; i++ {//第一层直接忽略
        tmpLen := len(triangle[i])
        for j:=0; j < tmpLen; j++ {
            if j == 0{
                triangle[i][j] = triangle[i-1][j] + triangle[i][j]
            } else if j == tmpLen-1{
                triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
            } else { 
                min := Min(triangle[i-1][j-1],triangle[i-1][j])
                triangle[i][j] = triangle[i][j] + min
            }
        }
    }
    
    //找到最后一层中最小的那个数
    res := triangle[lens-1][0]
    for i:=0; i < len(triangle[lens-1]); i++ {
        if res > triangle[lens-1][i] {
            res = triangle[lens-1][i]
        }
    }
    
    
    return res
    
}

func Min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}
