type Node struct{
    Next [26]*Node
    Word string
}


//这题TMD是要按照需要找的单词来构建前缀树，然而一开始我是用的board来构建前缀树，不论怎样都MLE，后来看了答案才知道要用word来构建，再用board根据前缀树去dfs

func findWords(board [][]byte, words []string) []string {
    rows := len(board)
    res := make([]string,0)
    if rows == 0 {
        return res
    }
    cols := len(board[0])
    
    //先按照要查的单词构造一个前缀树，而不是board来构建不然会MLE
    node := Node{Next:[26]*Node{}}
    for i:=0;i<len(words);i++{
       getTree(&node, words[i])
    } 
    
    //node已经构造完成，开始查找
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
            search(board, &node, i, j, rows, cols,&res)
        }
    }
  
    return res
}

func getTree(node *Node,word string) {
    for i:=0;i<len(word);i++{
        nowChar := int(word[i]-'a')
        if node.Next[nowChar] == nil {
            node.Next[nowChar] = &Node{Next:[26]*Node{}}
        }
        node = node.Next[nowChar]
    }
    node.Word = word
}

func search(board [][]byte, node *Node, i,j,rows, cols int, res *[]string) {
    if board[i][j] == '.' || node.Next[int(board[i][j] - 'a')] == nil {
        return 
    }
    
    tmp,index:= board[i][j],int(board[i][j] - 'a')
    
    if node.Next[index].Word != "" {
        *res = append(*res, node.Next[index].Word)
        node.Next[index].Word = ""   //大佬极限去重我拜读+抄袭一波，之前使用map去重的
    }
  
    board[i][j] = '.'
    //上
    if i>0{
        search(board,node.Next[index],i-1,j,rows,cols,res)
    }
    //下
    if i<rows-1{
        search(board,node.Next[index],i+1,j,rows,cols,res)
    }
    //左
    if j>0{
        search(board,node.Next[index],i,j-1,rows,cols,res)
    }
    //右
     if j<cols-1{
        search(board,node.Next[index],i,j+1,rows,cols,res)
    }
    
    board[i][j] = tmp
}




