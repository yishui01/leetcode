func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    if numRows <= 0 {
        return res
    }
    res[0] = []int{1}
    for i:=1; i < numRows; i++ {
        lens := i+1
        res[i] = make([]int, lens)
        res[i][0] = 1
        res[i][lens-1] = 1
        for j:=1; j < lens-1; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]
        }
    }
    
    return res
}
