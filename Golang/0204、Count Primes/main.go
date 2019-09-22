func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    
    nums := make([]bool,n+1) //数组空间换时间，下标代表数字
    
    for i:=2;i*i<n;i++{
        if !nums[i]{
            t := n/i
            for j:=2;j<=t;j++{
                nums[i*j] = true
            }
        }
    }
    
    res := 0
    
    for i:=2;i<n;i++{
        if !nums[i]{
            res++
        }
    }
    
    return res
    
}

