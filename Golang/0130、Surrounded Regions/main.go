func solve(board [][]byte)  {
    //想了下，无解，看答案，一句话，所有不合格的O必定与某个边界O相连
    //思路：遍历四条边，找到O的时候DFS,将相连的O替换成#,最后将剩余的O变成X,将#还原成O
    
    rows := len(board)
    if rows == 0 {
        return 
    }
    cols := len(board[0])
    
    for i:=0; i < rows; i++{
        if i == 0 || i == rows - 1 {
            for j:=0; j < cols; j++ {
                if j == 0 ||  board[i][j-1] != 'O'{ //一行中相连的就不要重复dfs了，前一个O在dfs的时候会把当前的O遍历到的
                    dfs(board, rows, cols, i,j)
                }
            }
        } else {
            dfs(board, rows, cols, i,0)
            dfs(board, rows, cols, i,cols-1)
        }
    }
    
    for i:=0; i < rows; i++ {
        for j:=0; j < cols; j++ {
            if board[i][j] == '#' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
    
}

func dfs(board [][]byte, rows, cols, startRow, startCol int) {
    if board[startRow][startCol] == 'O' {
       board[startRow][startCol] = '#'
       //上
        if startRow > 0 {
            dfs(board,rows, cols, startRow - 1, startCol)
        }
        //下
        if startRow < rows-1{
            dfs(board,rows, cols, startRow + 1, startCol)
        }
        //左
        if startCol >0 {
            dfs(board, rows, cols, startRow, startCol-1)
        }
        //右
        if startCol < cols - 1 {
            dfs(board, rows, cols, startRow, startCol+1)
        }
    }
}
