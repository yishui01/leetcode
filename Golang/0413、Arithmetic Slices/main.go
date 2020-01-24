func numberOfArithmeticSlices(A []int) int {
    //求能划分出的等差子数组的个数
    //[1,2,3] 能划分出的子数组个数为1    (1)
    //[1,2,3,4]能划分出子数组的个数为3   (1+2)
    //[1,2,3,4,5]能划分出的子数组的个数为6  (1+2+3)
    //那么长度为n的等差数列能划分出的子数组为 1+2+...+n-2

    //设dp[i]为以i号元素为末尾元素的等差数组所具有的等差子数组个数
    //如果当前这个元素和之前的元素能组成一个等差数列，那么dp[i] = dp[i-1]+1
    
    if len(A) < 3 {
        return 0
    }

   /* dp := make([]int,len(A))
    res := 0

    for i:=2;i<len(A);i++{
        if A[i]-A[i-1] == A[i-1] - A[i-2] {
            dp[i] = dp[i-1]+1
        }
        res += dp[i]
    }*/

    //用一个变量代替DP数组
    res := 0
    tmp := 0
    for i:=2;i<len(A);i++{
        if A[i]-A[i-1] == A[i-1] - A[i-2] {
           tmp++
        } else {
            tmp = 0
        }
        res += tmp
    }
    return res
}
