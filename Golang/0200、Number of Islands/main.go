func numIslands(grid [][]byte) int {
    //遍历每一个节点，上下左右扩散，遍历的时候将其改为'-'等特殊符号代表已经遍历过，下次遇到这种就直接跳过
    lens := len(grid)
    if lens == 0 {
        return 0
    }
    cols := len(grid[0])
    res := 0
    for i:=0;i<lens;i++{
        for j:=0;j<cols;j++{
            if grid[i][j] == '-' || grid[i][j] == '0'{
                continue
            }
            res += dfs(grid,i,j,lens,cols)
        }
    }
    
    return res
}

func dfs(grid [][]byte, i,j,rows,cols int) int {
    //上下左右依次扩散，扩散之处全部修改为 '-'
    if grid[i][j] == '0' || grid[i][j] == '-' {
        return 0
    }
    grid[i][j] = '-'
    //上
    if i >0 {
        dfs(grid,i-1,j,rows,cols)
    }
    //下
    if i<rows-1{
        dfs(grid,i+1,j,rows,cols)
    }
    //左
     if j>0{
        dfs(grid,i,j-1,rows,cols)
    }
    //右
     if j<cols-1{
        dfs(grid,i,j+1,rows,cols)
    }
    return 1
}
