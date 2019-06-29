func exist(board [][]byte, word string) bool {
    //递归+回溯
    //大哥，递归+回溯超时了，我去，这个不用递归能玩？你告诉我怎么玩，这只能想办法剪枝处理
    //看到网上大佬剪枝，
    //1、我之前用maps保存已经录入过的key，然后超时，学大佬直接将录入过的位置赋值为'.',下一个到这个点的自然就将其排除了（总结，可以添加副作用来剪枝）
    //2、返回结果的时候之前是等上下左右全部找完后再返回，现在是只要有任意一个返回true，那就是true
    allRow := len(board)
    if allRow == 0 {
        return false
    }
    allCol := len(board[0])
    maps := make(map[string]bool)
    isFind := false
    
    for i:=0; i < allRow; i++ {
        for j:=0; j < allCol; j++ {
            if board[i][j] == word[0] {
                if recursion(allRow,allCol,i,j, 0,len(word), word, maps,board,  &isFind) {
                     return true
                }
                
            }
        }
    }
    
    return false
    
}

func recursion(allRow, allCol, row, col,level, wordLens int,word string, maps map[string]bool, board [][]byte, isFind *bool) bool {
    if *isFind == true {
        return true
    }
    
    if level == wordLens {
        *isFind = true
        return true
    }
    
    if row < 0 || col  < 0 || allRow <= row || allCol <= col {
        return false
    }
    
    //首先看当前的字母是否是合法字母
    if word[level] == board[row][col]{
        tmp := board[row][col]
            board[row][col] = '.'
             //进入下一层,下一层有4种情况，上下左右
        if  recursion(allRow,allCol,row+1,col, level+1,wordLens, word, maps,board, isFind) ||
            recursion(allRow,allCol,row,col+1, level+1,wordLens, word, maps,board, isFind) ||
            recursion(allRow,allCol,row,col-1, level+1,wordLens, word, maps,board, isFind) ||
            recursion(allRow,allCol,row-1,col, level+1,wordLens, word, maps,board, isFind) {
                return true
            }
        
         board[row][col] = tmp
           
    }
  
    
    return *isFind
    
}

