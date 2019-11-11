type NumMatrix struct {
    nums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    dp := make([][]int,len(matrix))
   
    if len(matrix) > 0&& len(matrix[0]) > 0 {
        for i:=0;i<len(dp);i++{

            dp[i] = make([]int,len(matrix[0]))
            
            for j:=0;j<len(matrix[0]);j++{
                
                dp[i][j] = matrix[i][j]
                
                //上面那个点+左边那个点-左上角那个点
                if i > 0 {
                    dp[i][j] += dp[i-1][j]
                }
                if j > 0 {
                    dp[i][j] += dp[i][j-1]
                }
                 
                if i>0 && j > 0{
                    dp[i][j] -= dp[i-1][j-1]
                }
                
            }
        }
    }
    return NumMatrix{nums:dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    if len(this.nums) == 0 {
        return 0
    }
    sum := this.nums[row2][col2]
    if row1 > 0 {
        sum -= this.nums[row1-1][col2]
    }
    if col1 > 0 {
        sum -= this.nums[row2][col1-1]
    }
    if row1 > 0 && col1 > 0 {
        sum += this.nums[row1-1][col1-1]
    }
    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
