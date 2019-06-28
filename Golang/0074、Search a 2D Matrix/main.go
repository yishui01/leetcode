func searchMatrix(matrix [][]int, target int) bool {
    //二分啊,对着每行的第一个元素二分，找到一行，target >= first && target < 下一行的first
    
    lens := len(matrix)
    
    if lens == 0 {
        return false
    }
    
    cols := len(matrix[0])
    
    if cols == 0 {
        return false
    }
    
    left := 0
    right := lens-1
    
    if target < matrix[0][0] || target >  matrix[lens-1][cols-1] {
        return false
    }
    
    indexRow := -1
    
    for left <= right {
        //找到一个当前行首元素小于target然后下一行又大于target的，如果没有下一行，那就看当前行的末尾
        mid := (left + right) / 2 
        if matrix[mid][0] <= target {
            if mid + 1 <= lens -1 && matrix[mid + 1][0] > target {
                //找到了，就是当前行内
                indexRow = mid
                break
            } else if  mid + 1 > lens -1 && matrix[mid][cols-1] >= target {
                //找到了，也是当前行内
                indexRow = mid
                break
            }
        }
        
        if matrix[mid][0] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
        
    }
    
    if indexRow == -1 {
        return false
    }
    
    for i:=0; i < cols; i++ {
        if matrix[indexRow][i] == target {
            return true
        }
    }
    
    return false
    
    
    
}
