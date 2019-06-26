func setZeroes(matrix [][]int)  {
    //遍历matrix每一个元素，将为0的那个元素的row和col分别记录到两个map中，最后根据map输出res
    lens := len(matrix)
    if lens == 0 {
        return 
    }
    mapsRow := make(map[int]bool) 
    mapsCol := make(map[int]bool)
    
    colLen := len(matrix[0])
    for i:= 0; i < lens; i++ {
        for j:=0; j < colLen; j++ {
            if matrix[i][j] == 0 {
                if _,ok := mapsRow[i];!ok {
                    mapsRow[i] = true
                }
                if _,ok := mapsCol[j];!ok {
                    mapsCol[j] = true
                }
            }
        }
    }
    
    for k,_ := range mapsRow {
        for i:=0; i < colLen; i++ {
            matrix[k][i] = 0
        }
    }
    for k,_ := range mapsCol {
        for i:=0; i < lens; i++ {
            matrix[i][k] = 0
        }
    }
    
}
