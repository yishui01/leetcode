func solveSudoku(board [][]byte)  {
    // 这个是回溯,递归
    //最外层遍历9*9, 当遇到第一个空格的时候，触发递归，并跳出循环，由第一个空格为起点，进行深度递归，当计数达到82时就是answer
    //每次进行下一轮递归之前保存当前表格的状态，用于回溯
    
    //优化版本，从92ms优化到68ms
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
    for i:=0; i < 9; i++ {
        for j := 0;j<9;j++ {
            if board[i][j] == '.' {
               // copyBoard := make([][]byte, 9)
               // mycopy(copyBoard, board)
                isFind := 0
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
       // copy(board,copyBoard)
        *isFind = 1
        return 
    }
    if i == 9{
        //已经到头了，
        return 
    }
    nextI,nextJ := findNextIJ(i, j)
    if copyBoard[i][j] == '.' {
        //循环该空格所有可能的数字
        possibleNums := findPossibleNum(copyBoard, i, j)
            for _,v := range possibleNums {
                 // tmpBoard := make([][]byte, 9)
                 // //当前层先保留一份coptBoard
                 // tmpBoard = mycopy(tmpBoard, copyBoard)
                 copyBoard[i][j] = v
               
                //将每一个可能的数字保存到copuBoard中
                recursion(copyBoard, level + 1, nextI, nextJ,isFind)
                if *isFind == 1 {
                    //找到答案了就不用回溯了，直接返回，精髓
                    return
                }
                copyBoard[i][j] = '.'
            }
        
    } else {
        recursion(copyBoard, level + 1, nextI,nextJ,isFind)
    }
    
}

// func mycopy(dist [][]byte, src [][]byte)[][]byte {
//     for k,_ := range src{
// 		dist[k] = append(dist[k], src[k]...)
// 	}
    
//     return dist
// }

//找出当前空格所有可能性
func findPossibleNum(copyBoard [][]byte, i int,j int) []byte {
    //1、横排
    //2、竖排
    //3、九宫格
    //4、如果以上三种情况中任意一种发生了冲突，直接返回空
    maps := make(map[byte]int)
    maps3 := make(map[byte]int)
    for z:=0; z < 9; z++ {
        if copyBoard[i][z] == '.' {
            continue
        }
        if _,ok := maps[copyBoard[i][z]];ok {
            return []byte{}
        } else {
            maps[copyBoard[i][z]] = 1
        }
    }
    //竖排,要排除本身
    for z:=0; z < 9; z++ {
        if copyBoard[z][j] == '.'{
            continue
        }
        if _,ok := maps3[copyBoard[z][j]];ok {
            return []byte{}
        } else {
            maps3[copyBoard[z][j]] = 1
        }
    }
    //九宫格,找出九宫格位置的起点,之前横排和竖排都验证过了，只要验证右下角的四格就行
    i = i/3*3 //九宫格左上角i
    j = j/3*3 //九宫格左上角j
   
    maps2 := make(map[byte]int)
    
    //现在i和j是九宫格的左上顶点
    for z:=0; z < 3;z++ {
        for c:=0; c < 3; c++ {
             if copyBoard[i+z][j+c] == '.'{
                  continue
                }
            if _,ok := maps2[copyBoard[i+z][j+c]];ok {
                return []byte{}
            } else {
                maps2[copyBoard[i+z][j+c]] = 1
            }
        }
    }
    
    res := make([]byte,0)
    for x:=1; x < 10; x++ {
        tmpKey := '0'+byte(x)
        _,ok :=maps[tmpKey];
        _,ok2 :=maps2[tmpKey];
        _,ok3 :=maps3[tmpKey]; 
        if !ok2 && !ok && !ok3{
            res = append(res,tmpKey)
        }
    }
    
    return res
    
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
