func generateMatrix(n int) [][]int {
    //全部初始化为0，螺旋填数
    
    res := make([][]int, n)
    for i:=0; i < n; i++ {
        tmp := make([]int, n)
        res[i] = tmp
    }
    lens := n*n
    direction := 0 //方向，0为从左往右，1为从上往下，2为从右往左，3为从下往上
    //下面是四面墙，碰到墙后改变方向
    top := 0
    right := n-1
    bottom := n-1
    left :=0
    
    //起点
    row := 0
    col := 0
    
    for i :=0; i < lens; i++ {
        if row < n && col < n {
            res[row][col] = i+1
        }
        switch(direction) {
            case -1:
                return res
            case 0:
                if col == right {
                    //碰到墙
                    if row+1 <= bottom {
                        direction = 1
                        row++
                        top++
                    } else {
                        direction = -1
                    }
                } else {
                    col++
                }
                
            case 1:
                if row == bottom {
                    if col-1 >= left {
                        direction = 2
                        right--
                        col--
                    } else {
                        fmt.Println(777)
                        direction = -1
                    }
                } else {
                    row++
                }
            case 2:
                if col == left {
                    if row-1 >= top {
                        direction = 3
                        bottom--
                        row--
                    } else {
                       direction = -1
                    }
                } else {
                    col--
                }
            case 3:
                if row == top {
                    if col+1 <= right {
                        direction = 0
                        left++
                        col++
                    } else {
                           direction = -1
                        }
                } else {
                    row--
                }
                
        }
    }
    
    return res
}
