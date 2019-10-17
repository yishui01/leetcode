func searchMatrix(matrix [][]int, target int) bool {
    //选择左下角为起点，
    //target > 当前数  右移 （往右移动是增大的过程，那有个问题，是不是可能会漏掉，比如说，target在上一层的右半部分，而我们是在本层右移，
                            //会不会漏掉target？答案是不会，因为本层在右移到一定程度的时候，一定会遇到一个比target大的数，然后那个时候会上移
                            //每一次的右移，是因为target比当前数字大，才会右移，又因为从上往下是递增的，
                            //那么也就是说，target同时也比这一列的上面的所有数字都大，右移的时候相当于已经把这一列全部排除了。）
    
    //target < 当前数  上移（因为当前数右边的全部是大于当前数的，更加大于target了，所以右边的数不可能存在target）
    
    rows := len(matrix)
    if rows == 0 {
        return false
    }
    cols := len(matrix[0]) 
    if cols == 0 {
        return false
    }
    
    i:=rows-1
    j := 0
    for i < 0 || j >= cols{
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            j++
        } else {
            i--
        }
    }
    
    return false
}
