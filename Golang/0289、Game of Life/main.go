func gameOfLife(board [][]int)  {
    //活细胞转挂 -1
    //挂细胞转活 2
    rows := len(board)
    if rows == 0 {
        return 
    }
    cols := len(board[0])
    if cols == 0 {
        return 
    }
    
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
                count := 0
            if i > 0 && (board[i-1][j] == 1 || board[i-1][j]== -1){
                    count++
                }
            if  i<rows-1 && (board[i+1][j] == 1 || board[i+1][j] == -1){
                    count++
                }
                                 if j > 0 && (board[i][j-1] == 1 || board[i][j-1] == -1){
                    count++
                }
                             if j < cols-1 && (board[i][j+1] == 1 || board[i][j+1] == -1){
                    count++
                }
                                  if i>=1 && j>=1 && (board[i-1][j-1] == 1 || board[i-1][j-1] == -1){
                    count++
                }
                                    if i>=1 && j <= cols-2 && (board[i-1][j+1] == 1 || board[i-1][j+1] == -1){
                    count++
                }
                                        if i<=rows-2 && j >=1 && (board[i+1][j-1] == 1 || board[i+1][j-1] == -1){
                    count++
                }
                  if i<= rows-2 && j<= cols-2 && (board[i+1][j+1] == 1 || board[i+1][j+1] == -1){
                    count++
                }
            
            if board[i][j] == 0 {
                //看是否复活
                if count == 3 {
                    board[i][j] = 2
                }
            } else {
                //看周围是否正好有2-3个活细胞
                if count <2 || count > 3 {
                    board[i][j] = -1
                }
            }
        }
    }
    
    
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
            if board[i][j] == 2 {
                board[i][j] = 1
                continue
            }
            if board[i][j] == -1{
                board[i][j] = 0
            }
        }
    }
    
    return 
    
    
}


