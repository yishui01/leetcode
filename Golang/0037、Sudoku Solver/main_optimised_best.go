func solveSudoku(board [][]byte)  {
    // 这个是回溯,递归
    //最外层遍历9*9, 当遇到第一个空格的时候，触发递归，并跳出循环，由第一个空格为起点，进行深度递归，当计数达到82时就是answer
    //每次进行下一轮递归之前保存当前表格的状态，用于回溯
    
/******************优化版本1，从92ms优化到68ms************************************/
    //优化内容：
    //1、之前回溯到当前层的原始状态是在递归之前保存一个新的二维切片副本，递归完成后利用副本还原
    //但现在还原是通过在递归之后还原当前层所进行的操作，如下
    /*
       copyBoard[i][j] = v  //当前层的操作
        recursion(copyBoard, level + 1, nextI, nextJ,isFind) //进行递归
        if *isFind == 1 { //这里如果不加判断，那么在找到答案后，依然会回溯成原始状态，并且由于切片是引用类型，会影响到result
            //找到答案了就不用回溯了，直接返回，精髓
            return
        }
          copyBoard[i][j] = '.' //这个就是精髓回溯，只管回溯当前层的操作，那么每一层都会回溯当前层，达到全部回溯
    */
    //2、进行递归时直接传board，不需要传递第二个副本
    
 /*******************优化进阶版2：从68ms优化到4ms****************************/
    //之前每一层是先找出下一层所有合法的数字，然后遍历这些合法的数字进入下一层
    //现在修改为每层都会尝试在当前层的格子中尝试1-9这9个字符，合法的就填入格子进入下一层，下一层再在那一层尝试1-9，合法的进入下一层
    //
    
   isFind := 0
    
    for i:=0; i < 9; i++ {
        for j := 0;j<9;j++ {
            if board[i][j] == '.' {
               // copyBoard := make([][]byte, 9)
               // mycopy(copyBoard, board)
                recursion(board, i*9+j+1,i,j, &isFind)
                break
            }
        }
    }
}

/*
 board为当前层的表格
 level为当前计数，到达82时产生solution
*/
func recursion(copyBoard [][]byte, level int,i int, j int, isFind *int) {
    if *isFind == 1 {
        return
    }
    if level == 82{
        *isFind = 1
        return 
    }
    nextI,nextJ := findNextIJ(i, j)
    if copyBoard[i][j] == '.' {
        //之前是找出该空格所有可能的数字，现在换种思路，尝试各种数字，合法就填入
        for z:=byte('1');z <= byte('9'); z++ {
            if isValid(copyBoard, i,j,z) {
                copyBoard[i][j] = z //设置当前层状态
                recursion(copyBoard, level+1, nextI,nextJ,isFind)
                if *isFind == 1 {
                    return
                }
                copyBoard[i][j] = '.' //回溯当前层状态
            }
        }
       
    } else {
        recursion(copyBoard, level + 1, nextI,nextJ,isFind)
    }
    
}


//验证传进来的数字在当前位置是否合法
func isValid(board [][]byte, i int, j int, num byte) bool {
    
    bi := i/3*3
    bj := j/3*3
    //bi和bj是当前i和j所在九宫格的左上顶点
    for z:=0; z < 9; z++ {
        if board[i][z] == num || board[z][j] == num || board[bi+z/3][bj+z%3] == num {
            return false
        }
    }
    return true
}


//找出下一层i和j的位置
func findNextIJ(i,j int) (int,int) {
    if j<8 {
        j=j+1
    } else {
        j = 0
        i = i+1
    }
    return i,j
    
}
