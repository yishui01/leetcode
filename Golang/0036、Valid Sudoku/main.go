func isValidSudoku(board [][]byte) bool {
    //一个一个来，先横着验证，再竖着验证、再九宫格，难吗？
   
    row := make(map[byte]int)
    col := make(map[byte]int)
    
    grid1 := make(map[byte]int)
    grid2 := make(map[byte]int)
    grid3 := make(map[byte]int)
    
    reset := func(arr map[byte]int){
        for k,_ := range arr {
            arr[k] = 0
        }
    }
    
    for i:=0; i < 9; i++ {
        if i != 0 {
             reset(row)
             reset(col)
        }
      
        if i != 0 && i%3 == 0 {
             reset(grid1)
             reset(grid2)
             reset(grid3)
        }
       
        for j:=0;j < 9; j++ {
            //验证横
            if val,_ := row[board[i][j]]; val == 1 && board[i][j] != '.' {
                return false
            } else {
                row[board[i][j]] = 1
            }
            //验证竖排
               if val,_ := col[board[j][i]]; val == 1 && board[j][i] != '.' {
                  return false
                } else {
                    col[board[j][i]] = 1
                } 
            
            //验证九宫格,根据j的位置决定这个元素是第几个grid的
            if j >=0 && j < 3 {
                if val,_ := grid1[board[i][j]]; val == 1 && board[i][j] != '.' {
                return false
                } else {
                    grid1[board[i][j]] = 1
                }
            } else if j >=3 && j < 6 {
                 if val,_ := grid2[board[i][j]]; val == 1 && board[i][j] != '.' {
                return false
                   
                } else {
                    grid2[board[i][j]] = 1
                }
            } else {
                if val,_ := grid3[board[i][j]];val == 1 && board[i][j] != '.' {
                return false
                } else {
                    grid3[board[i][j]] = 1
                }
            }
            
        }
        
    }
    
    return true
}
