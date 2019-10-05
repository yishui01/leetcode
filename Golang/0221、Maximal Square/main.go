func maximalSquare(matrix [][]byte) int {
    //dfs
    //边长1开始进行扩展，扩展成更新area变量，直到失败，将area变量与maxArea比对，进行选择更新
    
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])
    if cols == 0 {
        return 0
    }
    
    maxArea := 0
    
    for i:=0;i<rows;i++{
        for j:=0;j<cols;j++{
            if matrix[i][j] == '1' {
                dfs(i,j,rows,cols,1,matrix,&maxArea)
            }
        }
    }
    
    return maxArea
    
}

func dfs(i,j,rows,cols,border int, nums [][]byte, maxArea *int) {
    if i+border-1 >= rows || j+border-1 >= cols || i < 0 || j < 0 {
        return 
    }
    flag := 1
    //看看当前边长是否符合
    
    //1、在之前的边长面积范围区域进行扩展，倒数第一行之前都只要看最后一列
    for z:=0;z<border-1;z++{
        if nums[i+z][j+border-1] != '1' {
            flag = 0
            break
        }
    }
    //看最后一行
    for z:=0;z<border;z++{
        if nums[i+border-1][j+z] != '1' {
            flag = 0
            break
        }
    }
    
    if flag == 1 {
        if *maxArea < border*border{
            *maxArea = border*border
        }
        dfs(i,j,rows,cols,border+1,nums,maxArea)
    }
}
