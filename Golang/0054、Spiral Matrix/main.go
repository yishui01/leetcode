func spiralOrder(matrix [][]int) []int {
    //这种就是找规律，然后硬写，没有技巧
    //左往右，直到边界，再往下，边界，再往左，边界，再网上，边界，再左往右
    //直到碰到边界进入下一次变换的时候检查是否还有下一轮
    row := len(matrix)
    if row == 0 {
        return []int{}
    }
    col := len(matrix[0])
    
    //起点
    startRow := 0
    startCol := 0
    
    //右边的边界
    rightBoard := col-1
    
    //下边的边界
    downBoard := row-1
    
    //左边的边界
    leftBoard := 0
    
    //上边的边界
    topBoard := 0
    
    res := make([]int, 0)
    
    flag := 0  //0为左往右1为右往下2为下往左3为左往上
    change := 1
    for {
         if change != 0 && flag != -1{
           res = append(res, matrix[startRow][startCol])
         }
        switch flag {
            case -1:
                return res
            case 0: //左往右
                if startCol < rightBoard {
                    change = 1
                    startCol++
                } else {
                    //碰到边界，检查是否还有下一层，有的话就转换flag
                    if startRow+1 <= downBoard {
                        flag = 1
                        topBoard++
                        change = 0
                    } else {
                        flag = -1
                    }
                }
            case 1: //右往下
                if startRow < downBoard {
                     change = 1
                    startRow++
                } else {
                    //要看左边还有没有
                    if startCol-1 >= leftBoard {
                        flag = 2
                        rightBoard--
                         change = 0
                    } else {
                        flag = -1
                    }
                }
            case 2: //下往左
                if startCol > leftBoard {
                     change = 1
                    startCol--
                } else {
                    if startRow-1 >= topBoard {
                        flag = 3
                        downBoard--
                        change = 0
                    } else {
                        flag = -1
                    }
                }
            case 3: //左往上
            if startRow > topBoard {
                 change = 1
                startRow--
            } else {
                //看右边还有没有
                if startCol + 1 <= rightBoard {
                    flag = 0
                    leftBoard++
                     change = 0
                } else {
                    flag = -1
                }
            }
        }
    }
    
    return res
    
}
