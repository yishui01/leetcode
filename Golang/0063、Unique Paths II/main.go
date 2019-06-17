func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    //一样的dp，左或者上有障碍物的点做处理就行了
    //dp[i][j] = dp[i-1][j] + dp[i][j-1]
    
    //dp还是不够熟练，差点绕晕了，实际上不要想那么多，只管当前格子的状态转移方程就行了，当前格子只需要考虑左上两个，其他的不用管，dp太菜了我，╮(╯▽╰)╭
    
    rows := len(obstacleGrid)
    if rows == 0 || obstacleGrid[0][0] == 1 {
        return 0
    }
    cols := len(obstacleGrid[0])
    obstacleGrid[0][0] = 1 //起点设为1，递推，除了起点以外的1都是墙
    for i:=0;i < rows; i++ {
        for j:=0; j < cols; j++ {
            if i == 0 && j == 0 {
                continue
            }
            if obstacleGrid[i][j] == 1 { //墙
                obstacleGrid[i][j] = 0
                continue
            }
            col := j-1
            if j == 0 {
                col = 0
            }
            row := i - 1
            if i == 0 {
                row = 0
            }
            left := obstacleGrid[i][col]
            top := obstacleGrid[row][j]
            
            obstacleGrid[i][j] =  left + top
        }
    }
    return obstacleGrid[rows-1][cols-1]
}