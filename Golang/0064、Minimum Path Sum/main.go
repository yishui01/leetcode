func minPathSum(grid [][]int) int {
    //dp dp[i][j] 为到达当前位置的最小和
    //dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    cols := len(grid[0])
    for i:=0;i < rows; i++ {
        for j:=0;j < cols; j++ {
            if i == 0 && j == 0 {
                continue
            }
            min := 0
            left := j-1
            top := i-1
            if top < 0 {
                min = grid[i][j-1]
            } else if left < 0 { 
                min = grid[i-1][j]
            } else {
                if grid[i][j-1] >= grid[i-1][j] {
                    min = grid[i-1][j]
                } else {
                    min = grid[i][j-1]
                }
            }
            
            grid[i][j] = min+grid[i][j]
        }
    }
    
    return grid[rows-1][cols-1]
}
