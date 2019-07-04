func maximalRectangle(matrix [][]byte) int {
    lens := len(matrix)
    if lens == 0 {
        return 0
    }
    
    cols := len(matrix[0])
    if cols == 0 {
        return 0
    }
    
     res := 0
    
    for i:=0; i < lens; i++ {
        tmp := make([]int, cols)
        for j:=0; j < cols; j++ {
            tmpI := i
            tmpEle := 0
            for tmpI >= 0 {
                if matrix[tmpI][j] == '0' {
                    break
                } else {
                    tmpI--
                    tmpEle++
                }
            }
            tmp[j] = tmpEle
        }
        tmpRes := largestRectangleArea(tmp)
        if tmpRes > res {
            res = tmpRes
        }
    }
    
    return res
   
}


func largestRectangleArea(heights []int) int {
  lens := len(heights)
    if lens == 0 {
        return 0
    }
    
    res := 0
    for i:=0; i < lens; i++ {
        if i+1 < lens && heights[i] <= heights[i+1]{
            continue
        }
        minH := heights[i]
        tmp := heights[i]
        for j:=i; j >=0; j-- {
            if heights[j] < minH {
                minH = heights[j]
            }
            nowArea := minH * (i-j + 1)
            if nowArea > tmp {
                tmp = nowArea
            }
        }
        
        if tmp > res {
            res = tmp
        }
    }
    
    return res
    
}


