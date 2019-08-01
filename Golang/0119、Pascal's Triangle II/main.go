func getRow(rowIndex int) []int {
    //滚动数组
    if rowIndex < 0 {
        return []int{}
    }
    
    res := make([]int, rowIndex+1)
    if rowIndex == 0 {
        return []int{1}
    }
    
    res[0] = 1
    for i:=1; i < rowIndex+1; i++ {
        res[i] = 1
        prev := res[0]
        for j:=1; j < i;j++{
            tmp := res[j]
            res[j] = prev+res[j]
            prev = tmp
        }
    }
    
    return res
}
