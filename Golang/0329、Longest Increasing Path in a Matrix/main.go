func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    
    res := 1
    rows := len(matrix)
    cols := len(matrix[0])
    
    directions := [][]int{{0,-1},{-1,0},{0,1},{1,0}}
    dp := make([][]int,rows)
    
    for i:=0;i<rows;i++{
         dp[i] = make([]int,cols)
    }
           
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
          res = max(res,dfs(matrix,directions,dp,rows,cols,i,j))
        }
    }
    return res
}

func dfs(matrix,directions,dp [][]int,rows,cols,i,j int) int {
    if dp[i][j] != 0 {
        return dp[i][j]
    }
    res := 1
    for _,v := range directions {
        x := i+v[0]
        y := j+v[1]
        if x < 0 || x >= rows || y <0 || y >= cols || matrix[x][y] <= matrix[i][j] {
            continue
        }
        nowLen := 1+dfs(matrix,directions,dp,rows,cols,x,y)
        res = max(res,nowLen)
    }
    dp[i][j] = res
    
    return res
}

func max(a,b int) int {
    if a >= b{
        return a
    }
    
    return b
}


