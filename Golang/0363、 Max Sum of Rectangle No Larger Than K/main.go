func maxSumSubmatrix(matrix [][]int, k int) int {
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])
    if cols == 0 {
        return 0
    }

    res := math.MinInt32
    sum := make([][]int,rows)

    for i:=0;i<rows;i++{
        sum[i] = make([]int,cols)
        for j:=0;j<cols;j++{
            t := matrix[i][j]
            if i > 0 {
                t += sum[i-1][j]
            }
            if j > 0 {
                t += sum[i][j-1]
            }
            if i > 0 && j > 0 {
                t -= sum[i-1][j-1]
            }
            sum[i][j] = t

            for r := 0;r <=i;r++ {
                for c := 0;c <= j;c++{
                    d := sum[i][j]
                    if r > 0 {
                        d -= sum[r-1][j]
                    }
                    if c > 0 {
                        d -= sum[i][c-1]
                    }
                    if r > 0 && c > 0 {
                        d += sum[r-1][c-1]
                    }
                    if d <= k {
                        res = max(res,d)
                    }
                }
            }
            if res == k {
                return k
            }
        }
    }
    return res

}

func max(a,b int)int{
    if a > b {
        return a
    }
    return b
}
