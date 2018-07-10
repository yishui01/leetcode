bool isValidSudoku(char** board, int boardRowSize, int boardColSize) {
    
    //那就按规则来咯，一条一条过
    int i = 0,j = 0,z = 0, index = 0, row_index = 0, col_index = 0;
    char tmp[9] = {'\0'};
    //检查横排
    for(i = 0; i < boardRowSize; i++)
    {
        memset(tmp, '\0', 9);
        for(j = 0;j < boardColSize; j++)
        {
            if(board[i][j] == '.')continue;
            index = board[i][j] - '0';
            if(tmp[index-1] == '\0') {
                tmp[index-1] = board[i][j];
            } else {
                printf("123");
                return false;
            }
        }
       
    }
    //检查竖排
    for(i = 0; i < boardRowSize; i++)
    {
         memset(tmp, '\0', 9);
        for(j = 0;j < boardColSize; j++)
        {
            if(board[j][i] == '.')continue;
             index = board[j][i] - '0';
            if(tmp[index-1] == '\0') {
                tmp[index-1] = board[j][i];
            } else {
                 printf("456");
                return false;
            }
        }
       
    }
    //检查九宫格
    for(i = 0; i < 9; i++) //九个九宫格，所以是就九次循环
    {
        
         memset(tmp, '\0', 9);
        for(j = 0;j < 3; j++) //横排3格
        {
      
            for(z = 0; z < 3; z++) //竖排3格
            {
                row_index = i/3*3+j; 
                col_index = i%3*3+z;
                 if(board[row_index][col_index] == '.')continue;
                 index = board[row_index][col_index] - '0';
             
                if(tmp[index-1] == '\0') {
                    tmp[index-1] = board[row_index][col_index];
                } else {
                    return false;
                }
            }
           
        }
       
    }
    
    
    
    
    return true;
    
}